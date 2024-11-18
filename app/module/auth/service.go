package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc/client/context"
	pbauth "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/auth"
)

// Service is the service for auth
type Service struct {
	client pbauth.AuthClient
	flag   *commonflag.ModeFlag
}

// NewService creates a new service
func NewService(flag *commonflag.ModeFlag, client pbauth.AuthClient) *Service {
	return &Service{client: client, flag: flag}
}

// LogIn logs in the user
func (s *Service) LogIn(ctx *gin.Context, grpcCtx context.Context, request *pbauth.LogInRequest) (*pbauth.LogInResponse, error) {
	response, err := s.client.LogIn(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RefreshToken refreshes the user's token
func (s *Service) RefreshToken(ctx *gin.Context, grpcCtx context.Context, request *pbauth.RefreshTokenRequest) (
	*pbauth.RefreshTokenResponse, error,
) {
	response, err := s.client.RefreshToken(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// LogOut logs out the user
func (s *Service) LogOut(ctx *gin.Context, grpcCtx context.Context, request *pbauth.LogOutRequest) (*pbauth.LogOutResponse, error) {
	response, err := s.client.LogOut(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// CloseSessions closes all the user's sessions
func (s *Service) CloseSessions(ctx *gin.Context, grpcCtx context.Context, request *pbauth.CloseSessionsRequest) (
	*pbauth.CloseSessionsResponse, error,
) {
	response, err := s.client.CloseSessions(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetSessions gets all the user's sessions
func (s *Service) GetSessions(ctx *gin.Context, grpcCtx context.Context, request *pbauth.GetSessionsRequest) (
	*pbauth.GetSessionsResponse, error,
) {
	response, err := s.client.GetSessions(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddPermission adds a permission to the user
func (s *Service) AddPermission(ctx *gin.Context, grpcCtx context.Context, request *pbauth.AddPermissionRequest) (
	*pbauth.AddPermissionResponse, error,
) {
	response, err := s.client.AddPermission(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokePermission revokes a permission from the user
func (s *Service) RevokePermission(ctx *gin.Context, grpcCtx context.Context, request *pbauth.RevokePermissionRequest) (
	*pbauth.RevokePermissionResponse, error,
) {
	response, err := s.client.RevokePermission(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetPermission gets all the user's permissions
func (s *Service) GetPermission(ctx *gin.Context, grpcCtx context.Context, request *pbauth.GetPermissionRequest) (
	*pbauth.GetPermissionResponse, error,
) {
	response, err := s.client.GetPermission(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetPermissions gets all the user's permissions
func (s *Service) GetPermissions(ctx *gin.Context, grpcCtx context.Context, request *pbauth.GetPermissionsRequest) (
	*pbauth.GetPermissionsResponse, error,
) {
	response, err := s.client.GetPermissions(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddRole adds a role to the user
func (s *Service) AddRole(ctx *gin.Context, grpcCtx context.Context, request *pbauth.AddRoleRequest) (*pbauth.AddRoleResponse, error) {
	response, err := s.client.AddRole(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokeRole revokes a role from the user
func (s *Service) RevokeRole(ctx *gin.Context, grpcCtx context.Context, request *pbauth.RevokeRoleRequest) (
	*pbauth.RevokeRoleResponse, error,
) {
	response, err := s.client.RevokeRole(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetRoles gets all the user's roles
func (s *Service) GetRoles(ctx *gin.Context, grpcCtx context.Context, request *pbauth.GetRolesRequest) (*pbauth.GetRolesResponse, error) {
	response, err := s.client.GetRoles(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddRolePermission adds a permission to the role
func (s *Service) AddRolePermission(ctx *gin.Context, grpcCtx context.Context, request *pbauth.AddRolePermissionRequest) (
	*pbauth.AddRolePermissionResponse, error,
) {
	response, err := s.client.AddRolePermission(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokeRolePermission revokes a permission from the role
func (s *Service) RevokeRolePermission(ctx *gin.Context, grpcCtx context.Context, request *pbauth.RevokeRolePermissionRequest) (
	*pbauth.RevokeRolePermissionResponse, error,
) {
	response, err := s.client.RevokeRolePermission(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetRolePermissions gets all the role's permissions
func (s *Service) GetRolePermissions(ctx *gin.Context, grpcCtx context.Context, request *pbauth.GetRolePermissionsRequest) (
	*pbauth.GetRolePermissionsResponse, error,
) {
	response, err := s.client.GetRolePermissions(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddUserRole adds a role to the user
func (s *Service) AddUserRole(ctx *gin.Context, grpcCtx context.Context, request *pbauth.AddUserRoleRequest) (
	*pbauth.AddUserRoleResponse, error,
) {
	response, err := s.client.AddUserRole(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokeUserRole revokes a role from the user
func (s *Service) RevokeUserRole(ctx *gin.Context, grpcCtx context.Context, request *pbauth.RevokeUserRoleRequest) (
	*pbauth.RevokeUserRoleResponse, error,
) {
	response, err := s.client.RevokeUserRole(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetUserRoles gets all the user's roles
func (s *Service) GetUserRoles(ctx *gin.Context, grpcCtx context.Context, request *pbauth.GetUserRolesRequest) (
	*pbauth.GetUserRolesResponse, error,
) {
	response, err := s.client.GetUserRoles(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}
