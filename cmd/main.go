package main

import (
	"github.com/jinzhu/gorm"
	"github.com/warriorsyn/codepix-go/application/grpc"
	"github.com/warriorsyn/codepix-go/infra/db"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))

	grpc.StartGrpcServer(database, 50051)
}
