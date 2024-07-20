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

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	UserRepository := &UserRepository{}
	return UserRepository
}

func (UserRepository *UserRepository) CreateUser(begin *mongo.Client, toCreateUser *pb.User) (result *pb.User, err error) {
	db := begin.Database("appDb")
	createAcc := bson.D{
		{Key: "user_name", Value: toCreateUser.UserName},
		{Key: "address", Value: toCreateUser.Address},
		{Key: "post_code", Value: toCreateUser.PostCode},
		{Key: "province", Value: toCreateUser.Province},
		{Key: "city", Value: toCreateUser.City},
		{Key: "created_at", Value: toCreateUser.CreatedAt},
		{Key: "updated_at", Value: toCreateUser.UpdatedAt},
	}
	_, queryErr := db.Collection("users").InsertOne(context.TODO(), createAcc)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateUser
	err = nil
	return result, err
}

func (UserRepository *UserRepository) GetUserById(begin *mongo.Client, id string) (result *pb.User, err error) {
	var foundUser entity.User
	db := begin.Database("appDb")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return result, err
	}
	queryErr := db.Collection("users").FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&foundUser)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = &pb.User{
		UserName:  foundUser.UserName.String,
		PostCode:  foundUser.PostCode.String,
		City:      foundUser.City.String,
		Province:  foundUser.Province.String,
		CreatedAt: timestamppb.New(foundUser.CreatedAt.Time),
		UpdatedAt: timestamppb.New(foundUser.UpdatedAt.Time),
	}
	err = nil
	return result, err
}
func (UserRepository *UserRepository) PatchOneById(begin *mongo.Client, id string, toPatchUser *pb.User) (result *pb.User, err error) {
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
			{Key: "user_name", Value: toPatchUser.UserName},
			{Key: "address", Value: toPatchUser.Address},
			{Key: "post_code", Value: toPatchUser.PostCode},
			{Key: "province", Value: toPatchUser.Province},
			{Key: "city", Value: toPatchUser.City},
			{Key: "created_at", Value: toPatchUser.CreatedAt},
			{Key: "updated_at", Value: toPatchUser.UpdatedAt},
		},
		},
	}
	_, queryErr := db.Collection("users").UpdateOne(context.TODO(), filter, update)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	result = toPatchUser
	err = nil
	return result, err
}

func (UserRepository *UserRepository) DeleteUser(begin *mongo.Client, id string) (result *pb.User, err error) {
	db := begin.Database("appDb")
	var foundUser entity.User
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	queryErr := db.Collection("users").FindOne(context.Background(), filter).Decode(&foundUser)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	_, deleteError := db.Collection("users").DeleteOne(context.TODO(), filter)
	if deleteError != nil {
		return nil, err
	}
	result = &pb.User{
		UserName:  foundUser.UserName.String,
		PostCode:  foundUser.PostCode.String,
		City:      foundUser.City.String,
		Province:  foundUser.Province.String,
		CreatedAt: timestamppb.New(foundUser.CreatedAt.Time),
		UpdatedAt: timestamppb.New(foundUser.UpdatedAt.Time),
	}
	err = nil
	return result, err
}

func (UserRepository *UserRepository) ListUser(begin *mongo.Client) (result *pb.UserResponseRepeated, err error) {
	db := begin.Database("appDb")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("users").Find(context.TODO(), bson.D{{}}, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var ListUsersPb []*pb.User
	var createdAt, updatedAt null.Time

	for cursor.Next(context.TODO()) {
		ListUser := &entity.User{}
		scanErr := cursor.Decode(&ListUser)
		ListUser.CreatedAt = createdAt
		ListUser.UpdatedAt = updatedAt
		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListUserPb := &pb.User{
			UserName:  ListUser.UserName.String,
			PostCode:  ListUser.PostCode.String,
			City:      ListUser.City.String,
			Province:  ListUser.Province.String,
			CreatedAt: timestamppb.New(ListUser.CreatedAt.Time),
			UpdatedAt: timestamppb.New(ListUser.UpdatedAt.Time),
		}
		ListUsersPb = append(ListUsersPb, ListUserPb)
	}

	result = &pb.UserResponseRepeated{
		Data: ListUsersPb,
	}
	err = nil
	return result, err
}
