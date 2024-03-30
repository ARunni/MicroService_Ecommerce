package di

import (
	server "GO-GRPC-PRODUCT-SVC/pkg/api"
	"GO-GRPC-PRODUCT-SVC/pkg/api/service"
	"GO-GRPC-PRODUCT-SVC/pkg/config"
	"GO-GRPC-PRODUCT-SVC/pkg/db"
	"GO-GRPC-PRODUCT-SVC/pkg/repository"
	"GO-GRPC-PRODUCT-SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	adminRepository := repository.NewProductRepository(gormDB)
	adminUseCase := usecase.NewProductUseCase(adminRepository)

	adminServiceServer := service.NewProductServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}