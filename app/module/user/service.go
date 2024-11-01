package user

import (
	"github.com/gin-gonic/gin"
	commonclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/grpc/client/context"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/user"
)

// Service is the service for user
type Service struct {
	client pbuser.UserClient
}

// NewService creates a new service
func NewService(client pbuser.UserClient) *Service {
	return &Service{
		client: client,
	}
}

// SignUp signs up a user
func (s *Service) SignUp(ctx *gin.Context) (*pbuser.SignUpResponse, error) {
	var signUpRequest pbuser.SignUpRequest

	// Bind the request
	if err := ctx.BindJSON(&signUpRequest); err != nil {
		return nil, err
	}

	// Call the client
	signUpResponse, err := s.client.SignUp(ctx, &signUpRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return signUpResponse, nil
}
