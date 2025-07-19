package menu

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	"medical-record-api/pkg/jwt"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/role"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetByRoleID(c *gin.Context) {
	claims, ok := jwt.GetUserClaims(c)
	if !ok {
		response.Respond(c, http.StatusUnauthorized, false, constant.MsgUserNotFoundInContext, nil, nil)
		return
	}

	roleID := role.Resolve(claims.Role)
	menus, err := h.service.GetByRoleID(roleID)
	if err != nil {
		response.Respond(c, http.StatusInternalServerError, false, err.Error(), nil, nil)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, menus, nil)
}
