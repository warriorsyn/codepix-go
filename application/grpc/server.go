package grpc

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/warriorsyn/codepix-go/application/grpc/pb"
	"github.com/warriorsyn/codepix-go/application/usecase"
	"github.com/warriorsyn/codepix-go/infra/repository"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	grpc := grpc.NewServer()

	reflection.Register(grpc)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: &pixRepository}

	pixGrpcService := NewPixGrpcService(pixUseCase)

	pb.RegisterPixServiceServer(grpc, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)

	listener, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatal("Cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started at port %d", port)

	err = grpc.Serve(listener)

	if err != nil {
		log.Fatal("Cannot start grpc server", err)
	}

}
