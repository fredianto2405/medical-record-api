package medicine

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	model "medical-record-api/internal/medicine/model"
	service "medical-record-api/internal/medicine/service"
	"medical-record-api/pkg/errors"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/sanitize"
	"net/http"
)

type UnitHandler struct {
	service *service.UnitService
}

func NewUnitHandler(service *service.UnitService) *UnitHandler {
	return &UnitHandler{service}
}

func (h *UnitHandler) GetUnits(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	isPaginated := c.DefaultQuery("pagination", "false") == "true"
	if !isPaginated {
		units, err := h.service.GetAll(search)
		if err != nil {
			c.Error(err)
			return
		}

		response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, units, nil)
		return
	}

	var pagination response.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgInvalidPagination, nil, nil)
		return
	}

	units, total, err := h.service.GetAllPaginated(pagination.Page, pagination.Limit, search)
	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, pagination.Page, pagination.Limit)
	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, units, meta)
}

func sanitizeUnitRequest(request *model.UnitRequest) {
	request.Name = sanitize.SanitizeStrict(request.Name)
}

func (h *UnitHandler) AddUnit(c *gin.Context) {
	var request model.UnitRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeUnitRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	unit, err := h.service.Create(&request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, unit, nil)
}

func (h *UnitHandler) UpdateUnit(c *gin.Context) {
	id := c.Param("id")
	var request model.UnitRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeUnitRequest(&request)

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

func (h *UnitHandler) DeleteUnit(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
