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

type InsurancePatientHandler struct {
	service *service.InsurancePatientService
}

func NewInsurancePatientHandler(service *service.InsurancePatientService) *InsurancePatientHandler {
	return &InsurancePatientHandler{service}
}

func sanitizeInsurancePatientRequest(request *model.InsurancePatientRequest) {
	request.PatientID = sanitize.SanitizeUGC(request.PatientID)
	request.InsuranceID = sanitize.SanitizeUGC(request.InsuranceID)
	request.InsuranceNumber = sanitize.SanitizeUGC(request.InsuranceNumber)
}

func (h *InsurancePatientHandler) AddInsurancePatient(c *gin.Context) {
	var request model.InsurancePatientRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeInsurancePatientRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	err := h.service.Create(&request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, nil, nil)
}

func (h *InsurancePatientHandler) DeleteInsurancePatient(c *gin.Context) {
	patientID := c.Param("patientID")
	if err := h.service.Delete(patientID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
