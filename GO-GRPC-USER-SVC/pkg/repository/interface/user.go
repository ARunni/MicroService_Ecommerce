package interfaces

import (
	"GO-GRPC-USER-SVC/pkg/domain"
	"GO-GRPC-USER-SVC/pkg/utils/models"
)

type UserRepository interface {
	CheckUserExistsByEmail(email string) (*domain.User, error)
	CheckUserExistsByPhone(phone string) (*domain.User, error)
	UserSignUp(user models.UserSignUp) (models.UserDetails, error)
	FindUserByEmail(user models.UserLogin) (models.UserDetail, error)
}
