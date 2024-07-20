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

type RoleRepository struct {
}

func NewRoleRepository() *RoleRepository {
	RoleRepository := &RoleRepository{}
	return RoleRepository
}

type MongoDataRole struct {
	ID        string                `bson:"_id,omitempty"`
	RoleName  string                `bson:"role_name"`
	RoleCode  string                `bson:"role_code"`
	CreatedAt timestamppb.Timestamp `bson:"created_at"`
	UpdatedAt timestamppb.Timestamp `bson:"updated_at"`
}

func (RoleRepository *RoleRepository) CreateRole(begin *mongo.Client, toCreateRole *pb.Role) (result *pb.Role, err error) {
	db := begin.Database("appDb")
	createAcc := bson.D{
		{Key: "role_name", Value: toCreateRole.RoleName},
		{Key: "role_code", Value: toCreateRole.RoleCode},
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
	var foundRole *pb.Role
	db := begin.Database("appDb")
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
	err = nil
	return result, err
}
func (RoleRepository *RoleRepository) PatchOneById(begin *mongo.Client, id string, toPatchRole *pb.Role) (result *pb.Role, err error) {
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
			{Key: "role_name", Value: toPatchRole.RoleName},
			{Key: "role_code", Value: toPatchRole.RoleCode},
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
	db := begin.Database("appDb")
	var foundRole *pb.Role
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
	err = nil
	return result, err
}

func (RoleRepository *RoleRepository) ListRoles(begin *mongo.Client) (result *pb.RoleResponseRepeated, err error) {
	db := begin.Database("appDb")
	findOptions := options.Find()
	cursor, cursorErr := db.Collection("roles").Find(context.TODO(), bson.D{{}}, findOptions)
	if cursorErr != nil {
		result = nil
		err = cursorErr
		return result, err
	}
	var listRolesPb []*pb.Role
	for cursor.Next(context.TODO()) {
		var role MongoDataRole
		if err := cursor.Decode(&role); err != nil {
			return nil, err
		}
		rolePb := &pb.Role{
			Id:        role.ID,
			RoleName:  role.RoleName,
			RoleCode:  role.RoleCode,
			CreatedAt: &role.CreatedAt,
			UpdatedAt: &role.UpdatedAt,
		}
		listRolesPb = append(listRolesPb, rolePb)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	result = &pb.RoleResponseRepeated{
		Data: listRolesPb,
	}
	err = nil
	return result, err
}
