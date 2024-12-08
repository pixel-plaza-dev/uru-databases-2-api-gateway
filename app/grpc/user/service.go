package user

import (
	"context"
	"github.com/gin-gonic/gin"
	commonclientresponse "github.com/pixel-plaza-dev/uru-databases-2-go-api-common/http/grpc/client/response"
	pbuser "github.com/pixel-plaza-dev/uru-databases-2-protobuf-common/compiled/pixel_plaza/user"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Service is the service for user
type Service struct {
	client          pbuser.UserClient
	responseHandler commonclientresponse.Handler
}

// NewService creates a new service
func NewService(client pbuser.UserClient, responseHandler commonclientresponse.Handler) *Service {
	return &Service{
		client:          client,
		responseHandler: responseHandler,
	}
}

// SignUp signs up a user
func (s *Service) SignUp(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.SignUpRequest,
) (*pbuser.SignUpResponse, error) {
	return s.client.SignUp(grpcCtx, request)
}

// UpdateUser updates the user
func (s *Service) UpdateUser(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.UpdateUserRequest,
) (
	*pbuser.UpdateUserResponse, error,
) {
	return s.client.UpdateUser(
		grpcCtx, request,
	)
}

// GetProfile gets the user's profile
func (s *Service) GetProfile(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetProfileRequest,
) (
	*pbuser.GetProfileResponse, error,
) {
	return s.client.GetProfile(
		grpcCtx, request,
	)
}

// GetMyProfile gets the user's profile
func (s *Service) GetMyProfile(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbuser.GetMyProfileResponse, error,
) {
	return s.client.GetMyProfile(
		grpcCtx, &emptypb.Empty{},
	)
}

// GetUserIdByUsername gets the user's ID by username
func (s *Service) GetUserIdByUsername(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetUserIdByUsernameRequest,
) (
	*pbuser.GetUserIdByUsernameResponse, error,
) {
	return s.client.GetUserIdByUsername(
		grpcCtx, request,
	)
}

// ChangePassword changes the user's password
func (s *Service) ChangePassword(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangePasswordRequest,
) (
	*pbuser.ChangePasswordResponse, error,
) {
	return s.client.ChangePassword(
		grpcCtx, request,
	)
}

// UsernameExists checks if the username exists
func (s *Service) UsernameExists(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.UsernameExistsRequest,
) (
	*pbuser.UsernameExistsResponse, error,
) {
	return s.client.UsernameExists(
		grpcCtx, request,
	)
}

// GetUsernameByUserId gets the username by user ID
func (s *Service) GetUsernameByUserId(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.GetUsernameByUserIdRequest,
) (
	*pbuser.GetUsernameByUserIdResponse, error,
) {
	return s.client.GetUsernameByUserId(
		grpcCtx, request,
	)
}

// ChangeUsername changes the users' username
func (s *Service) ChangeUsername(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangeUsernameRequest,
) (
	*pbuser.ChangeUsernameResponse, error,
) {
	return s.client.ChangeUsername(
		grpcCtx, request,
	)
}

// AddEmail adds an email to a user
func (s *Service) AddEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.AddEmailRequest,
) (*pbuser.AddEmailResponse, error) {
	return s.client.AddEmail(grpcCtx, request)
}

// DeleteEmail deletes an email from a user
func (s *Service) DeleteEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.DeleteEmailRequest,
) (
	*pbuser.DeleteEmailResponse, error,
) {
	return s.client.DeleteEmail(grpcCtx, request)
}

// ChangePrimaryEmail changes the user's primary email
func (s *Service) ChangePrimaryEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangePrimaryEmailRequest,
) (
	*pbuser.ChangePrimaryEmailResponse, error,
) {
	return s.client.ChangePrimaryEmail(
		grpcCtx, request,
	)
}

// SendVerificationEmail sends a verification email to a user
func (s *Service) SendVerificationEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.SendVerificationEmailRequest,
) (
	*pbuser.SendVerificationEmailResponse, error,
) {
	return s.client.SendVerificationEmail(
		grpcCtx, request,
	)
}

// VerifyEmail verifies the user's email
func (s *Service) VerifyEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.VerifyEmailRequest,
) (
	*pbuser.VerifyEmailResponse, error,
) {
	return s.client.VerifyEmail(grpcCtx, request)
}

// GetPrimaryEmail gets the user's primary email
func (s *Service) GetPrimaryEmail(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbuser.GetPrimaryEmailResponse, error,
) {
	return s.client.GetPrimaryEmail(
		grpcCtx, &emptypb.Empty{},
	)
}

// GetActiveEmails gets the user's active emails
func (s *Service) GetActiveEmails(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbuser.GetActiveEmailsResponse, error,
) {
	return s.client.GetActiveEmails(
		grpcCtx, &emptypb.Empty{},
	)
}

// ChangePhoneNumber changes the user's phone number
func (s *Service) ChangePhoneNumber(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ChangePhoneNumberRequest,
) (
	*pbuser.ChangePhoneNumberResponse, error,
) {
	return s.client.ChangePhoneNumber(
		grpcCtx, request,
	)
}

// SendVerificationSMS sends a verification SMS to a user
func (s *Service) SendVerificationSMS(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.SendVerificationSMSRequest,
) (
	*pbuser.SendVerificationSMSResponse, error,
) {
	return s.client.SendVerificationSMS(
		grpcCtx, request,
	)
}

// VerifyPhoneNumber verifies the user's phone number
func (s *Service) VerifyPhoneNumber(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.VerifyPhoneNumberRequest,
) (
	*pbuser.VerifyPhoneNumberResponse, error,
) {
	return s.client.VerifyPhoneNumber(
		grpcCtx, request,
	)
}

// GetPhoneNumber gets the user's phone number
func (s *Service) GetPhoneNumber(
	ctx *gin.Context,
	grpcCtx context.Context,
) (
	*pbuser.GetPhoneNumberResponse, error,
) {
	return s.client.GetPhoneNumber(
		grpcCtx, &emptypb.Empty{},
	)
}

// ForgotPassword sends a forgot password email to a user
func (s *Service) ForgotPassword(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ForgotPasswordRequest,
) (
	*pbuser.ForgotPasswordResponse, error,
) {
	return s.client.ForgotPassword(
		grpcCtx, request,
	)
}

// ResetPassword resets the user's password
func (s *Service) ResetPassword(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.ResetPasswordRequest,
) (
	*pbuser.ResetPasswordResponse, error,
) {
	return s.client.ResetPassword(
		grpcCtx, request,
	)
}

// DeleteUser deletes a user
func (s *Service) DeleteUser(
	ctx *gin.Context,
	grpcCtx context.Context,
	request *pbuser.DeleteUserRequest,
) (
	*pbuser.DeleteUserResponse, error,
) {
	return s.client.DeleteUser(grpcCtx, request)
}
