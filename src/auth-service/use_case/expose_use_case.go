package use_case

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/auth-service/config"
	"assesement-test-MicroServices/src/auth-service/delivery/grpc/client"
	"assesement-test-MicroServices/src/auth-service/entity"
	model_request "assesement-test-MicroServices/src/auth-service/model/request"
	model_response "assesement-test-MicroServices/src/auth-service/model/response"
	"assesement-test-MicroServices/src/auth-service/repository"
	"net/http"

	"github.com/guregu/null"
)

type ExposeUseCase struct {
	DatabaseConfig *config.DatabaseConfig
	AuthRepository *repository.AuthRepository
	Env            *config.EnvConfig
	AccountClient  *client.AccountServiceClient
	RoleClient     *client.RoleServiceClient
	OfficeClient   *client.OfficeServiceClient
	ScreenClient   *client.ScreenServiceClient
	UserClient     *client.UserServiceClient
}

func NewExposeUseCase(
	databaseConfig *config.DatabaseConfig,
	authRepository *repository.AuthRepository,
	env *config.EnvConfig,
	initAccountClient *client.AccountServiceClient,
	initRoleClient *client.RoleServiceClient,
	initOfficeClient *client.OfficeServiceClient,
	initScreenClient *client.ScreenServiceClient,
	initUserClient *client.UserServiceClient,
) *ExposeUseCase {
	accountUseCase := &ExposeUseCase{
		AccountClient:  initAccountClient,
		RoleClient:     initRoleClient,
		OfficeClient:   initOfficeClient,
		ScreenClient:   initScreenClient,
		UserClient:     initUserClient,
		DatabaseConfig: databaseConfig,
		AuthRepository: authRepository,
		Env:            env,
	}
	return accountUseCase
}

// accounts
func (exposeUseCase *ExposeUseCase) ListAccounts() (result *model_response.Response[[]*entity.Account]) {
	ListAccount, err := exposeUseCase.AccountClient.ListAccounts()
	if err != nil {
		result = &model_response.Response[[]*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return result
	}
	var accounts []*entity.Account
	for _, account := range ListAccount.Data {
		accountData := &entity.Account{
			AccountName: null.NewString(account.AccountName, true),
			Password:    null.NewString(account.Password, true),
			CreatedAt:   null.NewTime(account.CreatedAt.AsTime(), true),
			UpdatedAt:   null.NewTime(account.UpdatedAt.AsTime(), true),
		}

		accounts = append(accounts, accountData)
	}
	bodyResponseAccount := &model_response.Response[[]*entity.Account]{
		Code:    http.StatusOK,
		Message: ListAccount.Message,
		Data:    accounts,
	}
	return bodyResponseAccount
}
func (exposeUseCase *ExposeUseCase) CreateAccount(request *model_request.CreateAccountRequest) (result *model_response.Response[*entity.Account]) {
	req := &pb.CreateAccountRequest{
		Name:     request.AccountName.String,
		Password: request.Password.String,
	}
	createAccount, err := exposeUseCase.AccountClient.CreateAccount(req)
	if err != nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: err.Error(),
		}
		return
	}
	if createAccount.Data == nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createAccount.Message,
		}
		return
	}
	account := entity.Account{
		AccountName: null.NewString(createAccount.Data.AccountName, true),
		Password:    null.NewString(createAccount.Data.Password, true),
		CreatedAt:   null.NewTime(createAccount.Data.CreatedAt.AsTime(), true),
		UpdatedAt:   null.NewTime(createAccount.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseAccount := &model_response.Response[*entity.Account]{
		Code:    http.StatusCreated,
		Message: createAccount.Message,
		Data:    &account,
	}
	return bodyResponseAccount
}
func (exposeUseCase *ExposeUseCase) DeleteAccount(id string) (result *model_response.Response[*entity.Account]) {
	DeleteAccount, err := exposeUseCase.AccountClient.DeleteAccount(id)
	if err != nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: DeleteAccount.Message,
			Data:    nil,
		}
		return
	}
	if DeleteAccount.Data == nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: DeleteAccount.Message,
			Data:    nil,
		}
		return
	}
	account := entity.Account{
		AccountName: null.NewString(DeleteAccount.Data.AccountName, true),
		Password:    null.NewString(DeleteAccount.Data.Password, true),
		CreatedAt:   null.NewTime(DeleteAccount.Data.CreatedAt.AsTime(), true),
		UpdatedAt:   null.NewTime(DeleteAccount.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseAccount := &model_response.Response[*entity.Account]{
		Code:    http.StatusOK,
		Message: DeleteAccount.Message,
		Data:    &account,
	}
	return bodyResponseAccount
}
func (exposeUseCase *ExposeUseCase) UpdateAccount(id string, request *model_request.AccountPatchOneByIdRequest) (result *model_response.Response[*entity.Account]) {
	req := &pb.UpdateAccountRequest{}
	if id != "" {
		req.Id = id
	}
	if request.AccountName.Valid {
		req.Name = &request.AccountName.String
	}
	if request.Password.Valid {
		req.Password = &request.Password.String
	}
	UpdateAccount, err := exposeUseCase.AccountClient.UpdateAccount(req)
	if err != nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: UpdateAccount.Message,
			Data:    nil,
		}
		return
	}
	if UpdateAccount.Data == nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: UpdateAccount.Message,
			Data:    nil,
		}
		return
	}
	account := entity.Account{
		AccountName: null.NewString(UpdateAccount.Data.AccountName, true),
		Password:    null.NewString(UpdateAccount.Data.Password, true),
		CreatedAt:   null.NewTime(UpdateAccount.Data.CreatedAt.AsTime(), true),
		UpdatedAt:   null.NewTime(UpdateAccount.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseAccount := &model_response.Response[*entity.Account]{
		Code:    http.StatusOK,
		Message: UpdateAccount.Message,
		Data:    &account,
	}
	return bodyResponseAccount
}
func (exposeUseCase *ExposeUseCase) DetailAccount(id string) (result *model_response.Response[*entity.Account]) {
	GetAccount, err := exposeUseCase.AccountClient.GetOneById(id)
	if err != nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: GetAccount.Message,
			Data:    nil,
		}
		return
	}
	if GetAccount.Data == nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: GetAccount.Message,
			Data:    nil,
		}
		return
	}
	account := entity.Account{
		AccountName: null.NewString(GetAccount.Data.AccountName, true),
		Password:    null.NewString(GetAccount.Data.Password, true),
		CreatedAt:   null.NewTime(GetAccount.Data.CreatedAt.AsTime(), true),
		UpdatedAt:   null.NewTime(GetAccount.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseAccount := &model_response.Response[*entity.Account]{
		Code:    http.StatusOK,
		Message: GetAccount.Message,
		Data:    &account,
	}
	return bodyResponseAccount
}
func (exposeUseCase *ExposeUseCase) GetOneByAccountName(accountName string) (result *model_response.Response[*entity.Account]) {
	GetAccount, err := exposeUseCase.AccountClient.GetOneByAccountName(accountName)
	if err != nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: GetAccount.Message,
			Data:    nil,
		}
		return
	}
	if GetAccount.Data == nil {
		result = &model_response.Response[*entity.Account]{
			Code:    http.StatusBadRequest,
			Message: GetAccount.Message,
			Data:    nil,
		}
		return
	}
	account := entity.Account{
		AccountName: null.NewString(GetAccount.Data.AccountName, true),
		Password:    null.NewString(GetAccount.Data.Password, true),
		CreatedAt:   null.NewTime(GetAccount.Data.CreatedAt.AsTime(), true),
		UpdatedAt:   null.NewTime(GetAccount.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseAccount := &model_response.Response[*entity.Account]{
		Code:    http.StatusOK,
		Message: GetAccount.Message,
		Data:    &account,
	}
	return bodyResponseAccount
}

// roles

func (exposeUseCase *ExposeUseCase) CreateRole(request *model_request.CreateRoleRequest) (result *model_response.Response[*entity.Role]) {
	req := &pb.CreateRoleRequest{
		RoleName: request.RoleName.String,
	}
	createRole, err := exposeUseCase.RoleClient.CreateRole(req)
	if err != nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createRole.Message,
		}
		return
	}
	if createRole.Data == nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createRole.Message,
		}
		return
	}
	account := entity.Role{
		RoleName:  null.NewString(createRole.Data.RoleName, true),
		RoleCode:  null.NewString(createRole.Data.RoleCode, true),
		CreatedAt: null.NewTime(createRole.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(createRole.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseRole := &model_response.Response[*entity.Role]{
		Code:    http.StatusCreated,
		Message: createRole.Message,
		Data:    &account,
	}
	return bodyResponseRole
}
func (exposeUseCase *ExposeUseCase) ListRoles() (result *model_response.Response[[]*entity.Role]) {
	ListRole, err := exposeUseCase.RoleClient.ListRoles()
	if err != nil {
		result = &model_response.Response[[]*entity.Role]{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return result
	}
	var roles []*entity.Role
	for _, role := range ListRole.Data {
		roleData := &entity.Role{
			RoleName:  null.NewString(role.RoleName, true),
			RoleCode:  null.NewString(role.RoleCode, true),
			CreatedAt: null.NewTime(role.CreatedAt.AsTime(), true),
			UpdatedAt: null.NewTime(role.UpdatedAt.AsTime(), true),
		}

		roles = append(roles, roleData)
	}
	bodyResponseRole := &model_response.Response[[]*entity.Role]{
		Code:    http.StatusOK,
		Message: ListRole.Message,
		Data:    roles,
	}
	return bodyResponseRole
}
func (exposeUseCase *ExposeUseCase) DeleteRole(id string) (result *model_response.Response[*entity.Role]) {
	DeleteRole, err := exposeUseCase.RoleClient.DeleteRole(id)
	if err != nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Message: DeleteRole.Message,
			Data:    nil,
		}
		return
	}
	if DeleteRole.Data == nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Message: DeleteRole.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Role{
		RoleName:  null.NewString(DeleteRole.Data.RoleName, true),
		RoleCode:  null.NewString(DeleteRole.Data.RoleCode, true),
		CreatedAt: null.NewTime(DeleteRole.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(DeleteRole.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseRole := &model_response.Response[*entity.Role]{
		Code:    http.StatusOK,
		Message: DeleteRole.Message,
		Data:    &role,
	}
	return bodyResponseRole
}
func (exposeUseCase *ExposeUseCase) UpdateRole(id string, request *model_request.RolePatchOneByIdRequest) (result *model_response.Response[*entity.Role]) {
	req := &pb.UpdateRoleRequest{}
	if id != "" {
		req.Id = id
	}
	if request.RoleName.Valid {
		req.RoleName = &request.RoleName.String
	}
	UpdateRole, err := exposeUseCase.RoleClient.UpdateRole(req)
	if err != nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Message: UpdateRole.Message,
			Data:    nil,
		}
		return
	}
	if UpdateRole.Data == nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Message: UpdateRole.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Role{
		RoleName:  null.NewString(UpdateRole.Data.RoleName, true),
		RoleCode:  null.NewString(UpdateRole.Data.RoleCode, true),
		CreatedAt: null.NewTime(UpdateRole.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(UpdateRole.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseRole := &model_response.Response[*entity.Role]{
		Code:    http.StatusOK,
		Message: UpdateRole.Message,
		Data:    &role,
	}
	return bodyResponseRole
}
func (exposeUseCase *ExposeUseCase) DetailRole(id string) (result *model_response.Response[*entity.Role]) {
	GetRole, err := exposeUseCase.RoleClient.GetRoleById(id)
	if err != nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Message: GetRole.Message,
			Data:    nil,
		}
		return
	}
	if GetRole.Data == nil {
		result = &model_response.Response[*entity.Role]{
			Code:    http.StatusBadRequest,
			Message: GetRole.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Role{
		RoleName:  null.NewString(GetRole.Data.RoleName, true),
		RoleCode:  null.NewString(GetRole.Data.RoleCode, true),
		CreatedAt: null.NewTime(GetRole.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(GetRole.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseRole := &model_response.Response[*entity.Role]{
		Code:    http.StatusOK,
		Message: GetRole.Message,
		Data:    &role,
	}
	return bodyResponseRole
}

// offices
func (exposeUseCase *ExposeUseCase) CreateOffice(request *model_request.CreateOfficeRequest) (result *model_response.Response[*entity.Office]) {
	req := &pb.CreateOfficeRequest{
		BranchName: request.BranchName.String,
	}
	createOffice, err := exposeUseCase.OfficeClient.CreateOffice(req)
	if err != nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createOffice.Message,
		}
		return
	}
	if createOffice.Data == nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createOffice.Message,
		}
		return
	}
	account := entity.Office{
		BranchName: null.NewString(createOffice.Data.BranchName, true),
		BranchCode: null.NewString(createOffice.Data.BranchCode, true),
		CreatedAt:  null.NewTime(createOffice.Data.CreatedAt.AsTime(), true),
		UpdatedAt:  null.NewTime(createOffice.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseOffice := &model_response.Response[*entity.Office]{
		Code:    http.StatusCreated,
		Message: createOffice.Message,
		Data:    &account,
	}
	return bodyResponseOffice
}
func (exposeUseCase *ExposeUseCase) ListOffices() (result *model_response.Response[[]*entity.Office]) {
	ListOffice, err := exposeUseCase.OfficeClient.ListOffices()
	if err != nil {
		result = &model_response.Response[[]*entity.Office]{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return result
	}
	var roles []*entity.Office
	for _, role := range ListOffice.Data {
		roleData := &entity.Office{
			BranchName: null.NewString(role.BranchName, true),
			BranchCode: null.NewString(role.BranchCode, true),
			CreatedAt:  null.NewTime(role.CreatedAt.AsTime(), true),
			UpdatedAt:  null.NewTime(role.UpdatedAt.AsTime(), true),
		}

		roles = append(roles, roleData)
	}
	bodyResponseOffice := &model_response.Response[[]*entity.Office]{
		Code:    http.StatusOK,
		Message: ListOffice.Message,
		Data:    roles,
	}
	return bodyResponseOffice
}
func (exposeUseCase *ExposeUseCase) DeleteOffice(id string) (result *model_response.Response[*entity.Office]) {
	DeleteOffice, err := exposeUseCase.OfficeClient.DeleteOffice(id)
	if err != nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Message: DeleteOffice.Message,
			Data:    nil,
		}
		return
	}
	if DeleteOffice.Data == nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Message: DeleteOffice.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Office{
		BranchName: null.NewString(DeleteOffice.Data.BranchName, true),
		BranchCode: null.NewString(DeleteOffice.Data.BranchCode, true),
		CreatedAt:  null.NewTime(DeleteOffice.Data.CreatedAt.AsTime(), true),
		UpdatedAt:  null.NewTime(DeleteOffice.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseOffice := &model_response.Response[*entity.Office]{
		Code:    http.StatusOK,
		Message: DeleteOffice.Message,
		Data:    &role,
	}
	return bodyResponseOffice
}
func (exposeUseCase *ExposeUseCase) UpdateOffice(id string, request *model_request.OfficePatchOneByIdRequest) (result *model_response.Response[*entity.Office]) {
	req := &pb.UpdateOfficeRequest{}
	if id != "" {
		req.Id = id
	}
	if request.BranchName.Valid {
		req.BranchName = &request.BranchName.String
	}
	UpdateOffice, err := exposeUseCase.OfficeClient.UpdateOffice(req)
	if err != nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Message: UpdateOffice.Message,
			Data:    nil,
		}
		return
	}
	if UpdateOffice.Data == nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Message: UpdateOffice.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Office{
		BranchName: null.NewString(UpdateOffice.Data.BranchName, true),
		BranchCode: null.NewString(UpdateOffice.Data.BranchCode, true),
		CreatedAt:  null.NewTime(UpdateOffice.Data.CreatedAt.AsTime(), true),
		UpdatedAt:  null.NewTime(UpdateOffice.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseOffice := &model_response.Response[*entity.Office]{
		Code:    http.StatusOK,
		Message: UpdateOffice.Message,
		Data:    &role,
	}
	return bodyResponseOffice
}
func (exposeUseCase *ExposeUseCase) DetailOffice(id string) (result *model_response.Response[*entity.Office]) {
	GetOffice, err := exposeUseCase.OfficeClient.GetOfficeById(id)
	if err != nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Message: GetOffice.Message,
			Data:    nil,
		}
		return
	}
	if GetOffice.Data == nil {
		result = &model_response.Response[*entity.Office]{
			Code:    http.StatusBadRequest,
			Message: GetOffice.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Office{
		BranchName: null.NewString(GetOffice.Data.BranchName, true),
		BranchCode: null.NewString(GetOffice.Data.BranchCode, true),
		CreatedAt:  null.NewTime(GetOffice.Data.CreatedAt.AsTime(), true),
		UpdatedAt:  null.NewTime(GetOffice.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseOffice := &model_response.Response[*entity.Office]{
		Code:    http.StatusOK,
		Message: GetOffice.Message,
		Data:    &role,
	}
	return bodyResponseOffice
}

// screen

func (exposeUseCase *ExposeUseCase) ListScreens() (result *model_response.Response[[]*entity.Screen]) {
	ListScreen, err := exposeUseCase.ScreenClient.ListScreens()
	if err != nil {
		result = &model_response.Response[[]*entity.Screen]{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return result
	}
	var screens []*entity.Screen
	for _, screen := range ListScreen.Data {
		screenData := &entity.Screen{
			ScreenName: null.NewString(screen.ScreenName, true),
			ScreenCode: null.NewString(screen.ScreenCode, true),
			CreatedAt:  null.NewTime(screen.CreatedAt.AsTime(), true),
			UpdatedAt:  null.NewTime(screen.UpdatedAt.AsTime(), true),
		}

		screens = append(screens, screenData)
	}
	bodyResponseScreen := &model_response.Response[[]*entity.Screen]{
		Code:    http.StatusOK,
		Message: ListScreen.Message,
		Data:    screens,
	}
	return bodyResponseScreen
}
func (exposeUseCase *ExposeUseCase) CreateScreen(request *model_request.CreateScreenRequest) (result *model_response.Response[*entity.Screen]) {
	req := &pb.CreateScreenRequest{
		ScreenName: request.ScreenName.String,
	}
	createScreen, err := exposeUseCase.ScreenClient.CreateScreen(req)
	if err != nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createScreen.Message,
		}
		return
	}
	if createScreen.Data == nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createScreen.Message,
		}
		return
	}
	account := entity.Screen{
		ScreenName: null.NewString(createScreen.Data.ScreenName, true),
		ScreenCode: null.NewString(createScreen.Data.ScreenCode, true),
		CreatedAt:  null.NewTime(createScreen.Data.CreatedAt.AsTime(), true),
		UpdatedAt:  null.NewTime(createScreen.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseScreen := &model_response.Response[*entity.Screen]{
		Code:    http.StatusCreated,
		Message: createScreen.Message,
		Data:    &account,
	}
	return bodyResponseScreen
}
func (exposeUseCase *ExposeUseCase) DeleteScreen(id string) (result *model_response.Response[*entity.Screen]) {
	DeleteScreen, err := exposeUseCase.ScreenClient.DeleteScreen(id)
	if err != nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Message: DeleteScreen.Message,
			Data:    nil,
		}
		return
	}
	if DeleteScreen.Data == nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Message: DeleteScreen.Message,
			Data:    nil,
		}
		return
	}
	screen := entity.Screen{
		ScreenName: null.NewString(DeleteScreen.Data.ScreenName, true),
		ScreenCode: null.NewString(DeleteScreen.Data.ScreenCode, true),
		CreatedAt:  null.NewTime(DeleteScreen.Data.CreatedAt.AsTime(), true),
		UpdatedAt:  null.NewTime(DeleteScreen.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseScreen := &model_response.Response[*entity.Screen]{
		Code:    http.StatusOK,
		Message: DeleteScreen.Message,
		Data:    &screen,
	}
	return bodyResponseScreen
}
func (exposeUseCase *ExposeUseCase) UpdateScreen(id string, request *model_request.ScreenPatchOneByIdRequest) (result *model_response.Response[*entity.Screen]) {
	req := &pb.UpdateScreenRequest{}
	if id != "" {
		req.Id = id
	}
	if request.ScreenName.Valid {
		req.ScreenName = &request.ScreenName.String
	}
	UpdateScreen, err := exposeUseCase.ScreenClient.UpdateScreen(req)
	if err != nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Message: UpdateScreen.Message,
			Data:    nil,
		}
		return
	}
	if UpdateScreen.Data == nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Message: UpdateScreen.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Screen{
		ScreenName: null.NewString(UpdateScreen.Data.ScreenName, true),
		ScreenCode: null.NewString(UpdateScreen.Data.ScreenCode, true),

		CreatedAt: null.NewTime(UpdateScreen.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(UpdateScreen.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseScreen := &model_response.Response[*entity.Screen]{
		Code:    http.StatusOK,
		Message: UpdateScreen.Message,
		Data:    &role,
	}
	return bodyResponseScreen
}
func (exposeUseCase *ExposeUseCase) DetailScreen(id string) (result *model_response.Response[*entity.Screen]) {
	GetScreen, err := exposeUseCase.ScreenClient.GetScreenById(id)
	if err != nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Message: GetScreen.Message,
			Data:    nil,
		}
		return
	}
	if GetScreen.Data == nil {
		result = &model_response.Response[*entity.Screen]{
			Code:    http.StatusBadRequest,
			Message: GetScreen.Message,
			Data:    nil,
		}
		return
	}
	role := entity.Screen{
		ScreenName: null.NewString(GetScreen.Data.ScreenName, true),
		ScreenCode: null.NewString(GetScreen.Data.ScreenCode, true),
		CreatedAt:  null.NewTime(GetScreen.Data.CreatedAt.AsTime(), true),
		UpdatedAt:  null.NewTime(GetScreen.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseScreen := &model_response.Response[*entity.Screen]{
		Code:    http.StatusOK,
		Message: GetScreen.Message,
		Data:    &role,
	}
	return bodyResponseScreen
}

// user
func (exposeUseCase *ExposeUseCase) CreateUser(request *model_request.CreateUserRequest) (result *model_response.Response[*entity.User]) {
	req := &pb.CreateUserRequest{
		UserName: request.UserName.String,
	}
	createUser, err := exposeUseCase.UserClient.CreateUser(req)
	if err != nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createUser.Message,
		}
		return
	}
	if createUser.Data == nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Data:    nil,
			Message: createUser.Message,
		}
		return
	}
	account := entity.User{
		UserName:  null.NewString(createUser.Data.UserName, true),
		PostCode:  null.NewString(createUser.Data.PostCode, true),
		Address:   null.NewString(createUser.Data.Address, true),
		Province:  null.NewString(createUser.Data.Province, true),
		City:      null.NewString(createUser.Data.City, true),
		CreatedAt: null.NewTime(createUser.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(createUser.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseUser := &model_response.Response[*entity.User]{
		Code:    http.StatusCreated,
		Message: createUser.Message,
		Data:    &account,
	}
	return bodyResponseUser
}
func (exposeUseCase *ExposeUseCase) ListUsers() (result *model_response.Response[[]*entity.User]) {
	ListUser, err := exposeUseCase.UserClient.ListUsers()
	if err != nil {
		result = &model_response.Response[[]*entity.User]{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		}
		return result
	}
	var screens []*entity.User
	for _, screen := range ListUser.Data {
		screenData := &entity.User{
			UserName:  null.NewString(screen.UserName, true),
			PostCode:  null.NewString(screen.PostCode, true),
			Address:   null.NewString(screen.Address, true),
			Province:  null.NewString(screen.Province, true),
			City:      null.NewString(screen.City, true),
			CreatedAt: null.NewTime(screen.CreatedAt.AsTime(), true),
			UpdatedAt: null.NewTime(screen.UpdatedAt.AsTime(), true),
		}

		screens = append(screens, screenData)
	}
	bodyResponseUser := &model_response.Response[[]*entity.User]{
		Code:    http.StatusOK,
		Message: ListUser.Message,
		Data:    screens,
	}
	return bodyResponseUser
}
func (exposeUseCase *ExposeUseCase) DeleteUser(id string) (result *model_response.Response[*entity.User]) {
	DeleteUser, err := exposeUseCase.UserClient.DeleteUser(id)
	if err != nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Message: DeleteUser.Message,
			Data:    nil,
		}
		return
	}
	if DeleteUser.Data == nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Message: DeleteUser.Message,
			Data:    nil,
		}
		return
	}
	screen := entity.User{
		UserName:  null.NewString(DeleteUser.Data.UserName, true),
		PostCode:  null.NewString(DeleteUser.Data.PostCode, true),
		Address:   null.NewString(DeleteUser.Data.Address, true),
		Province:  null.NewString(DeleteUser.Data.Province, true),
		City:      null.NewString(DeleteUser.Data.City, true),
		CreatedAt: null.NewTime(DeleteUser.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(DeleteUser.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseUser := &model_response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: DeleteUser.Message,
		Data:    &screen,
	}
	return bodyResponseUser
}
func (exposeUseCase *ExposeUseCase) UpdateUser(id string, request *model_request.UserPatchOneByIdRequest) (result *model_response.Response[*entity.User]) {
	req := &pb.UpdateUserRequest{}
	if id != "" {
		req.Id = id
	}
	if request.UserName.Valid {
		req.UserName = &request.UserName.String
	}
	UpdateUser, err := exposeUseCase.UserClient.UpdateUser(req)
	if err != nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Message: UpdateUser.Message,
			Data:    nil,
		}
		return
	}
	if UpdateUser.Data == nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Message: UpdateUser.Message,
			Data:    nil,
		}
		return
	}
	role := entity.User{
		UserName:  null.NewString(UpdateUser.Data.UserName, true),
		PostCode:  null.NewString(UpdateUser.Data.PostCode, true),
		Address:   null.NewString(UpdateUser.Data.Address, true),
		Province:  null.NewString(UpdateUser.Data.Province, true),
		City:      null.NewString(UpdateUser.Data.City, true),
		CreatedAt: null.NewTime(UpdateUser.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(UpdateUser.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseUser := &model_response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: UpdateUser.Message,
		Data:    &role,
	}
	return bodyResponseUser
}
func (exposeUseCase *ExposeUseCase) DetailUser(id string) (result *model_response.Response[*entity.User]) {
	GetUser, err := exposeUseCase.UserClient.GetUserById(id)
	if err != nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Message: GetUser.Message,
			Data:    nil,
		}
		return
	}
	if GetUser.Data == nil {
		result = &model_response.Response[*entity.User]{
			Code:    http.StatusBadRequest,
			Message: GetUser.Message,
			Data:    nil,
		}
		return
	}
	role := entity.User{
		UserName:  null.NewString(GetUser.Data.UserName, true),
		PostCode:  null.NewString(GetUser.Data.PostCode, true),
		Address:   null.NewString(GetUser.Data.Address, true),
		Province:  null.NewString(GetUser.Data.Province, true),
		City:      null.NewString(GetUser.Data.City, true),
		CreatedAt: null.NewTime(GetUser.Data.CreatedAt.AsTime(), true),
		UpdatedAt: null.NewTime(GetUser.Data.UpdatedAt.AsTime(), true),
	}
	bodyResponseUser := &model_response.Response[*entity.User]{
		Code:    http.StatusOK,
		Message: GetUser.Message,
		Data:    &role,
	}
	return bodyResponseUser
}
