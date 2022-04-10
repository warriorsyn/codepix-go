package grpc

import (
	"context"
	"github.com/warriorsyn/codepix-go/application/grpc/pb"
	"github.com/warriorsyn/codepix-go/application/usecase"
)

type PixGrpcService struct {
	PixUseCase usecase.PixUseCase
	pb.UnimplementedPixServiceServer
}

func NewPixGrpcService(usecase usecase.PixUseCase) *PixGrpcService {
	return &PixGrpcService{PixUseCase: usecase}
}

func (p *PixGrpcService) RegisterPixKey(ctx context.Context, in *pb.PixKeyRegistration) (*pb.PixKeyCreatedResult, error) {
	key, err := p.PixUseCase.RegisterKey(in.Key, in.Kind, in.AccountId)

	if err != nil {
		return &pb.PixKeyCreatedResult{
			Status: "not created",
			Error:  err.Error(),
		}, err
	}

	return &pb.PixKeyCreatedResult{
		Id:     key.ID,
		Status: "created",
	}, nil
}

func (p *PixGrpcService) Find(ctx context.Context, in *pb.PixKey) (*pb.PixKeyInfo, error) {
	key, err := p.PixUseCase.FindKey(in.Key, in.Kind)

	if err != nil {
		return &pb.PixKeyInfo{}, err
	}

	return &pb.PixKeyInfo{
		Id:   key.ID,
		Kind: key.Kind,
		Key:  key.Key,
		Account: &pb.Account{
			AccountId:     key.AccountID,
			OwnerName:     key.Account.OwnerName,
			AccountNumber: key.Account.Number,
			BankId:        key.Account.BankID,
			BankName:      key.Account.Bank.Name,
			CreatedAt:     key.Account.CreatedAt.String(),
		},
		CreatedAt: key.CreatedAt.String(),
	}, nil
}
