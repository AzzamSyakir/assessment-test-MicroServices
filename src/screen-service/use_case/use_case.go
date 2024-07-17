package use_case

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/screen-service/config"
	"assesement-test-MicroServices/src/screen-service/repository"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/guregu/null"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ScreenUseCase struct {
	pb.UnimplementedScreenServiceServer
	DatabaseConfig   *config.DatabaseConfig
	ScreenRepository *repository.ScreenRepository
}

func NewScreenUseCase(
	databaseConfig *config.DatabaseConfig,
	ScreenRepository *repository.ScreenRepository,
) *ScreenUseCase {
	return &ScreenUseCase{
		UnimplementedScreenServiceServer: pb.UnimplementedScreenServiceServer{},
		DatabaseConfig:                   databaseConfig,
		ScreenRepository:                 ScreenRepository,
	}
}

func (ScreenUseCase *ScreenUseCase) GetScreenById(context context.Context, id *pb.ById) (result *pb.ScreenResponse, err error) {
	session, err := ScreenUseCase.DatabaseConfig.ScreenDB.Connection.StartSession()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase GetScreenById is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase GetScreenById is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	GetScreenById, GetScreenByIdErr := ScreenUseCase.ScreenRepository.GetScreenById(ScreenUseCase.DatabaseConfig.ScreenDB.Connection, id.Id)
	if GetScreenByIdErr != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("ScreenUseCase GetScreenById is failed, GetScreenById failed : %s", GetScreenByIdErr)
		result = &pb.ScreenResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	if GetScreenById == nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("Screen UseCase GetOneById is failed, Screen is not found by id %s", id)
		result = &pb.ScreenResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.ScreenResponse{
		Code:    int64(codes.OK),
		Message: "Screen UseCase GetOneById is succeed.",
		Data:    GetScreenById,
	}
	return result, commit
}

func (ScreenUseCase *ScreenUseCase) UpdateScreen(context context.Context, request *pb.UpdateScreenRequest) (result *pb.ScreenResponse, err error) {
	begin := ScreenUseCase.DatabaseConfig.ScreenDB.Connection
	session, err := ScreenUseCase.DatabaseConfig.ScreenDB.Connection.StartSession()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase UpdateScreen is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase UpdateScreen is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	founScreend, err := ScreenUseCase.ScreenRepository.GetScreenById(begin, request.Id)
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Canceled),
			Message: "ScreenUseCase UpdateScreen is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if founScreend == nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Canceled),
			Message: "ScreenUseCase UpdateScreen is failed, Screen is not found by id " + request.Id,
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if request.ScreenName != nil {
		founScreend.ScreenName = *request.ScreenName
	}
	founScreend.ScreenCode = uuid.NewString()
	time := time.Now()
	founScreend.UpdatedAt = timestamppb.New(time)
	patcheScreend, err := ScreenUseCase.ScreenRepository.PatchOneById(begin, request.Id, founScreend)
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase UpdateScreen is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	result = &pb.ScreenResponse{
		Code:    int64(codes.OK),
		Message: "ScreenUseCase UpdateScreen is succeed.",
		Data:    patcheScreend,
	}
	return result, session.CommitTransaction(context)
}
func (ScreenUseCase *ScreenUseCase) CreateScreen(context context.Context, request *pb.CreateScreenRequest) (result *pb.ScreenResponse, err error) {
	session, err := ScreenUseCase.DatabaseConfig.ScreenDB.Connection.StartSession()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase CreateScreen is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase CreateScreen is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	currentTime := null.NewTime(time.Now(), true)
	newScreen := &pb.Screen{
		ScreenName: request.ScreenName,
		ScreenCode: uuid.NewString(),
		CreatedAt:  timestamppb.New(currentTime.Time),
		UpdatedAt:  timestamppb.New(currentTime.Time),
	}
	createScreend, err := ScreenUseCase.ScreenRepository.CreateScreen(ScreenUseCase.DatabaseConfig.ScreenDB.Connection, newScreen)
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase Register is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}

	result = &pb.ScreenResponse{
		Code:    int64(codes.OK),
		Message: "ScreenUseCase Register is succeed.",
		Data:    createScreend,
	}
	return result, session.CommitTransaction(context)
}
func (ScreenUseCase *ScreenUseCase) DeleteScreen(context context.Context, id *pb.ById) (result *pb.ScreenResponse, err error) {
	session, err := ScreenUseCase.DatabaseConfig.ScreenDB.Connection.StartSession()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase DeleteScreen is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase DeleteScreen is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	deleteScreend, deleteScreendErr := ScreenUseCase.ScreenRepository.DeleteScreen(ScreenUseCase.DatabaseConfig.ScreenDB.Connection, id.Id)
	if deleteScreendErr != nil {
		err = session.AbortTransaction(context)
		result = &pb.ScreenResponse{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase DeleteScreen is failed, " + deleteScreendErr.Error(),
			Data:    nil,
		}
		return result, err
	}
	if deleteScreend == nil {
		err = session.AbortTransaction(context)
		result = &pb.ScreenResponse{
			Code:    int64(codes.Canceled),
			Message: "ScreenUseCase DeleteScreen is failed, Screen is not deleted by id, " + id.Id,
			Data:    nil,
		}
		return result, err
	}

	err = session.CommitTransaction(context)
	result = &pb.ScreenResponse{
		Code:    int64(codes.OK),
		Message: "ScreenUseCase DeleteScreen is succeed.",
		Data:    deleteScreend,
	}
	return result, err
}
func (ScreenUseCase *ScreenUseCase) ListScreens(context context.Context, empty *pb.Empty) (result *pb.ScreenResponseRepeated, err error) {
	session, err := ScreenUseCase.DatabaseConfig.ScreenDB.Connection.StartSession()
	if err != nil {
		result = &pb.ScreenResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase ListScreen is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.ScreenResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "ScreenUseCase ListScreen is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	ListScreen, err := ScreenUseCase.ScreenRepository.ListScreens(ScreenUseCase.DatabaseConfig.ScreenDB.Connection)
	if err != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("ScreenUseCase ListScreen is failed, query failed : %s", err)
		result = &pb.ScreenResponseRepeated{
			Code:    int64(codes.Internal),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}

	if ListScreen.Data == nil {
		rollback := session.AbortTransaction(context)
		result = &pb.ScreenResponseRepeated{
			Code:    int64(codes.Canceled),
			Message: "Screen UseCase ListScreen is failed, data Screen is empty ",
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.ScreenResponseRepeated{
		Code:    int64(codes.OK),
		Message: "Screen UseCase ListScreen is succeed.",
		Data:    ListScreen.Data,
	}
	return result, commit
}
