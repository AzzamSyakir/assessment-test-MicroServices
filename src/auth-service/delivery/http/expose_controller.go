package http

import (
	model_request "assesement-test-MicroServices/src/auth-service/model/request"
	"assesement-test-MicroServices/src/auth-service/model/response"
	"assesement-test-MicroServices/src/auth-service/use_case"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type ExposeController struct {
	ExposeUseCase *use_case.ExposeUseCase
}

func NewExposeController(exposeUseCase *use_case.ExposeUseCase) *ExposeController {
	exposeController := &ExposeController{
		ExposeUseCase: exposeUseCase,
	}
	return exposeController
}

// accounts

func (exposeController *ExposeController) ListAccount(writer http.ResponseWriter, reader *http.Request) {
	ListAccount := exposeController.ExposeUseCase.ListAccounts()
	response.NewResponse(writer, ListAccount)
}
func (exposeController *ExposeController) CreateAccount(writer http.ResponseWriter, reader *http.Request) {

	request := &model_request.CreateAccountRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, decodeErr.Error(), 404)
	}

	result := exposeController.ExposeUseCase.CreateAccount(request)

	response.NewResponse(writer, result)
}
func (exposeController *ExposeController) DeleteAccount(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	result := exposeController.ExposeUseCase.DeleteAccount(id)

	response.NewResponse(writer, result)
}
func (exposeController *ExposeController) UpdateAccount(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	request := &model_request.AccountPatchOneByIdRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, decodeErr.Error(), 404)
	}

	result := exposeController.ExposeUseCase.UpdateAccount(id, request)

	response.NewResponse(writer, result)
}
func (expoaseController *ExposeController) DetailAccount(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	foundAccount := expoaseController.ExposeUseCase.DetailAccount(id)
	response.NewResponse(writer, foundAccount)
}
func (exposeController *ExposeController) GetOneByAccountName(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	name := vars["account_name"]

	foundAccount := exposeController.ExposeUseCase.GetOneByAccountName(name)
	response.NewResponse(writer, foundAccount)
}

// office

func (exposeController *ExposeController) ListOffices(writer http.ResponseWriter, reader *http.Request) {
	office := exposeController.ExposeUseCase.ListOffices()
	response.NewResponse(writer, office)
}
func (exposeController *ExposeController) CreateOffice(writer http.ResponseWriter, reader *http.Request) {

	request := &model_request.CreateOfficeRequest{}

	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, decodeErr.Error(), 404)
	}

	result := exposeController.ExposeUseCase.CreateOffice(request)

	response.NewResponse(writer, result)
}

func (exposeController *ExposeController) DeleteOffice(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	result := exposeController.ExposeUseCase.DeleteOffice(id)

	response.NewResponse(writer, result)
}

func (exposeController *ExposeController) UpdateOffice(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	request := &model_request.OfficePatchOneByIdRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}
	result := exposeController.ExposeUseCase.UpdateOffice(id, request)

	response.NewResponse(writer, result)
}
func (exposeController *ExposeController) DetailOffice(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]
	foundOffice := exposeController.ExposeUseCase.DetailOffice(id)
	response.NewResponse(writer, foundOffice)
}

// role

func (exposeController *ExposeController) CreateRole(writer http.ResponseWriter, reader *http.Request) {

	request := &model_request.CreateRoleRequest{}

	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, "Failed to decode request body: "+decodeErr.Error(), http.StatusBadRequest)
		return
	}

	result := exposeController.ExposeUseCase.CreateRole(request)

	response.NewResponse(writer, result)
}

func (exposeController *ExposeController) ListRoles(writer http.ResponseWriter, reader *http.Request) {
	foundRole := exposeController.ExposeUseCase.ListRoles()
	response.NewResponse(writer, foundRole)
}

func (exposeController *ExposeController) DeleteRole(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	result := exposeController.ExposeUseCase.DeleteRole(id)

	response.NewResponse(writer, result)
}

func (exposeController *ExposeController) UpdateRole(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	request := &model_request.RolePatchOneByIdRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}
	result := exposeController.ExposeUseCase.UpdateRole(id, request)

	response.NewResponse(writer, result)
}
func (exposeController *ExposeController) DetailRole(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]
	foundRole := exposeController.ExposeUseCase.DetailRole(id)
	response.NewResponse(writer, foundRole)
}

// screen

func (exposeController *ExposeController) CreateScreen(writer http.ResponseWriter, reader *http.Request) {

	request := &model_request.CreateScreenRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		http.Error(writer, "Failed to decode request body: "+decodeErr.Error(), http.StatusBadRequest)
		return
	}
	result := exposeController.ExposeUseCase.CreateScreen(request)
	response.NewResponse(writer, result)
}

func (exposeController *ExposeController) DetailScreen(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]
	foundScreen := exposeController.ExposeUseCase.DetailScreen(id)
	response.NewResponse(writer, foundScreen)
}

func (exposeController *ExposeController) ListScreens(writer http.ResponseWriter, reader *http.Request) {
	foundScreens := exposeController.ExposeUseCase.ListScreens()
	response.NewResponse(writer, foundScreens)
}

func (exposeController *ExposeController) UpdateScreen(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	request := &model_request.ScreenPatchOneByIdRequest{}
	decodeErr := json.NewDecoder(reader.Body).Decode(request)
	if decodeErr != nil {
		panic(decodeErr)
	}
	result := exposeController.ExposeUseCase.UpdateScreen(id, request)

	response.NewResponse(writer, result)
}

func (exposeController *ExposeController) DeleteScreen(writer http.ResponseWriter, reader *http.Request) {
	vars := mux.Vars(reader)
	id := vars["id"]

	result := exposeController.ExposeUseCase.DeleteScreen(id)

	response.NewResponse(writer, result)
}
