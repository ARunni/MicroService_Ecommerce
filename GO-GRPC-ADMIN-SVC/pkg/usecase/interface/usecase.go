package interfaces

import (
	"GO-GRPC-ADMIN-SVC/pkg/domain"
	"GO-GRPC-ADMIN-SVC/pkg/utils/models"
)

type AdminUseCase interface {
	AdminSignUp(admindeatils models.AdminSignUp) (*domain.TokenAdmin, error)
	LoginHandler(adminDetails models.AdminLogin) (*domain.TokenAdmin, error)
}