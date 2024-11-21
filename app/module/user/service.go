package user

import (
	"context"
	"github.com/gin-gonic/gin"
	commonflag "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/config/flag"
	commonclientctx "github.com/pixel-plaza-dev/uru-databases-2-go-service-common/http/grpc/client/context"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/protobuf/compiled/user"
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
func (s *Service) SignUp(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.SignUpRequest,
) (*pbuser.SignUpResponse, error) {
	response, err := s.client.SignUp(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// UpdateProfile updates the user's profile
func (s *Service) UpdateProfile(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.UpdateProfileRequest,
) (
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

// GetProfile gets the user's profile
func (s *Service) GetProfile(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetProfileRequest,
) (
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

// GetFullProfile gets the user's full profile
func (s *Service) GetFullProfile(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetFullProfileRequest,
) (
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

// GetUserIdByUsername gets the user's ID by username
func (s *Service) GetUserIdByUsername(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetUserIdByUsernameRequest,
) (
	*pbuser.GetUserIdByUsernameResponse, error,
) {
	response, err := s.client.GetUserIdByUsername(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ChangePassword changes the user's password
func (s *Service) ChangePassword(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangePasswordRequest,
) (
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

// UsernameExists checks if the username exists
func (s *Service) UsernameExists(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.UsernameExistsRequest,
) (
	*pbuser.UsernameExistsResponse, error,
) {
	response, err := s.client.UsernameExists(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetUsernameByUserId gets the username by user ID
func (s *Service) GetUsernameByUserId(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetUsernameByUserIdRequest,
) (
	*pbuser.GetUsernameByUserIdResponse, error,
) {
	response, err := s.client.GetUsernameByUserId(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ChangeUsername changes the users' username
func (s *Service) ChangeUsername(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangeUsernameRequest,
) (
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
func (s *Service) AddEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.AddEmailRequest,
) (*pbuser.AddEmailResponse, error) {
	response, err := s.client.AddEmail(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// DeleteEmail deletes an email from a user
func (s *Service) DeleteEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.DeleteEmailRequest,
) (
	*pbuser.DeleteEmailResponse, error,
) {
	response, err := s.client.DeleteEmail(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// ChangePrimaryEmail changes the user's primary email
func (s *Service) ChangePrimaryEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangePrimaryEmailRequest,
) (
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
func (s *Service) SendVerificationEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.SendVerificationEmailRequest,
) (
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

// VerifyEmail verifies the user's email
func (s *Service) VerifyEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.VerifyEmailRequest,
) (
	*pbuser.VerifyEmailResponse, error,
) {
	response, err := s.client.VerifyEmail(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// GetPrimaryEmail gets the user's primary email
func (s *Service) GetPrimaryEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetPrimaryEmailRequest,
) (
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

// GetActiveEmails gets the user's active emails
func (s *Service) GetActiveEmails(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetActiveEmailsRequest,
) (
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

// ChangePhoneNumber changes the user's phone number
func (s *Service) ChangePhoneNumber(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangePhoneNumberRequest,
) (
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

// SendVerificationSMS sends a verification SMS to a user
func (s *Service) SendVerificationSMS(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.SendVerificationSMSRequest,
) (
	*pbuser.SendVerificationSMSResponse, error,
) {
	response, err := s.client.SendVerificationSMS(
		grpcCtx, request,
	)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}

// VerifyPhoneNumber verifies the user's phone number
func (s *Service) VerifyPhoneNumber(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.VerifyPhoneNumberRequest,
) (
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

// GetPhoneNumber gets the user's phone number
func (s *Service) GetPhoneNumber(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetPhoneNumberRequest,
) (
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
func (s *Service) ForgotPassword(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ForgotPasswordRequest,
) (
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

// ResetPassword resets the user's password
func (s *Service) ResetPassword(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ResetPasswordRequest,
) (
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
func (s *Service) DeleteUser(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.DeleteUserRequest,
) (
	*pbuser.DeleteUserResponse, error,
) {
	response, err := s.client.DeleteUser(grpcCtx, request)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(s.flag, err)
	}
	return response, nil
}
