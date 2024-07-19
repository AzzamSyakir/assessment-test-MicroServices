package http

import (
	model_request "assesement-test-MicroServices/src/auth-service/model/request"
	"assesement-test-MicroServices/src/auth-service/model/response"
	"assesement-test-MicroServices/src/auth-service/use_case"
	"encoding/json"
	"net/http"
	"strings"
)

type AuthController struct {
	AuthUseCase   *use_case.AuthUseCase
	ExposeUseCase *use_case.ExposeUseCase
}

func NewAuthController(authUseCase *use_case.AuthUseCase, exposeUseCase *use_case.ExposeUseCase) *AuthController {
	authController := &AuthController{
		AuthUseCase:   authUseCase,
		ExposeUseCase: exposeUseCase,
	}
	return authController
}
func (authController *AuthController) CreateAccount(writer http.ResponseWriter, reader *http.Request) {

	request := &model_request.CreateAccountRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, decodeErr.Error(), 404)
	}

	result := authController.ExposeUseCase.CreateAccount(request)

	response.NewResponse(writer, result)
}
func (authController *AuthController) Login(writer http.ResponseWriter, reader *http.Request) {
	request := &model_request.LoginRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, decodeErr.Error(), 404)
	}
	foundUser, _ := authController.AuthUseCase.Login(request)
	response.NewResponse(writer, foundUser)
}
func (authController *AuthController) Logout(writer http.ResponseWriter, reader *http.Request) {
	token := reader.Header.Get("Authorization")
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	result, _ := authController.AuthUseCase.Logout(tokenString)
	response.NewResponse(writer, result)
}

func (authController *AuthController) GetNewAccessToken(writer http.ResponseWriter, reader *http.Request) {
	token := reader.Header.Get("Authorization")
	tokenString := strings.Replace(token, "Bearer ", "", 1)

	result, _ := authController.AuthUseCase.GetNewAccessToken(tokenString)
	response.NewResponse(writer, result)
}
