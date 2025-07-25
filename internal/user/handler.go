package user

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	"medical-record-api/pkg/errors"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/sanitize"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsers(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	var pagination response.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgInvalidPagination, nil, nil)
		return
	}

	users, total, err := h.service.GetAllPaginated(pagination.Page, pagination.Limit, search)

	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, pagination.Page, pagination.Limit)
	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, users, meta)
}

func sanitizeRequest(request *Request) {
	request.Email = sanitize.SanitizeStrict(request.Email)
}

func (h *Handler) AddUser(c *gin.Context) {
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	emailExists, err := h.service.EmailExists(request.Email)
	if err != nil {
		c.Error(err)
		return
	}

	if emailExists {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgEmailExists, nil, nil)
		return
	}

	if err := h.service.Create(&request); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, nil, nil)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
