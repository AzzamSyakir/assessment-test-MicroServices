package request

import (
	"github.com/guregu/null"
)

type LoginRequest struct {
	Email    null.String `json:"email"`
	Password null.String `json:"password"`
}

// account req
type CreateAccountRequest struct {
	Name     null.String `json:"name"`
	Password null.String `json:"password"`
}
type AccountPatchOneByIdRequest struct {
	Name     null.String `json:"name"`
	Password null.String `json:"password"`
}

// office req
type CreateOffice struct {
	BranchName null.String `json:"branch_name"`
}
type OfficePatchOneByIdRequest struct {
	BranchName null.String `json:"branch_name"`
}

// role req
type CreateRole struct {
	RoleName null.String `json:"role_name"`
}
type RolePatchOneByIdRequest struct {
	RoleName null.String `json:"role_name"`
}

// screen req
type CreateScreen struct {
	ScreenName null.String `json:"screen_name"`
}
type ScreenPatchOneByIdRequest struct {
	ScreenName null.String `json:"screen_name"`
}
