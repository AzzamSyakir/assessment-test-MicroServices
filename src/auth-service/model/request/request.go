package request

import (
	"github.com/guregu/null"
)

type LoginRequest struct {
	Email    null.String `json:"email"`
	Password null.String `json:"password"`
}

// account req
type AccountPatchOneByIdRequest struct {
	AccountName null.String `json:"account_name"`
	Password    null.String `json:"password"`
}
type CreateAccountRequest struct {
	AccountName null.String `json:"account_name"`
	Password    null.String `json:"password"`
}

// role req
type RolePatchOneByIdRequest struct {
	RoleName null.String `json:"role_name"`
}
type CreateRoleRequest struct {
	RoleName null.String `json:"role_name"`
}

// office req
type OfficePatchOneByIdRequest struct {
	BranchName null.String `json:"branch_name"`
}
type CreateOfficeRequest struct {
	BranchName null.String `json:"branch_name"`
}

// office req
type ScreenPatchOneByIdRequest struct {
	ScreenName null.String `json:"screen_name"`
}
type CreateScreenRequest struct {
	ScreenName null.String `json:"screen_name"`
}
