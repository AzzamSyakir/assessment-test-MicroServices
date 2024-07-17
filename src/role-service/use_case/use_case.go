package use_case

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/role-service/config"
	"assesement-test-MicroServices/src/role-service/repository"
	"context"
	"fmt"
	"time"

	"github.com/guregu/null"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RoleUseCase struct {
	pb.UnimplementedRoleServiceServer
	DatabaseConfig *config.DatabaseConfig
	RoleRepository *repository.RoleRepository
}

func NewRoleUseCase(
	databaseConfig *config.DatabaseConfig,
	RoleRepository *repository.RoleRepository,
) *RoleUseCase {
	return &RoleUseCase{
		UnimplementedRoleServiceServer: pb.UnimplementedRoleServiceServer{},
		DatabaseConfig:                 databaseConfig,
		RoleRepository:                 RoleRepository,
	}
}

func (RoleUseCase *RoleUseCase) GetRoleById(context context.Context, id *pb.ById) (result *pb.RoleResponse, err error) {
	session, err := RoleUseCase.DatabaseConfig.RoleDB.Connection.StartSession()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase Register is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase Register is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	GetRoleById, GetRoleByIdErr := RoleUseCase.RoleRepository.GetRoleById(RoleUseCase.DatabaseConfig.RoleDB.Connection, id.Id)
	if GetRoleByIdErr != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("RoleUseCase GetRoleById is failed, GetRoleById failed : %s", GetRoleByIdErr)
		result = &pb.RoleResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	if GetRoleById == nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("Role UseCase GetOneById is failed, Role is not found by id %s", id)
		result = &pb.RoleResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.RoleResponse{
		Code:    int64(codes.OK),
		Message: "Role UseCase GetOneById is succeed.",
		Data:    GetRoleById,
	}
	return result, commit
}

func (RoleUseCase *RoleUseCase) UpdateRole(context context.Context, request *pb.UpdateRoleRequest) (result *pb.RoleResponse, err error) {
	begin := RoleUseCase.DatabaseConfig.RoleDB.Connection
	session, err := RoleUseCase.DatabaseConfig.RoleDB.Connection.StartSession()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase UpdateRole is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase UpdateRole is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	foundRole, err := RoleUseCase.RoleRepository.GetRoleById(begin, request.Id)
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Canceled),
			Message: "RoleUseCase UpdateRole is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if foundRole == nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Canceled),
			Message: "RoleRoleCase UpdateRole is failed, Role is not found by id " + request.Id,
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if request.RoleName != nil {
		foundRole.RoleName = *request.RoleName
	}

	time := time.Now()
	foundRole.UpdatedAt = timestamppb.New(time)
	patchedRole, err := RoleUseCase.RoleRepository.PatchOneById(begin, request.Id, foundRole)
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase UpdateRole is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	result = &pb.RoleResponse{
		Code:    int64(codes.OK),
		Message: "RoleRoleCase UpdateRole is succeed.",
		Data:    patchedRole,
	}
	return result, session.CommitTransaction(context)
}
func (RoleUseCase *RoleUseCase) CreateRole(context context.Context, request *pb.CreateRoleRequest) (result *pb.RoleResponse, err error) {
	session, err := RoleUseCase.DatabaseConfig.RoleDB.Connection.StartSession()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase Register is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase Register is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	currentTime := null.NewTime(time.Now(), true)
	newRole := &pb.Role{
		RoleName:  request.RoleName,
		CreatedAt: timestamppb.New(currentTime.Time),
		UpdatedAt: timestamppb.New(currentTime.Time),
	}
	createdRole, err := RoleUseCase.RoleRepository.CreateRole(RoleUseCase.DatabaseConfig.RoleDB.Connection, newRole)
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase Register is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}

	result = &pb.RoleResponse{
		Code:    int64(codes.OK),
		Message: "RoleUseCase Register is succeed.",
		Data:    createdRole,
	}
	return result, session.CommitTransaction(context)
}
func (RoleUseCase *RoleUseCase) DeleteRole(context context.Context, id *pb.ById) (result *pb.RoleResponse, err error) {
	session, err := RoleUseCase.DatabaseConfig.RoleDB.Connection.StartSession()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase DeleteRole is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase DeleteRole is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	deletedRole, deletedRoleErr := RoleUseCase.RoleRepository.DeleteRole(RoleUseCase.DatabaseConfig.RoleDB.Connection, id.Id)
	if deletedRoleErr != nil {
		err = session.AbortTransaction(context)
		result = &pb.RoleResponse{
			Code:    int64(codes.Internal),
			Message: "RoleRoleCase DeleteRole is failed, " + deletedRoleErr.Error(),
			Data:    nil,
		}
		return result, err
	}
	if deletedRole == nil {
		err = session.AbortTransaction(context)
		result = &pb.RoleResponse{
			Code:    int64(codes.Canceled),
			Message: "RoleRoleCase DeleteRole is failed, Role is not deleted by id, " + id.Id,
			Data:    nil,
		}
		return result, err
	}

	err = session.CommitTransaction(context)
	result = &pb.RoleResponse{
		Code:    int64(codes.OK),
		Message: "RoleRoleCase DeleteRole is succeed.",
		Data:    deletedRole,
	}
	return result, err
}
func (RoleUseCase *RoleUseCase) ListRoles(context context.Context, empty *pb.Empty) (result *pb.RoleResponseRepeated, err error) {
	session, err := RoleUseCase.DatabaseConfig.RoleDB.Connection.StartSession()
	if err != nil {
		result = &pb.RoleResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase ListRole is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.RoleResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "RoleUseCase ListRole is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	ListRole, err := RoleUseCase.RoleRepository.ListRole(RoleUseCase.DatabaseConfig.RoleDB.Connection)
	if err != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("RoleUseCase ListRole is failed, query failed : %s", err)
		result = &pb.RoleResponseRepeated{
			Code:    int64(codes.Internal),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}

	if ListRole.Data == nil {
		rollback := session.AbortTransaction(context)
		result = &pb.RoleResponseRepeated{
			Code:    int64(codes.Canceled),
			Message: "Role UseCase ListRole is failed, data Role is empty ",
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.RoleResponseRepeated{
		Code:    int64(codes.OK),
		Message: "Role UseCase ListRole is succeed.",
		Data:    ListRole.Data,
	}
	return result, commit
}
