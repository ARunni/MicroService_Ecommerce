package client

import (
	interfaces "GO-GRPC-API-GATEWAY/pkg/client/interface"
	"GO-GRPC-API-GATEWAY/pkg/config"
	pb "GO-GRPC-API-GATEWAY/pkg/pb/user"
	"GO-GRPC-API-GATEWAY/pkg/utils/models"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

type userClient struct {
	Client pb.UserClient
}

func NewUserClient(cfg config.Config) interfaces.UserClient {

	grpcConnection, err := grpc.Dial(cfg.UserSvcUrl, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Could not connect", err)
	}

	grpcClient := pb.NewUserClient(grpcConnection)

	return &userClient{
		Client: grpcClient,
	}

}
func (c *userClient) UsersSignUp(user models.UserSignUp) (models.TokenUser, error) {
	res, err := c.Client.UserSignUp(context.Background(), &pb.UserSignUpRequest{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		Password:  user.Password,
		Phone:     user.Phone,
	})
	if err != nil {
		return models.TokenUser{}, err
	}
	userDetails := models.UserDetails{
		ID:        uint(res.UserDetails.Id),
		Firstname: res.UserDetails.Firstname,
		Lastname:  res.UserDetails.Lastname,
		Email:     res.UserDetails.Email,
		Phone:     res.UserDetails.Phone,
	}

	return models.TokenUser{
		User:         userDetails,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
func (c *userClient) UserLogin(user models.UserLogin) (models.TokenUser, error) {
	res, err := c.Client.UserLogin(context.Background(), &pb.UserLoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		fmt.Println("🤷‍♂️")
		return models.TokenUser{}, err
	}
	userDetails := models.UserDetails{
		ID:        uint(res.UserDetails.Id),
		Firstname: res.UserDetails.Firstname,
		Lastname:  res.UserDetails.Lastname,
		Email:     res.UserDetails.Email,
		Phone:     res.UserDetails.Phone,
	}

	return models.TokenUser{
		User:         userDetails,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}, nil
}
