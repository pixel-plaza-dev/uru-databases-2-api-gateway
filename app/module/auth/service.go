package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	commonclientrequest "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/request"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/auth"
)

// Service is the service for auth
type Service struct {
	client  pbauth.AuthClient
	handler commonclientrequest.Handler
}

// NewService creates a new service
func NewService(
	client pbauth.AuthClient,
	handler commonclientrequest.Handler,
) *Service {
	return &Service{client: client, handler: handler}
}

// LogIn logs in q user
func (s *Service) LogIn(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.LogInRequest,
) (*pbauth.LogInResponse, error) {
	response, err := s.client.LogIn(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// IsAccessTokenValid checks if the access token is valid
func (s *Service) IsAccessTokenValid(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.IsAccessTokenValidRequest,
) (
	*pbauth.IsAccessTokenValidResponse, error,
) {
	response, err := s.client.IsAccessTokenValid(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// IsRefreshTokenValid checks if the refresh token is valid
func (s *Service) IsRefreshTokenValid(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.IsRefreshTokenValidRequest,
) (
	*pbauth.IsRefreshTokenValidResponse, error,
) {
	response, err := s.client.IsRefreshTokenValid(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// RefreshToken refreshes the user's token
func (s *Service) RefreshToken(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RefreshTokenRequest,
) (
	*pbauth.RefreshTokenResponse, error,
) {
	response, err := s.client.RefreshToken(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// LogOut logs out the user
func (s *Service) LogOut(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.LogOutRequest,
) (*pbauth.LogOutResponse, error) {
	response, err := s.client.LogOut(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// GetRefreshTokenInformation gets the refresh token information
func (s *Service) GetRefreshTokenInformation(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetRefreshTokenInformationRequest,
) (
	response *pbauth.GetRefreshTokenInformationResponse, err error,
) {
	response, err = s.client.GetRefreshTokenInformation(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// GetRefreshTokensInformation gets all refresh tokens information
func (s *Service) GetRefreshTokensInformation(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetRefreshTokensInformationRequest,
) (
	*pbauth.GetRefreshTokensInformationResponse, error,
) {
	response, err := s.client.GetRefreshTokensInformation(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// RevokeRefreshToken revokes a user's refresh token
func (s *Service) RevokeRefreshToken(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeRefreshTokenRequest,
) (
	*pbauth.RevokeRefreshTokenResponse, error,
) {
	response, err := s.client.RevokeRefreshToken(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// RevokeRefreshTokens revokes all the user's refresh tokens
func (s *Service) RevokeRefreshTokens(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeRefreshTokensRequest,
) (
	*pbauth.RevokeRefreshTokensResponse, error,
) {
	response, err := s.client.RevokeRefreshTokens(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// AddPermission adds a permission
func (s *Service) AddPermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddPermissionRequest,
) (
	*pbauth.AddPermissionResponse, error,
) {
	response, err := s.client.AddPermission(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// RevokePermission revokes a permission
func (s *Service) RevokePermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokePermissionRequest,
) (
	*pbauth.RevokePermissionResponse, error,
) {
	response, err := s.client.RevokePermission(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// GetPermission gets all the permissions
func (s *Service) GetPermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetPermissionRequest,
) (
	*pbauth.GetPermissionResponse, error,
) {
	response, err := s.client.GetPermission(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// GetPermissions gets all the permissions
func (s *Service) GetPermissions(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetPermissionsRequest,
) (
	*pbauth.GetPermissionsResponse, error,
) {
	response, err := s.client.GetPermissions(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// AddRolePermission adds a permission to a role
func (s *Service) AddRolePermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddRolePermissionRequest,
) (
	*pbauth.AddRolePermissionResponse, error,
) {
	response, err := s.client.AddRolePermission(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// RevokeRolePermission revokes a permission from a role
func (s *Service) RevokeRolePermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeRolePermissionRequest,
) (
	*pbauth.RevokeRolePermissionResponse, error,
) {
	response, err := s.client.RevokeRolePermission(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// GetRolePermissions gets all the role's permissions
func (s *Service) GetRolePermissions(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetRolePermissionsRequest,
) (
	*pbauth.GetRolePermissionsResponse, error,
) {
	response, err := s.client.GetRolePermissions(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// AddRole adds a role
func (s *Service) AddRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddRoleRequest,
) (*pbauth.AddRoleResponse, error) {
	response, err := s.client.AddRole(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// RevokeRole revokes a role
func (s *Service) RevokeRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeRoleRequest,
) (
	*pbauth.RevokeRoleResponse, error,
) {
	response, err := s.client.RevokeRole(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// GetRoles gets all the roles
func (s *Service) GetRoles(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetRolesRequest,
) (*pbauth.GetRolesResponse, error) {
	response, err := s.client.GetRoles(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// AddUserRole adds a role to a user
func (s *Service) AddUserRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddUserRoleRequest,
) (
	*pbauth.AddUserRoleResponse, error,
) {
	response, err := s.client.AddUserRole(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// RevokeUserRole revokes a role from a user
func (s *Service) RevokeUserRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeUserRoleRequest,
) (
	*pbauth.RevokeUserRoleResponse, error,
) {
	response, err := s.client.RevokeUserRole(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}

// GetUserRoles gets all the user's roles
func (s *Service) GetUserRoles(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetUserRolesRequest,
) (
	*pbauth.GetUserRolesResponse, error,
) {
	response, err := s.client.GetUserRoles(grpcCtx, request)
	if err != nil {
		return nil, s.handler.HandleError(err)
	}
	return response, nil
}
