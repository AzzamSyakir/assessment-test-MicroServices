package model

import "google.golang.org/protobuf/types/known/timestamppb"

type Role struct {
	RoleName  string                 `bson:"account_name"`
	Password  string                 `bson:"password"`
	CreatedAt *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt *timestamppb.Timestamp `bson:"updated_at"`
}
