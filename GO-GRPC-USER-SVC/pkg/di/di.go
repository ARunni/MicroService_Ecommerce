package di

import (
	server "GO-GRPC-USER-SVC/pkg/api"
	"GO-GRPC-USER-SVC/pkg/api/service"
	"GO-GRPC-USER-SVC/pkg/config"
	"GO-GRPC-USER-SVC/pkg/db"
	"GO-GRPC-USER-SVC/pkg/repository"
	"GO-GRPC-USER-SVC/pkg/usecase"
)


func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewUserRepository(gormDB)
	adminUseCase := usecase.NewUserUseCase(adminRepository)

	userServiceServer := service.NewUserServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, userServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
