package repository

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/role-service/model"
	"context"

	"github.com/guregu/null"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RoleRepository struct {
}

func NewRoleRepository() *RoleRepository {
	RoleRepository := &RoleRepository{}
	return RoleRepository
}

func (RoleRepository *RoleRepository) CreateRole(begin *mongo.Client, toCreateRole *pb.Role) (result *pb.Role, err error) {
	db := begin.Database("db")
	createAcc := bson.D{
		{Key: "account_name", Value: toCreateRole.RoleName},
		{Key: "created_at", Value: toCreateRole.CreatedAt},
		{Key: "updated_at", Value: toCreateRole.UpdatedAt},
	}
	_, queryErr := db.Collection("roles").InsertOne(context.TODO(), createAcc)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}

	result = toCreateRole
	err = nil
	return result, err
}

func (RoleRepository *RoleRepository) GetRoleById(begin *mongo.Client, id string) (result *pb.Role, err error) {
	var foundRole model.Role
	db := begin.Database("db")
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return result, err
	}
	queryErr := db.Collection("roles").FindOne(context.Background(), bson.D{{Key: "_id", Value: objID}}).Decode(&foundRole)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	result = &pb.Role{
		RoleName:  foundRole.RoleName,
		CreatedAt: foundRole.CreatedAt,
		UpdatedAt: foundRole.UpdatedAt,
	}
	err = nil
	return result, err
}
func (RoleRepository *RoleRepository) PatchOneById(begin *mongo.Client, id string, toPatchRole *pb.Role) (result *pb.Role, err error) {
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
			{Key: "account_name", Value: toPatchRole.RoleName},
			{Key: "created_at", Value: toPatchRole.CreatedAt},
			{Key: "updated_at", Value: toPatchRole.UpdatedAt},
		},
		},
	}
	_, queryErr := db.Collection("roles").UpdateOne(context.TODO(), filter, update)
	if queryErr != nil {
		result = nil
		err = queryErr
		return
	}
	result = toPatchRole
	err = nil
	return result, err
}

func (RoleRepository *RoleRepository) DeleteRole(begin *mongo.Client, id string) (result *pb.Role, err error) {
	db := begin.Database("db")
	var foundRole model.Role
	objID, objErr := primitive.ObjectIDFromHex(id)
	if objErr != nil {
		result = nil
		err = objErr
		return
	}
	filter := bson.D{{Key: "_id", Value: objID}}
	queryErr := db.Collection("roles").FindOne(context.Background(), filter).Decode(&foundRole)
	if queryErr != nil {
		result = nil
		err = queryErr
		return result, err
	}
	_, deleteError := db.Collection("roles").DeleteOne(context.TODO(), filter)
	if deleteError != nil {
		return nil, err
	}
	result = &pb.Role{
		RoleName:  foundRole.RoleName,
		CreatedAt: foundRole.CreatedAt,
		UpdatedAt: foundRole.UpdatedAt,
	}
	err = nil
	return result, err
}

func (RoleRepository *RoleRepository) ListRole(begin *mongo.Client) (result *pb.RoleResponseRepeated, err error) {
	db := begin.Database("db")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("roles").Find(context.TODO(), bson.D{{}}, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var ListRolesPb []*pb.Role
	var createdAt, updatedAt null.Time

	for cursor.Next(context.TODO()) {
		ListRole := &model.Role{}
		scanErr := cursor.Decode(&ListRole)
		ListRole.CreatedAt = timestamppb.New(createdAt.Time)
		ListRole.UpdatedAt = timestamppb.New(updatedAt.Time)
		if scanErr != nil {
			result = nil
			err = scanErr
			return result, err
		}
		ListRolePb := &pb.Role{
			RoleName:  ListRole.RoleName,
			CreatedAt: ListRole.CreatedAt,
			UpdatedAt: ListRole.UpdatedAt,
		}
		ListRolesPb = append(ListRolesPb, ListRolePb)
	}

	result = &pb.RoleResponseRepeated{
		Data: ListRolesPb,
	}
	err = nil
	return result, err
}
