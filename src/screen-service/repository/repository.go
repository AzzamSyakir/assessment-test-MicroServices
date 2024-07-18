package repository

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/gateway-service/entity"
	"context"

	"github.com/guregu/null"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScreenRepository struct {
}

func NewScreenRepository() *ScreenRepository {
	ScreenRepository := &ScreenRepository{}
	return ScreenRepository
}

func (ScreenRepository *ScreenRepository) CreateScreen(begin *mongo.Client, toCreateScreen *pb.Screen) (result *pb.Screen, err error) {
	db := begin.Database("appDb")
	createAcc := bson.D{
		{Key: "screen_name", Value: toCreateScreen.ScreenName},
		{Key: "screen_code", Value: toCreateScreen.ScreenCode},
		{Key: "created_at", Value: toCreateScreen.CreatedAt},
		{Key: "updated_at", Value: toCreateScreen.UpdatedAt},
	}
	_, queryErr := db.Collection("screens").InsertOne(context.TODO(), createAcc)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateScreen
	err = nil
	return result, err
}

func (ScreenRepository *ScreenRepository) GetScreenById(begin *mongo.Client, id string) (result *pb.Screen, err error) {
	var foundScreen entity.Screen
	db := begin.Database("appDb")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return result, err
	}
	queryErr := db.Collection("screens").FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&foundScreen)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = &pb.Screen{
		ScreenName: foundScreen.ScreenName,
		ScreenCode: foundScreen.ScreenCode,
		CreatedAt:  foundScreen.CreatedAt,
		UpdatedAt:  foundScreen.UpdatedAt,
	}
	err = nil
	return result, err
}
func (ScreenRepository *ScreenRepository) PatchOneById(begin *mongo.Client, id string, toPatchScreen *pb.Screen) (result *pb.Screen, err error) {
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
			{Key: "screen_name", Value: toPatchScreen.ScreenName},
			{Key: "screen_code", Value: toPatchScreen.ScreenCode},
			{Key: "created_at", Value: toPatchScreen.CreatedAt},
			{Key: "updated_at", Value: toPatchScreen.UpdatedAt},
		},
		},
	}
	_, queryErr := db.Collection("screens").UpdateOne(context.TODO(), filter, update)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	result = toPatchScreen
	err = nil
	return result, err
}

func (ScreenRepository *ScreenRepository) DeleteScreen(begin *mongo.Client, id string) (result *pb.Screen, err error) {
	db := begin.Database("appDb")
	var foundScreen entity.Screen
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	queryErr := db.Collection("screens").FindOne(context.Background(), filter).Decode(&foundScreen)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	_, deleteError := db.Collection("screens").DeleteOne(context.TODO(), filter)
	if deleteError != nil {
		return nil, err
	}
	result = &pb.Screen{
		ScreenName: foundScreen.ScreenName,
		ScreenCode: foundScreen.ScreenCode,
		CreatedAt:  foundScreen.CreatedAt,
		UpdatedAt:  foundScreen.UpdatedAt,
	}
	err = nil
	return result, err
}

func (ScreenRepository *ScreenRepository) ListScreens(begin *mongo.Client) (result *pb.ScreenResponseRepeated, err error) {
	db := begin.Database("appDb")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("screens").Find(context.TODO(), bson.D{{}}, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var ListScreenssPb []*pb.Screen
	var createdAt, updatedAt null.Time

	for cursor.Next(context.TODO()) {
		ListScreens := &entity.Screen{}
		scanErr := cursor.Decode(&ListScreens)
		ListScreens.CreatedAt = timestamppb.New(createdAt.Time)
		ListScreens.UpdatedAt = timestamppb.New(updatedAt.Time)
		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListScreensPb := &pb.Screen{
			ScreenName: ListScreens.ScreenName,
			ScreenCode: ListScreens.ScreenCode,
			CreatedAt:  ListScreens.CreatedAt,
			UpdatedAt:  ListScreens.UpdatedAt,
		}
		ListScreenssPb = append(ListScreenssPb, ListScreensPb)
	}

	result = &pb.ScreenResponseRepeated{
		Data: ListScreenssPb,
	}
	err = nil
	return result, err
}
