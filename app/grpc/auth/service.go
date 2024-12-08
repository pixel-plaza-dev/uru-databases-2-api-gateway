package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/auth"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Service is the service for auth
type Service struct {
	client          pbauth.AuthClient
	responseHandler commonclientresponse.Handler
}

// NewService creates a new service
func NewService(
	client pbauth.AuthClient,
	responseHandler commonclientresponse.Handler,
) *Service {
	return &Service{client: client, responseHandler: responseHandler}
}

// LogIn logs in q user
func (s *Service) LogIn(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.LogInRequest,
) (*pbauth.LogInResponse, error) {
	return s.client.LogIn(grpcCtx, request)
}

// IsAccessTokenValid checks if the access token is valid
func (s *Service) IsAccessTokenValid(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.IsAccessTokenValidRequest,
) (
	*pbauth.IsAccessTokenValidResponse, error,
) {
	return s.client.IsAccessTokenValid(grpcCtx, request)
}

// IsRefreshTokenValid checks if the refresh token is valid
func (s *Service) IsRefreshTokenValid(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.IsRefreshTokenValidRequest,
) (
	*pbauth.IsRefreshTokenValidResponse, error,
) {
	return s.client.IsRefreshTokenValid(grpcCtx, request)
}

// RefreshToken refreshes the user's token
func (s *Service) RefreshToken(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbauth.RefreshTokenResponse, error,
) {
	return s.client.RefreshToken(grpcCtx, &emptypb.Empty{})
}

// LogOut logs out the user
func (s *Service) LogOut(
	ctx *gin.Context,
	grpcCtx context.Context,
) (*pbauth.LogOutResponse, error) {
	return s.client.LogOut(grpcCtx, &emptypb.Empty{})
}

// GetRefreshTokenInformation gets the refresh token information
func (s *Service) GetRefreshTokenInformation(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetRefreshTokenInformationRequest,
) (
	*pbauth.GetRefreshTokenInformationResponse, error,
) {
	return s.client.GetRefreshTokenInformation(grpcCtx, request)
}

// GetRefreshTokensInformation gets all refresh tokens information
func (s *Service) GetRefreshTokensInformation(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbauth.GetRefreshTokensInformationResponse, error,
) {
	return s.client.GetRefreshTokensInformation(grpcCtx, &emptypb.Empty{})
}

// RevokeRefreshToken revokes a user's refresh token
func (s *Service) RevokeRefreshToken(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeRefreshTokenRequest,
) (
	*pbauth.RevokeRefreshTokenResponse, error,
) {
	return s.client.RevokeRefreshToken(grpcCtx, request)
}

// RevokeRefreshTokens revokes all the user's refresh tokens
func (s *Service) RevokeRefreshTokens(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbauth.RevokeRefreshTokensResponse, error,
) {
	return s.client.RevokeRefreshTokens(grpcCtx, &emptypb.Empty{})
}

// AddPermission adds a permission
func (s *Service) AddPermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddPermissionRequest,
) (
	*pbauth.AddPermissionResponse, error,
) {
	return s.client.AddPermission(grpcCtx, request)
}

// RevokePermission revokes a permission
func (s *Service) RevokePermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokePermissionRequest,
) (
	*pbauth.RevokePermissionResponse, error,
) {
	return s.client.RevokePermission(grpcCtx, request)
}

// GetPermission gets all the permissions
func (s *Service) GetPermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetPermissionRequest,
) (
	*pbauth.GetPermissionResponse, error,
) {
	return s.client.GetPermission(grpcCtx, request)
}

// GetPermissions gets all the permissions
func (s *Service) GetPermissions(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbauth.GetPermissionsResponse, error,
) {
	return s.client.GetPermissions(grpcCtx, &emptypb.Empty{})
}

// AddRolePermission adds a permission to a role
func (s *Service) AddRolePermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddRolePermissionRequest,
) (
	*pbauth.AddRolePermissionResponse, error,
) {
	return s.client.AddRolePermission(grpcCtx, request)
}

// RevokeRolePermission revokes a permission from a role
func (s *Service) RevokeRolePermission(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeRolePermissionRequest,
) (
	*pbauth.RevokeRolePermissionResponse, error,
) {
	return s.client.RevokeRolePermission(grpcCtx, request)
}

// GetRolePermissions gets all the role's permissions
func (s *Service) GetRolePermissions(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetRolePermissionsRequest,
) (
	*pbauth.GetRolePermissionsResponse, error,
) {
	return s.client.GetRolePermissions(grpcCtx, request)
}

// AddRole adds a role
func (s *Service) AddRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddRoleRequest,
) (*pbauth.AddRoleResponse, error) {
	return s.client.AddRole(grpcCtx, request)
}

// RevokeRole revokes a role
func (s *Service) RevokeRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeRoleRequest,
) (
	*pbauth.RevokeRoleResponse, error,
) {
	return s.client.RevokeRole(grpcCtx, request)
}

// GetRoles gets all the roles
func (s *Service) GetRoles(
	ctx *gin.Context,
	grpcCtx context.Context,
) (*pbauth.GetRolesResponse, error) {
	return s.client.GetRoles(grpcCtx, &emptypb.Empty{})
}

// AddUserRole adds a role to a user
func (s *Service) AddUserRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.AddUserRoleRequest,
) (
	*pbauth.AddUserRoleResponse, error,
) {
	return s.client.AddUserRole(grpcCtx, request)
}

// RevokeUserRole revokes a role from a user
func (s *Service) RevokeUserRole(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.RevokeUserRoleRequest,
) (
	*pbauth.RevokeUserRoleResponse, error,
) {
	return s.client.RevokeUserRole(grpcCtx, request)
}

// GetUserRoles gets all the user's roles
func (s *Service) GetUserRoles(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbauth.GetUserRolesRequest,
) (
	*pbauth.GetUserRolesResponse, error,
) {
	return s.client.GetUserRoles(grpcCtx, request)
}
