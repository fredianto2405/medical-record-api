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

type EmergencyContactHandler struct {
	service *service.EmergencyContactService
}

func NewEmergencyContactHandler(service *service.EmergencyContactService) *EmergencyContactHandler {
	return &EmergencyContactHandler{service}
}

func sanitizeEmergencyContactRequest(request *model.EmergencyContactRequest) {
	request.PatientID = sanitize.SanitizeUGC(request.PatientID)
	request.Name = sanitize.SanitizeStrict(request.Name)
	request.Phone = sanitize.SanitizeStrict(request.Phone)
	request.Relation = sanitize.SanitizeStrict(request.Relation)
}

func (h *EmergencyContactHandler) AddEmergencyContact(c *gin.Context) {
	var request model.EmergencyContactRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeEmergencyContactRequest(&request)

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

func (h *EmergencyContactHandler) DeleteEmergencyContact(c *gin.Context) {
	patientID := c.Param("patientID")
	if err := h.service.Delete(patientID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
