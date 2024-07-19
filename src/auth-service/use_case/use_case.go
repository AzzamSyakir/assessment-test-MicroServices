package use_case

import (
	"assesement-test-MicroServices/src/auth-service/config"
	"assesement-test-MicroServices/src/auth-service/delivery/grpc/client"
	"assesement-test-MicroServices/src/auth-service/entity"
	model_request "assesement-test-MicroServices/src/auth-service/model/request/controller"
	model_response "assesement-test-MicroServices/src/auth-service/model/response"
	"assesement-test-MicroServices/src/auth-service/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	DatabaseConfig *config.DatabaseConfig
	AuthRepository *repository.AuthRepository
	Env            *config.EnvConfig
	accountClient     *client.AccountServiceClient
}

func NewAuthUseCase(
	databaseConfig *config.DatabaseConfig,
	authRepository *repository.AuthRepository,
	env *config.EnvConfig,
	initAccountClient *client.AccountServiceClient,
) *AuthUseCase {
	authUseCase := &AuthUseCase{
		accountClient:     initAccountClient,
		DatabaseConfig: databaseConfig,
		AuthRepository: authRepository,
		Env:            env,
	}
	return authUseCase
}

func (authUseCase *AuthUseCase) Login(request *model_request.LoginRequest) (result *model_response.Response[*entity.Session], err error) {
	begin, err := authUseCase.DatabaseConfig.AuthDB.Connection.Begin()
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Login failed, begin fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	foundAccount, err := authUseCase.accountClient.GetAccountByEmail(request.Email.String)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: foundAccount.Message,
			Data:    nil,
		}
		return result, rollback
	}
	if foundAccount.Data == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: foundAccount.Message,
			Data:    nil,
		}
		return result, rollback
	}

	comparePasswordErr := bcrypt.CompareHashAndPassword([]byte(foundAccount.Data.Password), []byte(request.Password.String))
	if comparePasswordErr != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase Login is failed, password is not match.",
			Data:    nil,
		}
		return result, rollback
	}

	accessToken := null.NewString(uuid.NewString(), true)
	refreshToken := null.NewString(uuid.NewString(), true)
	currentTime := null.NewTime(time.Now(), true)
	accessTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Minute*10), true)
	refreshTokenExpiredAt := null.NewTime(currentTime.Time.Add(time.Hour*24*2), true)

	foundSession, err := authUseCase.AuthRepository.GetOneByAccountId(begin, foundAccount.Data.Id)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase Login failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	if foundSession != nil {

		foundSession.AccessToken = accessToken
		foundSession.RefreshToken = refreshToken
		foundSession.AccessTokenExpiredAt = accessTokenExpiredAt
		foundSession.RefreshTokenExpiredAt = refreshTokenExpiredAt
		foundSession.UpdatedAt = currentTime
		patchedSession, err := authUseCase.AuthRepository.PatchOneById(begin, foundSession.Id.String, foundSession)
		if err != nil {
			rollback := begin.Rollback()
			result = &model_response.Response[*entity.Session]{
				Code:    http.StatusBadRequest,
				Message: "AuthUseCase Login failed, query updateSession  fail, " + err.Error(),
				Data:    nil,
			}
			return result, rollback
		}

		commit := begin.Commit()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusOK,
			Message: "AuthUseCase Login is succeed",
			Data:    patchedSession,
		}
		return result, commit
	}

	newSession := &entity.Session{
		Id:                    null.NewString(uuid.NewString(), true),
		AccountId:                null.NewString(foundAccount.Data.Id, true),
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiredAt:  accessTokenExpiredAt,
		RefreshTokenExpiredAt: refreshTokenExpiredAt,
		CreatedAt:             currentTime,
		UpdatedAt:             currentTime,
	}

	createdSession, err := authUseCase.AuthRepository.CreateSession(begin, newSession)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase Login failed, query createSession fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	commit := begin.Commit()
	result = &model_response.Response[*entity.Session]{
		Code:    http.StatusOK,
		Message: "AuthUseCase Login is succeed",
		Data:    createdSession,
	}
	return result, commit
}

func (authUseCase *AuthUseCase) Logout(accessToken string) (result *model_response.Response[*entity.Session], err error) {
	begin, err := authUseCase.DatabaseConfig.AuthDB.Connection.Begin()
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase Logout failed, begin fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	foundSession, err := authUseCase.AuthRepository.FindOneByAccToken(begin, accessToken)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase Logout failed, Invalid token, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	if foundSession == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase Logout is failed, session is not found by access token.",
			Data:    nil,
		}
		return result, rollback
	}
	deletedSession, err := authUseCase.AuthRepository.DeleteOneById(begin, foundSession.Id.String)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase Logout failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	if deletedSession == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase Logout failed, delete session failed",
			Data:    nil,
		}
		return result, rollback
	}

	commit := begin.Commit()
	result = &model_response.Response[*entity.Session]{
		Code:    http.StatusOK,
		Message: "AuthUseCase Logout is succeed.",
		Data:    deletedSession,
	}
	return result, commit
}

func (authUseCase *AuthUseCase) GetNewAccessToken(refreshToken string) (result *model_response.Response[*entity.Session], err error) {
	begin, err := authUseCase.DatabaseConfig.AuthDB.Connection.Begin()
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusInternalServerError,
			Message: "AuthUseCase GetNewAccesToken failed, begin fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}
	foundSession, err := authUseCase.AuthRepository.FindOneByRefToken(begin, refreshToken)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase GetNewAccesToken failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	if foundSession == nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase GetNewAccesToken  failed, session is not found by refresh token.",
			Data:    nil,
		}
		return result, rollback
	}

	if foundSession.RefreshTokenExpiredAt.Time.Before(time.Now()) {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusNotFound,
			Message: "AuthUseCase GetNewAccessToken is failed, refresh token is expired.",
			Data:    nil,
		}
		return result, rollback
	}

	foundSession.AccessToken = null.NewString(uuid.NewString(), true)
	foundSession.UpdatedAt = null.NewTime(time.Now(), true)
	patchedSession, err := authUseCase.AuthRepository.PatchOneById(begin, foundSession.Id.String, foundSession)
	if err != nil {
		rollback := begin.Rollback()
		result = &model_response.Response[*entity.Session]{
			Code:    http.StatusBadRequest,
			Message: "AuthUseCase GetNewAccesToken  failed, query to db fail," + err.Error(),
			Data:    nil,
		}
		return result, rollback
	}

	commit := begin.Commit()
	result = &model_response.Response[*entity.Session]{
		Code:    http.StatusOK,
		Message: "AuthUseCase GetNewAccessToken is succeed.",
		Data:    patchedSession,
	}
	return result, commit

}
