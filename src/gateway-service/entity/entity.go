package entity

import "google.golang.org/protobuf/types/known/timestamppb"

type Account struct {
	AccountName string                 `bson:"account_name"`
	Password    string                 `bson:"password"`
	CreatedAt   *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt   *timestamppb.Timestamp `bson:"updated_at"`
}
type Office struct {
	BranchName string                 `bson:"branch_name"`
	BranchCode string                 `bson:"branch_code"`
	CreatedAt  *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt  *timestamppb.Timestamp `bson:"updated_at"`
}
type Role struct {
	RoleName  string                 `bson:"role_name"`
	RoleCode  string                 `bson:"role_code"`
	CreatedAt *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt *timestamppb.Timestamp `bson:"updated_at"`
}

type Screen struct {
	ScreenName string                 `bson:"branch_name"`
	ScreenCode string                 `bson:"branch_code"`
	CreatedAt  *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt  *timestamppb.Timestamp `bson:"updated_at"`
}
