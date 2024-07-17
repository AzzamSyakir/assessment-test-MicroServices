package model

import "google.golang.org/protobuf/types/known/timestamppb"

type Role struct {
	RoleName  string                 `bson:"role_name"`
	RoleCode  string                 `bson:"role_code"`
	CreatedAt *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt *timestamppb.Timestamp `bson:"updated_at"`
}
