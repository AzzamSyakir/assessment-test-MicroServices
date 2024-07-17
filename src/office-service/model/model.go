package model

import "google.golang.org/protobuf/types/known/timestamppb"

type Office struct {
	BranchName string                 `bson:"branch_name"`
	BranchCode string                 `bson:"branch_code"`
	CreatedAt  *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt  *timestamppb.Timestamp `bson:"updated_at"`
}
