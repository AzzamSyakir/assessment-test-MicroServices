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

type AccountRepository struct {
}

func NewAccountRepository() *AccountRepository {
	AccountRepository := &AccountRepository{}
	return AccountRepository
}

func (AccountRepository *AccountRepository) CreateAccount(begin *mongo.Client, toCreateAccount *pb.Account) (result *pb.Account, err error) {
	db := begin.Database("appDb")
	createAcc := bson.D{
		{Key: "account_name", Value: toCreateAccount.AccountName},
		{Key: "password", Value: toCreateAccount.Password},
		{Key: "created_at", Value: toCreateAccount.CreatedAt},
		{Key: "updated_at", Value: toCreateAccount.UpdatedAt},
	}
	_, queryErr := db.Collection("accounts").InsertOne(context.TODO(), createAcc)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateAccount
	err = nil
	return result, err
}

func (AccountRepository *AccountRepository) GetAccountById(begin *mongo.Client, id string) (result *pb.Account, err error) {
	var foundAccount entity.Account
	db := begin.Database("appDb")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return result, err
	}
	queryErr := db.Collection("accounts").FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&foundAccount)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = &pb.Account{
		AccountName: foundAccount.AccountName,
		Password:    foundAccount.Password,
		CreatedAt:   foundAccount.CreatedAt,
		UpdatedAt:   foundAccount.UpdatedAt,
	}
	err = nil
	return result, err
}
func (AccountRepository *AccountRepository) PatchOneById(begin *mongo.Client, id string, toPatchAccount *pb.Account) (result *pb.Account, err error) {
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
			{Key: "account_name", Value: toPatchAccount.AccountName},
			{Key: "password", Value: toPatchAccount.Password},
			{Key: "created_at", Value: toPatchAccount.CreatedAt},
			{Key: "updated_at", Value: toPatchAccount.UpdatedAt},
		},
		},
	}
	_, queryErr := db.Collection("accounts").UpdateOne(context.TODO(), filter, update)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	result = toPatchAccount
	err = nil
	return result, err
}

func (AccountRepository *AccountRepository) DeleteAccount(begin *mongo.Client, id string) (result *pb.Account, err error) {
	db := begin.Database("appDb")
	var foundAccount entity.Account
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	queryErr := db.Collection("accounts").FindOne(context.Background(), filter).Decode(&foundAccount)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	_, deleteError := db.Collection("accounts").DeleteOne(context.TODO(), filter)
	if deleteError != nil {
		return nil, err
	}
	result = &pb.Account{
		AccountName: foundAccount.AccountName,
		Password:    foundAccount.Password,
		CreatedAt:   foundAccount.CreatedAt,
		UpdatedAt:   foundAccount.UpdatedAt,
	}
	err = nil
	return result, err
}

func (AccountRepository *AccountRepository) ListAccount(begin *mongo.Client) (result *pb.AccountResponseRepeated, err error) {
	db := begin.Database("appDb")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("accounts").Find(context.TODO(), bson.D{{}}, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var ListAccountsPb []*pb.Account
	var createdAt, updatedAt null.Time

	for cursor.Next(context.TODO()) {
		ListAccount := &entity.Account{}
		scanErr := cursor.Decode(&ListAccount)
		ListAccount.CreatedAt = timestamppb.New(createdAt.Time)
		ListAccount.UpdatedAt = timestamppb.New(updatedAt.Time)
		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListAccountPb := &pb.Account{
			AccountName: ListAccount.AccountName,
			Password:    ListAccount.Password,
			CreatedAt:   ListAccount.CreatedAt,
			UpdatedAt:   ListAccount.UpdatedAt,
		}
		ListAccountsPb = append(ListAccountsPb, ListAccountPb)
	}

	result = &pb.AccountResponseRepeated{
		Data: ListAccountsPb,
	}
	err = nil
	return result, err
}
