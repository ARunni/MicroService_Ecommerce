package interfaces

import (
	"GO-GRPC-USER-SVC/pkg/domain"
	"GO-GRPC-USER-SVC/pkg/utils/models"
)

type UserUseCase interface {
	UsersSignUp(user models.UserSignUp) (domain.TokenUser, error)
	UsersLogin(user models.UserLogin) (domain.TokenUser, error)
}
