package auth

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	"medical-record-api/pkg/errors"
	"medical-record-api/pkg/jwt"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/sanitize"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
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
