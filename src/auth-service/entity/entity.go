package entity

import (
	"github.com/guregu/null"
)

type Account struct {
	AccountName null.String `bson:"auth_name"`
	Password    null.String `bson:"password"`
	CreatedAt   null.Time   `bson:"created_at"`
	UpdatedAt   null.Time   `bson:"updated_at"`
}
type Office struct {
	BranchName null.String `bson:"branch_name"`
	BranchCode null.String `bson:"branch_code"`
	CreatedAt  null.Time   `bson:"created_at"`
	UpdatedAt  null.Time   `bson:"updated_at"`
}
type Role struct {
	RoleName  null.String `bson:"role_name"`
	RoleCode  null.String `bson:"role_code"`
	CreatedAt null.Time   `bson:"created_at"`
	UpdatedAt null.Time   `bson:"updated_at"`
}

type Screen struct {
	ScreenName null.String `bson:"branch_name"`
	ScreenCode null.String `bson:"branch_code"`
	CreatedAt  null.Time   `bson:"created_at"`
	UpdatedAt  null.Time   `bson:"updated_at"`
}
