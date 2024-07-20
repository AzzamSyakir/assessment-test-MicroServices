package repository

import (
	"assesement-test-MicroServices/src/auth-service/entity"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
}

func NewAuthRepository() *AuthRepository {
	AuthRepository := &AuthRepository{}
	return AuthRepository
}

func (sessionRepository *AuthRepository) CreateSession(begin *mongo.Client, toCreateSession *entity.Session) (result *entity.Session, err error) {
	db := begin.Database("appDb")
	createSession := bson.D{
		{Key: "account_id", Value: toCreateSession.AccountId},
		{Key: "access_token", Value: toCreateSession.AccessToken},
		{Key: "refresh_token", Value: toCreateSession.RefreshToken},
		{Key: "access_token_expired_at", Value: toCreateSession.AccessTokenExpiredAt},
		{Key: "refresh_token_expired_at", Value: toCreateSession.RefreshTokenExpiredAt},
		{Key: "created_at", Value: toCreateSession.CreatedAt},
		{Key: "updated_at", Value: toCreateSession.UpdatedAt},
	}
	_, queryErr := db.Collection("sessions").InsertOne(context.TODO(), createSession)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateSession
	err = nil
	return result, err
}

func (sessionRepository *AuthRepository) FindOneByAccToken(begin *mongo.Client, accessToken string) (result *entity.Session, err error) {
	var foundSession *entity.Session
	db := begin.Database("appDb")
	queryErr := db.Collection("sessions").FindOne(context.Background(), bson.D{{Key: "access_token", Value: accessToken}}).Decode(&foundSession)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = foundSession
	err = nil
	return result, err
}

func (sessionRepository *AuthRepository) GetOneByAccountId(begin *mongo.Client, accountId string) (result *entity.Session, err error) {
	var foundSession *entity.Session
	db := begin.Database("appDb")
	err = db.Collection("sessions").FindOne(context.Background(), bson.D{{Key: "account_id", Value: accountId}}).Decode(&foundSession)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return foundSession, nil
}

func (sessionRepository *AuthRepository) FindOneByRefToken(begin *mongo.Client, refreshToken string) (result *entity.Session, err error) {
	var foundSession *entity.Session
	db := begin.Database("appDb")
	queryErr := db.Collection("sessions").FindOne(context.Background(), bson.D{{Key: "refresh_token", Value: refreshToken}}).Decode(&foundSession)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = foundSession
	err = nil
	return result, err
}

func (sessionRepository *AuthRepository) PatchOneById(begin *mongo.Client, id string, toPatchSession *entity.Session) (result *entity.Session, err error) {
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
			{Key: "account_id", Value: toPatchSession.AccountId},
			{Key: "access_token", Value: toPatchSession.AccessToken},
			{Key: "refresh_token", Value: toPatchSession.RefreshToken},
			{Key: "access_token_expired_at", Value: toPatchSession.AccessTokenExpiredAt},
			{Key: "refresh_token_expired_at", Value: toPatchSession.RefreshTokenExpiredAt},
			{Key: "updated_at", Value: toPatchSession.UpdatedAt},
		},
		},
	}
	_, queryErr := db.Collection("accounts").UpdateOne(context.TODO(), filter, update)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	result = toPatchSession
	err = nil
	return result, err
}
func (sessionRepository *AuthRepository) DeleteOneById(begin *mongo.Client, id string) (result *entity.Session, err error) {
	db := begin.Database("appDb")
	var foundSession *entity.Session
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	queryErr := db.Collection("accounts").FindOne(context.Background(), filter).Decode(&foundSession)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	_, deleteError := db.Collection("accounts").DeleteOne(context.TODO(), filter)
	if deleteError != nil {
		return nil, err
	}
	result = foundSession
	err = nil
	return result, err
}
