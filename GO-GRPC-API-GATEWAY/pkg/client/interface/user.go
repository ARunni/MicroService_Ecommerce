package interfaces

import "GO-GRPC-API-GATEWAY/pkg/utils/models"

type UserClient interface {
	UsersSignUp(user models.UserSignUp) (models.TokenUser, error)
	UserLogin(user models.UserLogin) (models.TokenUser, error)
}
