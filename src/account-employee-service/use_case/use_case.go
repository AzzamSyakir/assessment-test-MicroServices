package use_case

import (
	"assesement-test-MicroServices/grpc/pb"
	"assesement-test-MicroServices/src/account-employee-service/config"
	"assesement-test-MicroServices/src/account-employee-service/repository"
	"context"
	"fmt"
	"time"

	"github.com/guregu/null"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountUseCase struct {
	pb.UnimplementedAccountServiceServer
	DatabaseConfig    *config.DatabaseConfig
	AccountRepository *repository.AccountRepository
}

func NewAccountUseCase(
	databaseConfig *config.DatabaseConfig,
	AccountRepository *repository.AccountRepository,
) *AccountUseCase {
	return &AccountUseCase{
		UnimplementedAccountServiceServer: pb.UnimplementedAccountServiceServer{},
		DatabaseConfig:                    databaseConfig,
		AccountRepository:                 AccountRepository,
	}
}

func (AccountUseCase *AccountUseCase) GetAccountById(context context.Context, id *pb.ById) (result *pb.AccountResponse, err error) {
	session, err := AccountUseCase.DatabaseConfig.AccountDB.Connection.StartSession()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase GetAccountById is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase GetAccountById is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	GetAccountById, GetAccountByIdErr := AccountUseCase.AccountRepository.GetAccountById(AccountUseCase.DatabaseConfig.AccountDB.Connection, id.Id)
	if GetAccountByIdErr != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("AccountUseCase GetAccountById is failed, GetAccountById failed : %s", GetAccountByIdErr)
		result = &pb.AccountResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	if GetAccountById == nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("Account UseCase GetOneById is failed, Account is not found by id %s", id)
		result = &pb.AccountResponse{
			Code:    int64(codes.Canceled),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.AccountResponse{
		Code:    int64(codes.OK),
		Message: "Account UseCase GetOneById is succeed.",
		Data:    GetAccountById,
	}
	return result, commit
}

func (AccountUseCase *AccountUseCase) UpdateAccount(context context.Context, request *pb.UpdateAccountRequest) (result *pb.AccountResponse, err error) {
	begin := AccountUseCase.DatabaseConfig.AccountDB.Connection
	session, err := AccountUseCase.DatabaseConfig.AccountDB.Connection.StartSession()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase UpdateAccount is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase UpdateAccount is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	foundAccount, err := AccountUseCase.AccountRepository.GetAccountById(begin, request.Id)
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Canceled),
			Message: "AccountUseCase UpdateAccount is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if foundAccount == nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Canceled),
			Message: "AccountAccountCase UpdateAccount is failed, Account is not found by id " + request.Id,
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	if request.Name != nil {
		foundAccount.AccountName = *request.Name
	}
	if request.Password != nil {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*request.Password), bcrypt.DefaultCost)
		if err != nil {
			result = &pb.AccountResponse{
				Code:    int64(codes.Canceled),
				Message: "AccountUseCase UpdateAccount is failed, password hashing is failed, " + err.Error(),
				Data:    nil,
			}
			return result, session.AbortTransaction(context)
		}

		foundAccount.Password = string(hashedPassword)
	}
	time := time.Now()
	foundAccount.UpdatedAt = timestamppb.New(time)
	patchedAccount, err := AccountUseCase.AccountRepository.PatchOneById(begin, request.Id, foundAccount)
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase UpdateAccount is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	result = &pb.AccountResponse{
		Code:    int64(codes.OK),
		Message: "AccountAccountCase UpdateAccount is succeed.",
		Data:    patchedAccount,
	}
	return result, session.CommitTransaction(context)
}
func (AccountUseCase *AccountUseCase) CreateAccount(context context.Context, request *pb.CreateAccountRequest) (result *pb.AccountResponse, err error) {
	session, err := AccountUseCase.DatabaseConfig.AccountDB.Connection.StartSession()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase Register is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase Register is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}

	hashedPassword, hashedPasswordErr := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if hashedPasswordErr != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Canceled),
			Message: "AccountUseCase Register is failed, password hashing is failed.",
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}

	currentTime := null.NewTime(time.Now(), true)
	newAccount := &pb.Account{
		AccountName: request.Name,
		Password:    string(hashedPassword),
		CreatedAt:   timestamppb.New(currentTime.Time),
		UpdatedAt:   timestamppb.New(currentTime.Time),
	}
	createdAccount, err := AccountUseCase.AccountRepository.CreateAccount(AccountUseCase.DatabaseConfig.AccountDB.Connection, newAccount)
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase Register is failed, query to db fail, " + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}

	result = &pb.AccountResponse{
		Code:    int64(codes.OK),
		Message: "AccountUseCase Register is succeed.",
		Data:    createdAccount,
	}
	return result, session.CommitTransaction(context)
}
func (AccountUseCase *AccountUseCase) DeleteAccount(context context.Context, id *pb.ById) (result *pb.AccountResponse, err error) {
	session, err := AccountUseCase.DatabaseConfig.AccountDB.Connection.StartSession()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase DeleteAccount is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase DeleteAccount is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	deletedAccount, deletedAccountErr := AccountUseCase.AccountRepository.DeleteAccount(AccountUseCase.DatabaseConfig.AccountDB.Connection, id.Id)
	if deletedAccountErr != nil {
		err = session.AbortTransaction(context)
		result = &pb.AccountResponse{
			Code:    int64(codes.Internal),
			Message: "AccountAccountCase DeleteAccount is failed, " + deletedAccountErr.Error(),
			Data:    nil,
		}
		return result, err
	}
	if deletedAccount == nil {
		err = session.AbortTransaction(context)
		result = &pb.AccountResponse{
			Code:    int64(codes.Canceled),
			Message: "AccountAccountCase DeleteAccount is failed, Account is not deleted by id, " + id.Id,
			Data:    nil,
		}
		return result, err
	}

	err = session.CommitTransaction(context)
	result = &pb.AccountResponse{
		Code:    int64(codes.OK),
		Message: "AccountAccountCase DeleteAccount is succeed.",
		Data:    deletedAccount,
	}
	return result, err
}
func (AccountUseCase *AccountUseCase) ListAccounts(context context.Context, empty *pb.Empty) (result *pb.AccountResponseRepeated, err error) {
	session, err := AccountUseCase.DatabaseConfig.AccountDB.Connection.StartSession()
	if err != nil {
		result = &pb.AccountResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase ListAccount is failed, startSession fail," + err.Error(),
			Data:    nil,
		}
		return result, session.AbortTransaction(context)
	}
	err = session.StartTransaction()
	if err != nil {
		result = &pb.AccountResponseRepeated{
			Code:    int64(codes.Internal),
			Message: "AccountUseCase ListAccount is failed, StartTransaction fail," + err.Error(),
			Data:    nil,
		}
		return result, nil
	}
	ListAccount, err := AccountUseCase.AccountRepository.ListAccount(AccountUseCase.DatabaseConfig.AccountDB.Connection)
	if err != nil {
		rollback := session.AbortTransaction(context)
		errorMessage := fmt.Sprintf("AccountUseCase ListAccount is failed, query failed : %s", err)
		result = &pb.AccountResponseRepeated{
			Code:    int64(codes.Internal),
			Message: errorMessage,
			Data:    nil,
		}
		return result, rollback
	}

	if ListAccount.Data == nil {
		rollback := session.AbortTransaction(context)
		result = &pb.AccountResponseRepeated{
			Code:    int64(codes.Canceled),
			Message: "Account UseCase ListAccount is failed, data Account is empty ",
			Data:    nil,
		}
		return result, rollback
	}
	commit := session.CommitTransaction(context)
	result = &pb.AccountResponseRepeated{
		Code:    int64(codes.OK),
		Message: "Account UseCase ListAccount is succeed.",
		Data:    ListAccount.Data,
	}
	return result, commit
}
