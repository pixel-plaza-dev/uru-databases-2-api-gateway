package auth

import (
	"github.com/gin-gonic/gin"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/flag"
	commonclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/grpc/client/context"
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
func (s *Service) LogIn(ctx *gin.Context) (*pbauth.LogInResponse, error) {
	var request pbauth.LogInRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Log in the user
	response, err := s.client.LogIn(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RefreshToken refreshes the user's token
func (s *Service) RefreshToken(ctx *gin.Context) (*pbauth.RefreshTokenResponse, error) {
	var request pbauth.RefreshTokenRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Refresh the token
	response, err := s.client.RefreshToken(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// LogOut logs out the user
func (s *Service) LogOut(ctx *gin.Context) (*pbauth.LogOutResponse, error) {
	var request pbauth.LogOutRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Log out the user
	response, err := s.client.LogOut(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// CloseSessions closes all the user's sessions
func (s *Service) CloseSessions(ctx *gin.Context) (*pbauth.CloseSessionsResponse, error) {
	var request pbauth.CloseSessionsRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Close all the user's sessions
	response, err := s.client.CloseSessions(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetSessions gets all the user's sessions
func (s *Service) GetSessions(ctx *gin.Context) (*pbauth.GetSessionsResponse, error) {
	var request pbauth.GetSessionsRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Get all the user's sessions
	response, err := s.client.GetSessions(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddPermission adds a permission to the user
func (s *Service) AddPermission(ctx *gin.Context) (*pbauth.AddPermissionResponse, error) {
	var request pbauth.AddPermissionRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Add a permission to the user
	response, err := s.client.AddPermission(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokePermission revokes a permission from the user
func (s *Service) RevokePermission(ctx *gin.Context) (*pbauth.RevokePermissionResponse, error) {
	var request pbauth.RevokePermissionRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Revoke a permission from the user
	response, err := s.client.RevokePermission(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetPermission gets all the user's permissions
func (s *Service) GetPermission(ctx *gin.Context) (*pbauth.GetPermissionResponse, error) {
	var request pbauth.GetPermissionRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Get all the user's permissions
	response, err := s.client.GetPermission(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetPermissions gets all the user's permissions
func (s *Service) GetPermissions(ctx *gin.Context) (*pbauth.GetPermissionsResponse, error) {
	var request pbauth.GetPermissionsRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Get all the user's permissions
	response, err := s.client.GetPermissions(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddRole adds a role to the user
func (s *Service) AddRole(ctx *gin.Context) (*pbauth.AddRoleResponse, error) {
	var request pbauth.AddRoleRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Add a role to the user
	response, err := s.client.AddRole(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokeRole revokes a role from the user
func (s *Service) RevokeRole(ctx *gin.Context) (*pbauth.RevokeRoleResponse, error) {
	var request pbauth.RevokeRoleRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Revoke a role from the user
	response, err := s.client.RevokeRole(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetRoles gets all the user's roles
func (s *Service) GetRoles(ctx *gin.Context) (*pbauth.GetRolesResponse, error) {
	var request pbauth.GetRolesRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Get all the user's roles
	response, err := s.client.GetRoles(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddRolePermission adds a permission to the role
func (s *Service) AddRolePermission(ctx *gin.Context) (*pbauth.AddRolePermissionResponse, error) {
	var request pbauth.AddRolePermissionRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Add a permission to the role
	response, err := s.client.AddRolePermission(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokeRolePermission revokes a permission from the role
func (s *Service) RevokeRolePermission(ctx *gin.Context) (*pbauth.RevokeRolePermissionResponse, error) {
	var request pbauth.RevokeRolePermissionRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Revoke a permission from the role
	response, err := s.client.RevokeRolePermission(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetRolePermissions gets all the role's permissions
func (s *Service) GetRolePermissions(ctx *gin.Context) (*pbauth.GetRolePermissionsResponse, error) {
	var request pbauth.GetRolePermissionsRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Get all the role's permissions
	response, err := s.client.GetRolePermissions(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddUserRole adds a role to the user
func (s *Service) AddUserRole(ctx *gin.Context) (*pbauth.AddUserRoleResponse, error) {
	var request pbauth.AddUserRoleRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Add a role to the user
	response, err := s.client.AddUserRole(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// RevokeUserRole revokes a role from the user
func (s *Service) RevokeUserRole(ctx *gin.Context) (*pbauth.RevokeUserRoleResponse, error) {
	var request pbauth.RevokeUserRoleRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Revoke a role from the user
	response, err := s.client.RevokeUserRole(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetUserRoles gets all the user's roles
func (s *Service) GetUserRoles(ctx *gin.Context) (*pbauth.GetUserRolesResponse, error) {
	var request pbauth.GetUserRolesRequest

	// Bind the request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return nil, err
	}

	// Get all the user's roles
	response, err := s.client.GetUserRoles(ctx, &request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}
