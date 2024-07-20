package repository

import (
	"assesement-test-MicroServices/grpc/pb"
	"context"

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

type MongoDataUser struct {
	Id        string                `bson:"_id,omitempty"`
	UserName  string                `bson:"user_name,omitempty"`
	PostCode  string                `bson:"post_code"`
	Province  string                `bson:"province"`
	Address   string                `bson:"address"`
	City      string                `bson:"city"`
	CreatedAt timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt timestamppb.Timestamp `bson:"updated_at"`
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
	var foundUser pb.User
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
	var foundUser pb.User
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

	for cursor.Next(context.TODO()) {
		var user MongoDataUser
		scanErr := cursor.Decode(&user)

		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListUserPb := &pb.User{
			Id:        user.Id,
			UserName:  user.UserName,
			PostCode:  user.PostCode,
			Address:   user.Address,
			Province:  user.Province,
			City:      user.City,
			CreatedAt: &user.CreatedAt,
			UpdatedAt: &user.UpdatedAt,
		}
		ListUsersPb = append(ListUsersPb, ListUserPb)
	}

	result = &pb.UserResponseRepeated{
		Data: ListUsersPb,
	}
	err = nil
	return result, err
}
