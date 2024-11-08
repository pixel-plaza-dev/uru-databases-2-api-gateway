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

// UpdateProfile updates the profile of a user
func (s *Service) UpdateProfile(ctx *gin.Context) (*pbuser.UpdateProfileResponse, error) {
	var updateProfileRequest pbuser.UpdateProfileRequest

	// Bind the request
	if err := ctx.BindJSON(&updateProfileRequest); err != nil {
		return nil, err
	}

	// Call the client
	updateProfileResponse, err := s.client.UpdateProfile(ctx, &updateProfileRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return updateProfileResponse, nil
}

// GetProfile gets the profile of a user
func (s *Service) GetProfile(ctx *gin.Context) (*pbuser.GetProfileResponse, error) {
	// Call the client
	getProfileResponse, err := s.client.GetProfile(ctx, &pbuser.GetProfileRequest{})
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return getProfileResponse, nil
}

// GetFullProfile gets the full profile of a user
func (s *Service) GetFullProfile(ctx *gin.Context) (*pbuser.GetFullProfileResponse, error) {
	// Call the client
	getFullProfileResponse, err := s.client.GetFullProfile(ctx, &pbuser.GetFullProfileRequest{})
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return getFullProfileResponse, nil
}

// ChangePassword changes the password of a user
func (s *Service) ChangePassword(ctx *gin.Context) (*pbuser.ChangePasswordResponse, error) {
	var changePasswordRequest pbuser.ChangePasswordRequest

	// Bind the request
	if err := ctx.BindJSON(&changePasswordRequest); err != nil {
		return nil, err
	}

	// Call the client
	changePasswordResponse, err := s.client.ChangePassword(ctx, &changePasswordRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return changePasswordResponse, nil
}

// ChangeUsername changes the username of a user
func (s *Service) ChangeUsername(ctx *gin.Context) (*pbuser.ChangeUsernameResponse, error) {
	var changeUsernameRequest pbuser.ChangeUsernameRequest

	// Bind the request
	if err := ctx.BindJSON(&changeUsernameRequest); err != nil {
		return nil, err
	}

	// Call the client
	changeUsernameResponse, err := s.client.ChangeUsername(ctx, &changeUsernameRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return changeUsernameResponse, nil
}

// AddEmail adds an email to a user
func (s *Service) AddEmail(ctx *gin.Context) (*pbuser.AddEmailResponse, error) {
	var addEmailRequest pbuser.AddEmailRequest

	// Bind the request
	if err := ctx.BindJSON(&addEmailRequest); err != nil {
		return nil, err
	}

	// Call the client
	addEmailResponse, err := s.client.AddEmail(ctx, &addEmailRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return addEmailResponse, nil
}

// DeleteEmail deletes an email from a user
func (s *Service) DeleteEmail(ctx *gin.Context) (*pbuser.DeleteEmailResponse, error) {
	var deleteEmailRequest pbuser.DeleteEmailRequest

	// Bind the request
	if err := ctx.BindJSON(&deleteEmailRequest); err != nil {
		return nil, err
	}

	// Call the client
	deleteEmailResponse, err := s.client.DeleteEmail(ctx, &deleteEmailRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return deleteEmailResponse, nil
}

// ChangePrimaryEmail changes the primary email of a user
func (s *Service) ChangePrimaryEmail(ctx *gin.Context) (*pbuser.ChangePrimaryEmailResponse, error) {
	var changePrimaryEmailRequest pbuser.ChangePrimaryEmailRequest

	// Bind the request
	if err := ctx.BindJSON(&changePrimaryEmailRequest); err != nil {
		return nil, err
	}

	// Call the client
	changePrimaryEmailResponse, err := s.client.ChangePrimaryEmail(ctx, &changePrimaryEmailRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return changePrimaryEmailResponse, nil
}

// SendVerificationEmail sends a verification email to a user
func (s *Service) SendVerificationEmail(ctx *gin.Context) (*pbuser.SendVerificationEmailResponse, error) {
	var sendVerificationEmailRequest pbuser.SendVerificationEmailRequest

	// Bind the request
	if err := ctx.BindJSON(&sendVerificationEmailRequest); err != nil {
		return nil, err
	}

	// Call the client
	sendVerificationEmailResponse, err := s.client.SendVerificationEmail(ctx, &sendVerificationEmailRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return sendVerificationEmailResponse, nil
}

// VerifyEmail verifies the email of a user
func (s *Service) VerifyEmail(ctx *gin.Context) (*pbuser.VerifyEmailResponse, error) {
	var verifyEmailRequest pbuser.VerifyEmailRequest

	// Bind the request
	if err := ctx.BindJSON(&verifyEmailRequest); err != nil {
		return nil, err
	}

	// Call the client
	verifyEmailResponse, err := s.client.VerifyEmail(ctx, &verifyEmailRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return verifyEmailResponse, nil
}

// GetPrimaryEmail gets the primary email of a user
func (s *Service) GetPrimaryEmail(ctx *gin.Context) (*pbuser.GetPrimaryEmailResponse, error) {
	// Call the client
	getPrimaryEmailResponse, err := s.client.GetPrimaryEmail(ctx, &pbuser.GetPrimaryEmailRequest{})
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return getPrimaryEmailResponse, nil
}

// GetActiveEmails gets the active emails of a user
func (s *Service) GetActiveEmails(ctx *gin.Context) (*pbuser.GetActiveEmailsResponse, error) {
	// Call the client
	getActiveEmailsResponse, err := s.client.GetActiveEmails(ctx, &pbuser.GetActiveEmailsRequest{})
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return getActiveEmailsResponse, nil
}

// ChangePhoneNumber changes the phone number of a user
func (s *Service) ChangePhoneNumber(ctx *gin.Context) (*pbuser.ChangePhoneNumberResponse, error) {
	var changePhoneNumberRequest pbuser.ChangePhoneNumberRequest

	// Bind the request
	if err := ctx.BindJSON(&changePhoneNumberRequest); err != nil {
		return nil, err
	}

	// Call the client
	changePhoneNumberResponse, err := s.client.ChangePhoneNumber(ctx, &changePhoneNumberRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return changePhoneNumberResponse, nil
}

// SendVerificationPhoneNumber sends a verification phone number to a user
func (s *Service) SendVerificationPhoneNumber(ctx *gin.Context) (*pbuser.SendVerificationPhoneNumberResponse, error) {
	var sendVerificationPhoneNumberRequest pbuser.SendVerificationPhoneNumberRequest

	// Bind the request
	if err := ctx.BindJSON(&sendVerificationPhoneNumberRequest); err != nil {
		return nil, err
	}

	// Call the client
	sendVerificationPhoneNumberResponse, err := s.client.SendVerificationPhoneNumber(ctx, &sendVerificationPhoneNumberRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return sendVerificationPhoneNumberResponse, nil
}

// VerifyPhoneNumber verifies the phone number of a user
func (s *Service) VerifyPhoneNumber(ctx *gin.Context) (*pbuser.VerifyPhoneNumberResponse, error) {
	var verifyPhoneNumberRequest pbuser.VerifyPhoneNumberRequest

	// Bind the request
	if err := ctx.BindJSON(&verifyPhoneNumberRequest); err != nil {
		return nil, err
	}

	// Call the client
	verifyPhoneNumberResponse, err := s.client.VerifyPhoneNumber(ctx, &verifyPhoneNumberRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return verifyPhoneNumberResponse, nil
}

// GetPhoneNumber gets the phone number of a user
func (s *Service) GetPhoneNumber(ctx *gin.Context) (*pbuser.GetPhoneNumberResponse, error) {
	// Call the client
	getPhoneNumberResponse, err := s.client.GetPhoneNumber(ctx, &pbuser.GetPhoneNumberRequest{})
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return getPhoneNumberResponse, nil
}

// ForgotPassword sends a forgot password email to a user
func (s *Service) ForgotPassword(ctx *gin.Context) (*pbuser.ForgotPasswordResponse, error) {
	var forgotPasswordRequest pbuser.ForgotPasswordRequest

	// Bind the request
	if err := ctx.BindJSON(&forgotPasswordRequest); err != nil {
		return nil, err
	}

	// Call the client
	forgotPasswordResponse, err := s.client.ForgotPassword(ctx, &forgotPasswordRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return forgotPasswordResponse, nil
}

// ResetPassword resets the password of a user
func (s *Service) ResetPassword(ctx *gin.Context) (*pbuser.ResetPasswordResponse, error) {
	var resetPasswordRequest pbuser.ResetPasswordRequest

	// Bind the request
	if err := ctx.BindJSON(&resetPasswordRequest); err != nil {
		return nil, err
	}

	// Call the client
	resetPasswordResponse, err := s.client.ResetPassword(ctx, &resetPasswordRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return resetPasswordResponse, nil
}

// DeleteUser deletes a user
func (s *Service) DeleteUser(ctx *gin.Context) (*pbuser.DeleteUserResponse, error) {
	var deleteUserRequest pbuser.DeleteUserRequest

	// Bind the request
	if err := ctx.BindJSON(&deleteUserRequest); err != nil {
		return nil, err
	}

	// Call the client
	deleteUserResponse, err := s.client.DeleteUser(ctx, &deleteUserRequest)
	if err != nil {
		return nil, commonclientctx.ExtractErrorFromStatus(err)
	}
	return deleteUserResponse, nil
}
