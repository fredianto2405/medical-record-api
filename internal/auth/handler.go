package auth

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	"medical-record-api/internal/email"
	"medical-record-api/pkg/errors"
	"medical-record-api/pkg/jwt"
	"medical-record-api/pkg/logger"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/sanitize"
	"net/http"
)

type Handler struct {
	service      *Service
	emailService *email.Service
}

func NewHandler(service *Service, emailService *email.Service) *Handler {
	return &Handler{
		service:      service,
		emailService: emailService,
	}
}

func sanitizeLoginRequest(request *LoginRequest) {
	request.Email = sanitize.SanitizeStrict(request.Email)
	request.Password = sanitize.SanitizeStrict(request.Password)
}

func (h *Handler) Login(c *gin.Context) {
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	sanitizeLoginRequest(&request)

	if err := errors.Validate.Struct(request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	user, err := h.service.Login(&request)
	if err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	token, err := jwt.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		response.Respond(c, http.StatusInternalServerError, false, constant.MsgFailedGenerateJWT, nil, nil)
		return
	}

	data := &LoginDTO{
		Email:       user.Email,
		AccessToken: token,
		Role:        user.Role,
	}

	response.Respond(c, http.StatusOK, true, constant.MsgLoginSuccess, data, nil)
}

func (h *Handler) ChangePassword(c *gin.Context) {
	claims, ok := jwt.GetUserClaims(c)
	if !ok {
		response.Respond(c, http.StatusUnauthorized, false, constant.MsgUserNotFoundInContext, nil, nil)
		return
	}

	email := claims.Email

	var request ChangePasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if err := errors.Validate.Struct(request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	err := h.service.ChangePassword(email, &request)
	if err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgChangePasswordSuccess, nil, nil)
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	log := logger.Log

	var request ForgotPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if err := errors.Validate.Struct(request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	token, err := h.service.ForgotPassword(request.Email)
	if err != nil {
		log.WithError(err).Error("Error execute forgot password: ")
	}

	if token != "" {
		subject := "[Electronic Medical Record] Password Reset"
		emailBody := buildEmailForgotPassword(token)
		err = h.emailService.QueueEmail(request.Email, subject, emailBody, true)
		if err != nil {
			log.WithError(err).Error("Error queue email: ")
		}
	}

	response.Respond(c, http.StatusOK, true, constant.MsgForgotPassword, nil, nil)
}

func (h *Handler) ResetPassword(c *gin.Context) {
	token := c.Param("token")
	var request ResetPasswordRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if err := errors.Validate.Struct(request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	if err := h.service.ResetPassword(token, &request); err != nil {
		response.Respond(c, http.StatusBadRequest, false, err.Error(), nil, nil)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgResetPasswordSuccess, nil, nil)
}

func buildEmailForgotPassword(token string) string {
	urlRedirect := "http://157.66.54.26:9001/auth/reset-password/" + token
	return `<!DOCTYPE html>
			<html lang="en" style="margin: 0; padding: 0;">
			<head>
			  <meta charset="UTF-8" />
			  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
			  <title>[Electronic Medical Record] Password Reset</title>
			  <style>
				body {
				  font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif;
				  background-color: #f4f4f4;
				  padding: 0;
				  margin: 0;
				}
				.email-container {
				  max-width: 600px;
				  margin: auto;
				  background: #ffffff;
				  padding: 40px;
				  border-radius: 8px;
				  box-shadow: 0 2px 5px rgba(0,0,0,0.1);
				}
				.button {
				  display: inline-block;
				  padding: 12px 24px;
				  background-color: #4CAF50;
				  color: white;
				  text-decoration: none;
				  border-radius: 6px;
				  font-weight: bold;
				}
				.footer {
				  text-align: center;
				  font-size: 12px;
				  color: #888888;
				  margin-top: 30px;
				}
			  </style>
			</head>
			<body>
			  <div class="email-container">
				<h2>Password Reset Request</h2>
				<p>Hello,</p>
				<p>We received a request to reset your password. Click the button below to choose a new password:</p>
				<p style="text-align: center; margin: 30px 0;">
				  <a href="` + urlRedirect + `" class="button">Reset Password</a>
				</p>
				<p>If you didn’t request a password reset, you can ignore this email. Your password will remain unchanged.</p>
				<p>Thank you,<br>Electronic Medical Record Apps</p>
				<div class="footer">
				  © 2025 Electronic Medical Record Apps. All rights reserved.
				</div>
			  </div>
			</body>
			</html>`
}
