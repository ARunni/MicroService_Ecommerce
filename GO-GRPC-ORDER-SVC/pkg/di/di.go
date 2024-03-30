package di

import (
	server "GO-GRPC-ORDER-SVC/pkg/api"
	"GO-GRPC-ORDER-SVC/pkg/api/service"
	"GO-GRPC-ORDER-SVC/pkg/client"
	"GO-GRPC-ORDER-SVC/pkg/config"
	"GO-GRPC-ORDER-SVC/pkg/db"
	"GO-GRPC-ORDER-SVC/pkg/repository"
	"GO-GRPC-ORDER-SVC/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*server.Server, error) {

	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	orderRepository := repository.NewOrderRepository(gormDB)
	cartClient := client.NewCartClient(&cfg)
	productClient := client.NewProductClient(&cfg)
	orderUseCase := usecase.NewOrderUseCase(orderRepository, cartClient, productClient)

	orderServiceServer := service.NewOrderServer(orderUseCase)
	grpcServer, err := server.NewGRPCServer(cfg, orderServiceServer)

	if err != nil {
		return &server.Server{}, err
	}

	return grpcServer, nil

}
