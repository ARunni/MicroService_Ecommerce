package interfaces

import (
	"GO-GRPC-ORDER-SVC/pkg/domain"
	"GO-GRPC-ORDER-SVC/pkg/utils/models"
)

type OrderUseCase interface {
	OrderItemsFromCart(orderFromCart models.OrderFromCart, userID int) (domain.OrderSuccessResponse, error)
	GetOrderDetails(userId int, page int, count int) ([]models.FullOrderDetails, error)
}
