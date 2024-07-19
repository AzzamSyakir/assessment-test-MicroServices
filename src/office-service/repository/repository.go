package repository

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/auth-service/entity"
	"context"

	"github.com/guregu/null"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OfficeRepository struct {
}

func NewOfficeRepository() *OfficeRepository {
	OfficeRepository := &OfficeRepository{}
	return OfficeRepository
}

func (OfficeRepository *OfficeRepository) CreateOffice(begin *mongo.Client, toCreateOffice *pb.Office) (result *pb.Office, err error) {
	db := begin.Database("appDb")
	createAcc := bson.D{
		{Key: "branch_name", Value: toCreateOffice.BranchName},
		{Key: "branch_code", Value: toCreateOffice.BranchCode},
		{Key: "created_at", Value: toCreateOffice.CreatedAt},
		{Key: "updated_at", Value: toCreateOffice.UpdatedAt},
	}
	_, queryErr := db.Collection("offices").InsertOne(context.TODO(), createAcc)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateOffice
	err = nil
	return result, err
}

func (OfficeRepository *OfficeRepository) GetOfficeById(begin *mongo.Client, id string) (result *pb.Office, err error) {
	var foundOffice entity.Office
	db := begin.Database("appDb")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return result, err
	}
	queryErr := db.Collection("offices").FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&foundOffice)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = &pb.Office{
		BranchName: foundOffice.BranchName.String,
		BranchCode: foundOffice.BranchCode.String,
		CreatedAt:  timestamppb.New(foundOffice.CreatedAt.Time),
		UpdatedAt:  timestamppb.New(foundOffice.UpdatedAt.Time),
	}
	err = nil
	return result, err
}
func (OfficeRepository *OfficeRepository) PatchOneById(begin *mongo.Client, id string, toPatchOffice *pb.Office) (result *pb.Office, err error) {
	db := begin.Database("appDb")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "branch_name", Value: toPatchOffice.BranchName},
			{Key: "branch_code", Value: toPatchOffice.BranchCode},
			{Key: "created_at", Value: toPatchOffice.CreatedAt},
			{Key: "updated_at", Value: toPatchOffice.UpdatedAt},
		},
		},
	}
	_, queryErr := db.Collection("offices").UpdateOne(context.TODO(), filter, update)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	result = toPatchOffice
	err = nil
	return result, err
}

func (OfficeRepository *OfficeRepository) DeleteOffice(begin *mongo.Client, id string) (result *pb.Office, err error) {
	db := begin.Database("appDb")
	var foundOffice entity.Office
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	queryErr := db.Collection("offices").FindOne(context.Background(), filter).Decode(&foundOffice)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	_, deleteError := db.Collection("offices").DeleteOne(context.TODO(), filter)
	if deleteError != nil {
		return nil, err
	}
	result = &pb.Office{
		BranchName: foundOffice.BranchName.String,
		BranchCode: foundOffice.BranchCode.String,
		CreatedAt:  timestamppb.New(foundOffice.CreatedAt.Time),
		UpdatedAt:  timestamppb.New(foundOffice.UpdatedAt.Time),
	}
	err = nil
	return result, err
}

func (OfficeRepository *OfficeRepository) ListOffices(begin *mongo.Client) (result *pb.OfficeResponseRepeated, err error) {
	db := begin.Database("appDb")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("offices").Find(context.TODO(), bson.D{{}}, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var ListOfficessPb []*pb.Office
	var createdAt, updatedAt null.Time

	for cursor.Next(context.TODO()) {
		ListOffices := &entity.Office{}
		scanErr := cursor.Decode(&ListOffices)
		ListOffices.CreatedAt = createdAt
		ListOffices.UpdatedAt = updatedAt
		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListOfficesPb := &pb.Office{
			BranchName: ListOffices.BranchName.String,
			BranchCode: ListOffices.BranchCode.String,
			CreatedAt:  timestamppb.New(ListOffices.CreatedAt.Time),
			UpdatedAt:  timestamppb.New(ListOffices.UpdatedAt.Time),
		}
		ListOfficessPb = append(ListOfficessPb, ListOfficesPb)
	}

	result = &pb.OfficeResponseRepeated{
		Data: ListOfficessPb,
	}
	err = nil
	return result, err
}
