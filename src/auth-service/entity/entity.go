package entity

import (
	"github.com/guregu/null"
)

type Session struct {
	Id                    null.String `json:"id"`
	AccountId             null.String `json:"account_id"`
	AccessToken           null.String `json:"access_token"`
	RefreshToken          null.String `json:"refresh_token"`
	AccessTokenExpiredAt  null.Time   `json:"access_token_expired_at"`
	RefreshTokenExpiredAt null.Time   `json:"refresh_token_expired_at"`
	CreatedAt             null.Time   `json:"created_at"`
	UpdatedAt             null.Time   `json:"updated_at"`
}

type Account struct {
	AccountName null.String `bson:"auth_name"`
	Password    null.String `bson:"password"`
	CreatedAt   null.Time   `bson:"created_at"`
	UpdatedAt   null.Time   `bson:"updated_at"`
}
type User struct {
	UserName  null.String `bson:"user_name"`
	PostCode  null.String `bson:"post_code"`
	Address   null.String `bson:"address"`
	Province  null.String `bson:"province"`
	City      null.String `bson:"city"`
	CreatedAt null.Time   `bson:"created_at"`
	UpdatedAt null.Time   `bson:"updated_at"`
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
