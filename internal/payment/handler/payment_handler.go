package payment

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	model "medical-record-api/internal/payment/model"
	payment "medical-record-api/internal/payment/service"
	"medical-record-api/pkg/errors"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/sanitize"
	"net/http"
)

type Handler struct {
	service *payment.Service
}

func NewHandler(service *payment.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetMethods(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	isPaginated := c.DefaultQuery("pagination", "false") == "true"
	if !isPaginated {
		methods, err := h.service.GetAll(search)
		if err != nil {
			c.Error(err)
			return
		}

		response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, methods, nil)
		return
	}

	var pagination response.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgInvalidPagination, nil, nil)
		return
	}

	methods, total, err := h.service.GetAllPaginated(pagination.Page, pagination.Limit, search)
	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, pagination.Page, pagination.Limit)
	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, methods, meta)
}

func sanitizeMethodsRequest(request *model.MethodRequest) {
	request.Name = sanitize.SanitizeStrict(request.Name)
	request.Description = sanitize.SanitizeUGC(request.Description)
}

func (h *Handler) AddMethod(c *gin.Context) {
	var request model.MethodRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeMethodsRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	insurance, err := h.service.Create(&request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, insurance, nil)
}

func (h *Handler) UpdateMethod(c *gin.Context) {
	id := c.Param("id")
	var request model.MethodRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeMethodsRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	unit, err := h.service.Update(id, &request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataUpdated, unit, nil)
}

func (h *Handler) DeleteMethod(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}

func (h *Handler) GetStatuses(c *gin.Context) {
	statuses, err := h.service.GetAllStatus()
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, statuses, nil)
}
