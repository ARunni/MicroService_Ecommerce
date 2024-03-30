package interfaces

import (
	"GO-GRPC-ADMIN-SVC/pkg/domain"
	"GO-GRPC-ADMIN-SVC/pkg/utils/models"
)

type AdminRepository interface {
	AdminSignUp(adminDetails models.AdminSignUp) (models.AdminDetailsResponse, error)
	FindAdminByEmail(admin models.AdminLogin) (models.AdminSignUp, error)
	CheckAdminExistsByEmail(email string) (*domain.Admin, error)
}
