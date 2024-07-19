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

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	AuthRepository := &AuthRepository{}
	return AuthRepository
}

func (AuthRepository *AuthRepository) CreateAuth(begin *mongo.Client, toCreateAuth *pb.Auth) (result *pb.Auth, err error) {
	db := begin.Database("appDb")
	createAcc := bson.D{
		{Key: "auth_name", Value: toCreateAuth.AuthName},
		{Key: "password", Value: toCreateAuth.Password},
		{Key: "created_at", Value: toCreateAuth.CreatedAt},
		{Key: "updated_at", Value: toCreateAuth.UpdatedAt},
	}
	_, queryErr := db.Collection("auths").InsertOne(context.TODO(), createAcc)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateAuth
	err = nil
	return result, err
}

func (AuthRepository *AuthRepository) GetAuthById(begin *mongo.Client, id string) (result *pb.Auth, err error) {
	var foundAuth entity.Auth
	db := begin.Database("appDb")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return result, err
	}
	queryErr := db.Collection("auths").FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&foundAuth)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = &pb.Auth{
		AuthName:  foundAuth.AuthName,
		Password:  foundAuth.Password,
		CreatedAt: foundAuth.CreatedAt,
		UpdatedAt: foundAuth.UpdatedAt,
	}
	err = nil
	return result, err
}
func (AuthRepository *AuthRepository) PatchOneById(begin *mongo.Client, id string, toPatchAuth *pb.Auth) (result *pb.Auth, err error) {
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
			{Key: "auth_name", Value: toPatchAuth.AuthName},
			{Key: "password", Value: toPatchAuth.Password},
			{Key: "created_at", Value: toPatchAuth.CreatedAt},
			{Key: "updated_at", Value: toPatchAuth.UpdatedAt},
		},
		},
	}
	_, queryErr := db.Collection("auths").UpdateOne(context.TODO(), filter, update)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	result = toPatchAuth
	err = nil
	return result, err
}

func (AuthRepository *AuthRepository) DeleteAuth(begin *mongo.Client, id string) (result *pb.Auth, err error) {
	db := begin.Database("appDb")
	var foundAuth entity.Auth
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	queryErr := db.Collection("auths").FindOne(context.Background(), filter).Decode(&foundAuth)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	_, deleteError := db.Collection("auths").DeleteOne(context.TODO(), filter)
	if deleteError != nil {
		return nil, err
	}
	result = &pb.Auth{
		AuthName:  foundAuth.AuthName,
		Password:  foundAuth.Password,
		CreatedAt: foundAuth.CreatedAt,
		UpdatedAt: foundAuth.UpdatedAt,
	}
	err = nil
	return result, err
}

func (AuthRepository *AuthRepository) ListAuth(begin *mongo.Client) (result *pb.AuthResponseRepeated, err error) {
	db := begin.Database("appDb")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("auths").Find(context.TODO(), bson.D{{}}, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var ListAuthsPb []*pb.Auth
	var createdAt, updatedAt null.Time

	for cursor.Next(context.TODO()) {
		ListAuth := &entity.Auth{}
		scanErr := cursor.Decode(&ListAuth)
		ListAuth.CreatedAt = timestamppb.New(createdAt.Time)
		ListAuth.UpdatedAt = timestamppb.New(updatedAt.Time)
		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListAuthPb := &pb.Auth{
			AuthName:  ListAuth.AuthName,
			Password:  ListAuth.Password,
			CreatedAt: ListAuth.CreatedAt,
			UpdatedAt: ListAuth.UpdatedAt,
		}
		ListAuthsPb = append(ListAuthsPb, ListAuthPb)
	}

	result = &pb.AuthResponseRepeated{
		Data: ListAuthsPb,
	}
	err = nil
	return result, err
}
