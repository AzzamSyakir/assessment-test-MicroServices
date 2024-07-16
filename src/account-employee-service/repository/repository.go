package repository

import (
	"assesement-test-MicroServices/grpc/pb"
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
	db := begin.Database("db")
	createAcc := bson.D{
		{Key: "account_name", Value: toCreateAccount.Accountname},
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

func (AccountRepository *AccountRepository) ListAccount(begin *mongo.Client) (result *pb.AccountResponseRepeated, err error) {
	db := begin.Database("db")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("accounts").Find(context.TODO(), nil, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var ListAccounts []*pb.Account
	var createdAt, updatedAt null.Time

	for cursor.Next(context.TODO()) {
		ListAccount := &pb.Account{}
		scanErr := cursor.Decode(&ListAccount)
		ListAccount.CreatedAt = timestamppb.New(createdAt.Time)
		ListAccount.UpdatedAt = timestamppb.New(updatedAt.Time)
		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListAccounts = append(ListAccounts, ListAccount)
	}

	result = &pb.AccountResponseRepeated{
		Data: ListAccounts,
	}
	err = nil
	return result, err
}

func (AccountRepository *AccountRepository) GetAccountById(begin *mongo.Client, id string) (result *pb.Account, err error) {
	var foundAccounts *pb.Account
	var queryErr error
	db := begin.Database("db")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		result = nil
		err = queryErr
		return result, err
	}

	queryErr = db.Collection("accounts").FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&foundAccounts)

	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = foundAccounts
	err = nil
	return result, err
}
func (AccountRepository *AccountRepository) PatchOneById(begin *mongo.Client, id string, toPatchAccount *pb.Account) (result *pb.Account, err error) {
	db := begin.Database("db")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "account_name", Value: toPatchAccount.Accountname},
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
	db := begin.Database("db")
	filter := bson.D{{Key: "id", Value: id}}
	returnOpt := options.FindOneAndDeleteOptions{}
	res := db.Collection("accounts").FindOneAndDelete(context.TODO(), filter, &returnOpt)

	err = res.Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, err
}
