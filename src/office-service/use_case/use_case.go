package use_case

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/office-service/config"
	"assesement-test-MicroServices/src/office-service/repository"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OfficeUseCase struct {
	pb.UnimplementedOfficeServiceServer
	DatabaseConfig   *config.DatabaseConfig
	OfficeRepository *repository.OfficeRepository
}

func NewOfficeUseCase(
	databaseConfig *config.DatabaseConfig,
	OfficeRepository *repository.OfficeRepository,
) *OfficeUseCase {
	return &OfficeUseCase{
		UnimplementedOfficeServiceServer: pb.UnimplementedOfficeServiceServer{},
		DatabaseConfig:                   databaseConfig,
		OfficeRepository:                 OfficeRepository,
	}
}

func (OfficeUseCase *OfficeUseCase) GetOfficeById(context context.Context, id *pb.ById) (result *pb.OfficeResponse, err error) {
	session, err := OfficeUseCase.DatabaseConfig.OfficeDB.Connection.StartSession()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase GetOfficeById is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase GetOfficeById is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	GetOfficeById, GetOfficeByIdErr := OfficeUseCase.OfficeRepository.GetOfficeById(OfficeUseCase.DatabaseConfig.OfficeDB.Connection, id.Id)
	if GetOfficeByIdErr != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("OfficeUseCase GetOfficeById is failed, GetOfficeById failed : %s", GetOfficeByIdErr)
		result = &pb.OfficeResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	if GetOfficeById == nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("Office UseCase GetOneById is failed, Office is not found by id %s", id)
		result = &pb.OfficeResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.OfficeResponse{
		Code:    int64(codes.OK),
		Message: "Office UseCase GetOneById is succeed.",
		Data:    GetOfficeById,
	}
	return result, commit
}

func (OfficeUseCase *OfficeUseCase) UpdateOffice(context context.Context, request *pb.UpdateOfficeRequest) (result *pb.OfficeResponse, err error) {
	begin := OfficeUseCase.DatabaseConfig.OfficeDB.Connection
	session, err := OfficeUseCase.DatabaseConfig.OfficeDB.Connection.StartSession()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase UpdateOffice is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase UpdateOffice is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	foundOffice, err := OfficeUseCase.OfficeRepository.GetOfficeById(begin, request.Id)
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Canceled),
			Message: "OfficeUseCase UpdateOffice is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if foundOffice == nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Canceled),
			Message: "OfficeOfficeCase UpdateOffice is failed, Office is not found by id " + request.Id,
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if request.BranchName != nil {
		foundOffice.BranchName = *request.BranchName
	}
	foundOffice.BranchCode = uuid.NewString()
	time := time.Now()
	foundOffice.UpdatedAt = timestamppb.New(time)
	patchedOffice, err := OfficeUseCase.OfficeRepository.PatchOneById(begin, request.Id, foundOffice)
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase UpdateOffice is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	result = &pb.OfficeResponse{
		Code:    int64(codes.OK),
		Message: "OfficeOfficeCase UpdateOffice is succeed.",
		Data:    patchedOffice,
	}
	return result, session.CommitTransaction(context)
}
func (OfficeUseCase *OfficeUseCase) CreateOffice(context context.Context, request *pb.CreateOfficeRequest) (result *pb.OfficeResponse, err error) {
	session, err := OfficeUseCase.DatabaseConfig.OfficeDB.Connection.StartSession()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase CreateOffice is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase CreateOffice is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	currentTime := null.NewTime(time.Now(), true)
	newOffice := &pb.Office{
		BranchName: request.BranchName,
		BranchCode: uuid.NewString(),
		CreatedAt:  timestamppb.New(currentTime.Time),
		UpdatedAt:  timestamppb.New(currentTime.Time),
	}
	createdOffice, err := OfficeUseCase.OfficeRepository.CreateOffice(OfficeUseCase.DatabaseConfig.OfficeDB.Connection, newOffice)
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase Register is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}

	result = &pb.OfficeResponse{
		Code:    int64(codes.OK),
		Message: "OfficeUseCase Register is succeed.",
		Data:    createdOffice,
	}
	return result, session.CommitTransaction(context)
}
func (OfficeUseCase *OfficeUseCase) DeleteOffice(context context.Context, id *pb.ById) (result *pb.OfficeResponse, err error) {
	session, err := OfficeUseCase.DatabaseConfig.OfficeDB.Connection.StartSession()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase DeleteOffice is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase DeleteOffice is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	deletedOffice, deletedOfficeErr := OfficeUseCase.OfficeRepository.DeleteOffice(OfficeUseCase.DatabaseConfig.OfficeDB.Connection, id.Id)
	if deletedOfficeErr != nil {
		err = session.AbortTransaction(context)
		result = &pb.OfficeResponse{
			Code:    int64(codes.Internal),
			Message: "OfficeOfficeCase DeleteOffice is failed, " + deletedOfficeErr.Error(),
			Data:    nil,
		}
		return result, err
	}
	if deletedOffice == nil {
		err = session.AbortTransaction(context)
		result = &pb.OfficeResponse{
			Code:    int64(codes.Canceled),
			Message: "OfficeOfficeCase DeleteOffice is failed, Office is not deleted by id, " + id.Id,
			Data:    nil,
		}
		return result, err
	}

	err = session.CommitTransaction(context)
	result = &pb.OfficeResponse{
		Code:    int64(codes.OK),
		Message: "OfficeOfficeCase DeleteOffice is succeed.",
		Data:    deletedOffice,
	}
	return result, err
}
func (OfficeUseCase *OfficeUseCase) ListOffices(context context.Context, empty *pb.Empty) (result *pb.OfficeResponseRepeated, err error) {
	session, err := OfficeUseCase.DatabaseConfig.OfficeDB.Connection.StartSession()
	if err != nil {
		result = &pb.OfficeResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase ListOffice is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.OfficeResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "OfficeUseCase ListOffice is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	ListOffice, err := OfficeUseCase.OfficeRepository.ListOffices(OfficeUseCase.DatabaseConfig.OfficeDB.Connection)
	if err != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("OfficeUseCase ListOffice is failed, query failed : %s", err)
		result = &pb.OfficeResponseRepeated{
			Code:    int64(codes.Internal),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}

	if ListOffice.Data == nil {
		rollback := session.AbortTransaction(context)
		result = &pb.OfficeResponseRepeated{
			Code:    int64(codes.Canceled),
			Message: "Office UseCase ListOffice is failed, data Office is empty ",
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.OfficeResponseRepeated{
		Code:    int64(codes.OK),
		Message: "Office UseCase ListOffice is succeed.",
		Data:    ListOffice.Data,
	}
	return result, commit
}
