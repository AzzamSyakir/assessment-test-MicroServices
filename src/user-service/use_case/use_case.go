package use_case

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/user-service/config"
	"assesement-test-MicroServices/src/user-service/repository"
	"context"
	"fmt"
	"time"

	"github.com/guregu/null"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserUseCase struct {
	pb.UnimplementedUserServiceServer
	DatabaseConfig *config.DatabaseConfig
	UserRepository *repository.UserRepository
}

func NewUserUseCase(
	databaseConfig *config.DatabaseConfig,
	UserRepository *repository.UserRepository,
) *UserUseCase {
	return &UserUseCase{
		UnimplementedUserServiceServer: pb.UnimplementedUserServiceServer{},
		DatabaseConfig:                 databaseConfig,
		UserRepository:                 UserRepository,
	}
}

func (UserUseCase *UserUseCase) GetUserById(context context.Context, id *pb.ById) (result *pb.UserResponse, err error) {
	session, err := UserUseCase.DatabaseConfig.UserDB.Connection.StartSession()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase GetUserById is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase GetUserById is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	GetUserById, GetUserByIdErr := UserUseCase.UserRepository.GetUserById(UserUseCase.DatabaseConfig.UserDB.Connection, id.Id)
	if GetUserByIdErr != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("UserUseCase GetUserById is failed, GetUserById failed : %s", GetUserByIdErr)
		result = &pb.UserResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	if GetUserById == nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("User UseCase GetOneById is failed, User is not found by id %s", id)
		result = &pb.UserResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.UserResponse{
		Code:    int64(codes.OK),
		Message: "User UseCase GetOneById is succeed.",
		Data:    GetUserById,
	}
	return result, commit
}

func (UserUseCase *UserUseCase) UpdateUser(context context.Context, request *pb.UpdateUserRequest) (result *pb.UserResponse, err error) {
	begin := UserUseCase.DatabaseConfig.UserDB.Connection
	session, err := UserUseCase.DatabaseConfig.UserDB.Connection.StartSession()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase UpdateUser is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase UpdateUser is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	foundUser, err := UserUseCase.UserRepository.GetUserById(begin, request.Id)
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Canceled),
			Message: "UserUseCase UpdateUser is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if foundUser == nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Canceled),
			Message: "UserUserCase UpdateUser is failed, User is not found by id " + request.Id,
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if request.UserName != nil {
		foundUser.UserName = *request.UserName
	}
	if request.PostCode != nil {
		foundUser.PostCode = *request.PostCode
	}
	if request.Address != nil {
		foundUser.Address = *request.Address
	}
	if request.Province != nil {
		foundUser.Province = *request.Province
	}
	if request.City != nil {
		foundUser.City = *request.City
	}
	time := time.Now()
	foundUser.UpdatedAt = timestamppb.New(time)
	patchedUser, err := UserUseCase.UserRepository.PatchOneById(begin, request.Id, foundUser)
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase UpdateUser is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	result = &pb.UserResponse{
		Code:    int64(codes.OK),
		Message: "UserUserCase UpdateUser is succeed.",
		Data:    patchedUser,
	}
	return result, session.CommitTransaction(context)
}
func (UserUseCase *UserUseCase) CreateUser(context context.Context, request *pb.CreateUserRequest) (result *pb.UserResponse, err error) {
	session, err := UserUseCase.DatabaseConfig.UserDB.Connection.StartSession()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase Register is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase Register is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	currentTime := null.NewTime(time.Now(), true)
	newUser := &pb.User{
		UserName:  request.UserName,
		PostCode:  request.PostCode,
		Address:   request.Address,
		Province:  request.Province,
		City:      request.City,
		CreatedAt: timestamppb.New(currentTime.Time),
		UpdatedAt: timestamppb.New(currentTime.Time),
	}
	createdUser, err := UserUseCase.UserRepository.CreateUser(UserUseCase.DatabaseConfig.UserDB.Connection, newUser)
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase Register is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}

	result = &pb.UserResponse{
		Code:    int64(codes.OK),
		Message: "UserUseCase Register is succeed.",
		Data:    createdUser,
	}
	return result, session.CommitTransaction(context)
}
func (UserUseCase *UserUseCase) DeleteUser(context context.Context, id *pb.ById) (result *pb.UserResponse, err error) {
	session, err := UserUseCase.DatabaseConfig.UserDB.Connection.StartSession()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase DeleteUser is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUseCase DeleteUser is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	deletedUser, deletedUserErr := UserUseCase.UserRepository.DeleteUser(UserUseCase.DatabaseConfig.UserDB.Connection, id.Id)
	if deletedUserErr != nil {
		err = session.AbortTransaction(context)
		result = &pb.UserResponse{
			Code:    int64(codes.Internal),
			Message: "UserUserCase DeleteUser is failed, " + deletedUserErr.Error(),
			Data:    nil,
		}
		return result, err
	}
	if deletedUser == nil {
		err = session.AbortTransaction(context)
		result = &pb.UserResponse{
			Code:    int64(codes.Canceled),
			Message: "UserUserCase DeleteUser is failed, User is not deleted by id, " + id.Id,
			Data:    nil,
		}
		return result, err
	}

	err = session.CommitTransaction(context)
	result = &pb.UserResponse{
		Code:    int64(codes.OK),
		Message: "UserUserCase DeleteUser is succeed.",
		Data:    deletedUser,
	}
	return result, err
}
func (UserUseCase *UserUseCase) ListUsers(context context.Context, empty *pb.Empty) (result *pb.UserResponseRepeated, err error) {
	session, err := UserUseCase.DatabaseConfig.UserDB.Connection.StartSession()
	if err != nil {
		result = &pb.UserResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "UserUseCase ListUser is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.UserResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "UserUseCase ListUser is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	ListUser, err := UserUseCase.UserRepository.ListUser(UserUseCase.DatabaseConfig.UserDB.Connection)
	if err != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("UserUseCase ListUser is failed, query failed : %s", err)
		result = &pb.UserResponseRepeated{
			Code:    int64(codes.Internal),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}

	if ListUser.Data == nil {
		rollback := session.AbortTransaction(context)
		result = &pb.UserResponseRepeated{
			Code:    int64(codes.Canceled),
			Message: "User UseCase ListUser is failed, data User is empty ",
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.UserResponseRepeated{
		Code:    int64(codes.OK),
		Message: "User UseCase ListUser is succeed.",
		Data:    ListUser.Data,
	}
	return result, commit
}
