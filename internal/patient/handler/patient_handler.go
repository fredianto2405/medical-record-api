package patient

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	model "medical-record-api/internal/patient/model"
	service "medical-record-api/internal/patient/service"
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

func (h *Handler) GetPatients(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	isPaginated := c.DefaultQuery("pagination", "false") == "true"
	if !isPaginated {
		patients, err := h.service.GetAll(search)
		if err != nil {
			c.Error(err)
			return
		}

		response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, patients, nil)
		return
	}

	var pagination response.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgInvalidPagination, nil, nil)
		return
	}

	patients, total, err := h.service.GetAllPaginated(pagination.Page, pagination.Limit, search)
	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, pagination.Page, pagination.Limit)
	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, patients, meta)
}

func sanitizePatientRequest(request *model.Request) {
	request.Name = sanitize.SanitizeStrict(request.Name)
	request.Gender = sanitize.SanitizeStrict(request.Gender)
	request.BirthDate = sanitize.SanitizeUGC(request.BirthDate)
	request.BloodType = sanitize.SanitizeStrict(request.BloodType)
	request.Address = sanitize.SanitizeUGC(request.Address)
	request.Phone = sanitize.SanitizeStrict(request.Phone)
	request.Email = sanitize.SanitizeStrict(request.Email)
	request.HistoryOfIllness = sanitize.SanitizeUGC(request.HistoryOfIllness)
	request.Allergies = sanitize.SanitizeUGC(request.Allergies)
	request.IdentityType = sanitize.SanitizeStrict(request.IdentityType)
	request.IdentityNumber = sanitize.SanitizeUGC(request.IdentityNumber)
}

func (h *Handler) AddPatient(c *gin.Context) {
	var request model.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizePatientRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	patient, err := h.service.Create(&request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, patient, nil)
}

func (h *Handler) UpdatePatient(c *gin.Context) {
	id := c.Param("id")
	var request model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizePatientRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	patient, err := h.service.Update(id, &request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataUpdated, patient, nil)
}

func (h *Handler) DeletePatient(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
