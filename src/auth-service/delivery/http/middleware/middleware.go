package middleware

import (
	"assesement-test-MicroServices/src/auth-service/config"
	"assesement-test-MicroServices/src/auth-service/model/response"
	"assesement-test-MicroServices/src/auth-service/repository"
	"context"
	"net/http"
	"strings"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuthMiddleware struct {
	SessionRepository *repository.AuthRepository
	DatabaseConfig    *config.DatabaseConfig
}

func NewAuthMiddleware(sessionRepository repository.AuthRepository, databaseConfig *config.DatabaseConfig) *AuthMiddleware {
	return &AuthMiddleware{
		SessionRepository: &sessionRepository,
		DatabaseConfig:    databaseConfig,
	}
}

func (authMiddleware *AuthMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", 1)
		if token == "" {
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: Missing token",
			}
			response.NewResponse(w, result)
			return
		}

		session, err := authMiddleware.DatabaseConfig.AuthDB.Connection.StartSession()
		if err != nil {
			session.AbortTransaction(context.Background())
			result := &response.Response[interface{}]{
				Code:    http.StatusInternalServerError,
				Message: "transaction error",
			}
			response.NewResponse(w, result)
			return
		}

		findSession, err := authMiddleware.SessionRepository.FindOneByAccToken(authMiddleware.DatabaseConfig.AuthDB.Connection, token)
		if err != nil {
			session.AbortTransaction(context.Background())
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: token not found",
			}
			response.NewResponse(w, result)
			return
		}
		if findSession == nil {
			session.AbortTransaction(context.Background())
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: Invalid Token",
			}
			response.NewResponse(w, result)
			return
		}
		if findSession.AccessTokenExpiredAt == timestamppb.Now() {
			session.AbortTransaction(context.Background())
			result := &response.Response[interface{}]{
				Code:    http.StatusUnauthorized,
				Message: "Unauthorized: Token expired",
			}
			response.NewResponse(w, result)
			return
		}
		session.CommitTransaction(context.Background())
		next.ServeHTTP(w, r)
	})
}
