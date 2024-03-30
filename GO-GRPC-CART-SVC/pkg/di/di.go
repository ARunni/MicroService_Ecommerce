package di

import (
	server "GO-GRPC-CART-SVC/pkg/api"
	"GO-GRPC-CART-SVC/pkg/api/service"
	"GO-GRPC-CART-SVC/pkg/client"
	"GO-GRPC-CART-SVC/pkg/config"
	"GO-GRPC-CART-SVC/pkg/db"
	"GO-GRPC-CART-SVC/pkg/repository"
	"GO-GRPC-CART-SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	cartRepository := repository.NewCartRepository(gormDB)
	productClient := client.NewProductClient(&cfg)
	adminUseCase := usecase.NewCartUseCase(cartRepository, productClient)

	adminServiceServer := service.NewCartServer(adminUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, adminServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
