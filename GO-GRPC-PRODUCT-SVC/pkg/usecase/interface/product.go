package interfaces

import (
	"GO-GRPC-PRODUCT-SVC/pkg/domain"
	"GO-GRPC-PRODUCT-SVC/pkg/utils/models"
)

type ProductUseCase interface {
	ShowAllProducts(page int, count int) ([]models.ProductBrief, error)
	AddProducts(product models.Product) (domain.Product, error)
	DeleteProducts(id int) error
	UpdateProduct(pid int, stock int) (models.ProductUpdateReciever, error)
	GetQuantityFromProductID(id int) (int, error)
	GetPriceOfProductFromID(prodcut_id int) (float64, error)
	ProductStockMinus(productID, stock int) error
	CheckProduct(product_id int) (bool, error)
}
