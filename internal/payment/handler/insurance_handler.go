package payment

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	model "medical-record-api/internal/payment/model"
	service "medical-record-api/internal/payment/service"
	"medical-record-api/pkg/errors"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/sanitize"
	"net/http"
)

type InsuranceHandler struct {
	service *service.InsuranceService
}

func NewInsuranceHandler(service *service.InsuranceService) *InsuranceHandler {
	return &InsuranceHandler{service}
}

func (h *InsuranceHandler) GetInsurances(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	isPaginated := c.DefaultQuery("pagination", "false") == "true"
	if !isPaginated {
		insurances, err := h.service.GetAll(search)
		if err != nil {
			c.Error(err)
			return
		}

		response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, insurances, nil)
		return
	}

	var pagination response.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgInvalidPagination, nil, nil)
		return
	}

	insurances, total, err := h.service.GetAllPaginated(pagination.Page, pagination.Limit, search)
	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, pagination.Page, pagination.Limit)
	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, insurances, meta)
}

func sanitizeInsuranceRequest(request *model.InsuranceRequest) {
	request.Name = sanitize.SanitizeStrict(request.Name)
	request.Code = sanitize.SanitizeUGC(request.Code)
	request.Contact = sanitize.SanitizeStrict(request.Contact)
	request.Email = sanitize.SanitizeStrict(request.Email)
}

func (h *InsuranceHandler) AddInsurance(c *gin.Context) {
	var request model.InsuranceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeInsuranceRequest(&request)

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

func (h *InsuranceHandler) UpdateInsurance(c *gin.Context) {
	id := c.Param("id")
	var request model.InsuranceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeInsuranceRequest(&request)

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

func (h *InsuranceHandler) DeleteInsurance(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
