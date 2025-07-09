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

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service}
}

func (h *Handler) GetMedicines(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	isPaginated := c.DefaultQuery("pagination", "false") == "true"
	if !isPaginated {
		medicines, err := h.service.GetAll(search)
		if err != nil {
			c.Error(err)
			return
		}

		response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, medicines, nil)
		return
	}

	var pagination response.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgInvalidPagination, nil, nil)
		return
	}

	medicines, total, err := h.service.GetAllPaginated(pagination.Page, pagination.Limit, search)
	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, pagination.Page, pagination.Limit)
	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, medicines, meta)
}

func sanitizeMedicineRequest(request *model.Request) {
	request.Code = sanitize.SanitizeUGC(request.Code)
	request.Name = sanitize.SanitizeStrict(request.Name)
	request.CategoryID = sanitize.SanitizeUGC(request.CategoryID)
	request.UnitID = sanitize.SanitizeUGC(request.UnitID)
	request.Dosage = sanitize.SanitizeUGC(request.Dosage)
}

func (h *Handler) AddMedicine(c *gin.Context) {
	var request model.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeMedicineRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	medicine, err := h.service.Create(&request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, medicine, nil)
}

func (h *Handler) UpdateMedicine(c *gin.Context) {
	id := c.Param("id")
	var request model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeMedicineRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	medicine, err := h.service.Update(id, &request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataUpdated, medicine, nil)
}

func (h *Handler) DeleteMedicine(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
