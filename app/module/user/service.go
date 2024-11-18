package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/server/grpc/client/context"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled-protobuf/user"
)

// Service is the service for user
type Service struct {
	client pbuser.UserClient
	flag   *commonflag.ModeFlag
}

// NewService creates a new service
func NewService(flag *commonflag.ModeFlag, client pbuser.UserClient) *Service {
	return &Service{
		client: client,
		flag:   flag,
	}
}

// SignUp signs up a user
func (s *Service) SignUp(ctx *gin.Context, grpcCtx context.Context, request *pbuser.SignUpRequest) (*pbuser.SignUpResponse, error) {
	response, err := s.client.SignUp(grpcCtx, request)
	fmt.Println(err)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// UpdateProfile updates the profile of a user
func (s *Service) UpdateProfile(ctx *gin.Context, grpcCtx context.Context, request *pbuser.UpdateProfileRequest) (
	*pbuser.UpdateProfileResponse, error,
) {
	response, err := s.client.UpdateProfile(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetProfile gets the profile of a user
func (s *Service) GetProfile(ctx *gin.Context, grpcCtx context.Context, request *pbuser.GetProfileRequest) (
	*pbuser.GetProfileResponse, error,
) {
	response, err := s.client.GetProfile(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetFullProfile gets the full profile of a user
func (s *Service) GetFullProfile(ctx *gin.Context, grpcCtx context.Context, request *pbuser.GetFullProfileRequest) (
	*pbuser.GetFullProfileResponse, error,
) {
	response, err := s.client.GetFullProfile(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ChangePassword changes the password of a user
func (s *Service) ChangePassword(ctx *gin.Context, grpcCtx context.Context, request *pbuser.ChangePasswordRequest) (
	*pbuser.ChangePasswordResponse, error,
) {
	response, err := s.client.ChangePassword(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ChangeUsername changes the username of a user
func (s *Service) ChangeUsername(ctx *gin.Context, grpcCtx context.Context, request *pbuser.ChangeUsernameRequest) (
	*pbuser.ChangeUsernameResponse, error,
) {
	response, err := s.client.ChangeUsername(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// AddEmail adds an email to a user
func (s *Service) AddEmail(ctx *gin.Context, grpcCtx context.Context, request *pbuser.AddEmailRequest) (*pbuser.AddEmailResponse, error) {
	response, err := s.client.AddEmail(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// DeleteEmail deletes an email from a user
func (s *Service) DeleteEmail(ctx *gin.Context, grpcCtx context.Context, request *pbuser.DeleteEmailRequest) (
	*pbuser.DeleteEmailResponse, error,
) {
	response, err := s.client.DeleteEmail(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ChangePrimaryEmail changes the primary email of a user
func (s *Service) ChangePrimaryEmail(ctx *gin.Context, grpcCtx context.Context, request *pbuser.ChangePrimaryEmailRequest) (
	*pbuser.ChangePrimaryEmailResponse, error,
) {
	response, err := s.client.ChangePrimaryEmail(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// SendVerificationEmail sends a verification email to a user
func (s *Service) SendVerificationEmail(ctx *gin.Context, grpcCtx context.Context, request *pbuser.SendVerificationEmailRequest) (
	*pbuser.SendVerificationEmailResponse, error,
) {
	response, err := s.client.SendVerificationEmail(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// VerifyEmail verifies the email of a user
func (s *Service) VerifyEmail(ctx *gin.Context, grpcCtx context.Context, request *pbuser.VerifyEmailRequest) (
	*pbuser.VerifyEmailResponse, error,
) {
	response, err := s.client.VerifyEmail(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetPrimaryEmail gets the primary email of a user
func (s *Service) GetPrimaryEmail(ctx *gin.Context, grpcCtx context.Context, request *pbuser.GetPrimaryEmailRequest) (
	*pbuser.GetPrimaryEmailResponse, error,
) {
	response, err := s.client.GetPrimaryEmail(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetActiveEmails gets the active emails of a user
func (s *Service) GetActiveEmails(ctx *gin.Context, grpcCtx context.Context, request *pbuser.GetActiveEmailsRequest) (
	*pbuser.GetActiveEmailsResponse, error,
) {
	response, err := s.client.GetActiveEmails(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ChangePhoneNumber changes the phone number of a user
func (s *Service) ChangePhoneNumber(ctx *gin.Context, grpcCtx context.Context, request *pbuser.ChangePhoneNumberRequest) (
	*pbuser.ChangePhoneNumberResponse, error,
) {
	response, err := s.client.ChangePhoneNumber(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// SendVerificationPhoneNumber sends a verification phone number to a user
func (s *Service) SendVerificationPhoneNumber(ctx *gin.Context, grpcCtx context.Context, request *pbuser.SendVerificationPhoneNumberRequest) (
	*pbuser.SendVerificationPhoneNumberResponse, error,
) {
	response, err := s.client.SendVerificationPhoneNumber(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// VerifyPhoneNumber verifies the phone number of a user
func (s *Service) VerifyPhoneNumber(ctx *gin.Context, grpcCtx context.Context, request *pbuser.VerifyPhoneNumberRequest) (
	*pbuser.VerifyPhoneNumberResponse, error,
) {
	response, err := s.client.VerifyPhoneNumber(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetPhoneNumber gets the phone number of a user
func (s *Service) GetPhoneNumber(ctx *gin.Context, grpcCtx context.Context, request *pbuser.GetPhoneNumberRequest) (
	*pbuser.GetPhoneNumberResponse, error,
) {
	response, err := s.client.GetPhoneNumber(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ForgotPassword sends a forgot password email to a user
func (s *Service) ForgotPassword(ctx *gin.Context, grpcCtx context.Context, request *pbuser.ForgotPasswordRequest) (
	*pbuser.ForgotPasswordResponse, error,
) {
	response, err := s.client.ForgotPassword(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ResetPassword resets the password of a user
func (s *Service) ResetPassword(ctx *gin.Context, grpcCtx context.Context, request *pbuser.ResetPasswordRequest) (
	*pbuser.ResetPasswordResponse, error,
) {
	response, err := s.client.ResetPassword(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// DeleteUser deletes a user
func (s *Service) DeleteUser(ctx *gin.Context, grpcCtx context.Context, request *pbuser.DeleteUserRequest) (
	*pbuser.DeleteUserResponse, error,
) {
	response, err := s.client.DeleteUser(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}
