package di

import (
	server "GO-GRPC-API-GATEWAY/pkg/api"
	"GO-GRPC-API-GATEWAY/pkg/api/handler"
	"GO-GRPC-API-GATEWAY/pkg/client"
	"GO-GRPC-API-GATEWAY/pkg/config"
)

func InitializeAPI(cfg config.Config) (*server.ServerHTTP, error) {

	adminClient := client.NewAdminClient(cfg)
	adminHandler := handler.NewAdminHandler(adminClient)

	productClient := client.NewProductClient(cfg)
	productHandler := handler.NewProductHandler(productClient)

	userClient := client.NewUserClient(cfg)
	userHandler := handler.NewUserHandler(userClient)

	cartClient := client.NewCartClient(cfg)
	cartHandler := handler.NewCartHandler(cartClient)

	orderClient := client.NewOrderClient(cfg)
	orderHandler := handler.NewOrderHandler(orderClient)

	serverHTTP := server.NewServerHTTP(adminHandler, productHandler, userHandler, cartHandler, orderHandler)

	return serverHTTP, nil
}
