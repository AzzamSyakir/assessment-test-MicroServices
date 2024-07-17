package model

import "google.golang.org/protobuf/types/known/timestamppb"

type Screen struct {
	ScreenName string                 `bson:"branch_name"`
	ScreenCode string                 `bson:"branch_code"`
	CreatedAt  *timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt  *timestamppb.Timestamp `bson:"updated_at"`
}
