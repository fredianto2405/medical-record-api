package medical_record

import (
	"github.com/gin-gonic/gin"
	"medical-record-api/internal/constant"
	model "medical-record-api/internal/medical_record/model"
	service "medical-record-api/internal/medical_record/service"
	"medical-record-api/pkg/errors"
	"medical-record-api/pkg/response"
	"medical-record-api/pkg/sanitize"
	"net/http"
)

type Handler struct {
	service                *service.Service
	nurseAssignmentService *service.NurseAssignmentService
	treatmentDetailService *service.TreatmentDetailService
	recipeService          *service.RecipeService
	statusService          *service.StatusService
}

func NewHandler(service *service.Service,
	nurseAssignmentService *service.NurseAssignmentService,
	treatmentDetailService *service.TreatmentDetailService,
	recipeService *service.RecipeService,
	statusService *service.StatusService) *Handler {
	return &Handler{
		service:                service,
		nurseAssignmentService: nurseAssignmentService,
		treatmentDetailService: treatmentDetailService,
		recipeService:          recipeService,
		statusService:          statusService,
	}
}

func (h *Handler) GetStatuses(c *gin.Context) {
	statuses, err := h.statusService.GetAll()
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, statuses, nil)
}

func sanitizeMedicalRecordRequest(request *model.Request) {
	request.PatientID = sanitize.SanitizeUGC(request.PatientID)
	request.DoctorID = sanitize.SanitizeUGC(request.DoctorID)
	request.Diagnosis = sanitize.SanitizeStrict(request.Diagnosis)
	request.Notes = sanitize.SanitizeStrict(request.Notes)
	request.PaymentMethodID = sanitize.SanitizeUGC(request.PaymentMethodID)
	request.InsuranceID = sanitize.SanitizeUGC(request.InsuranceID)
	request.Anamnesis = sanitize.SanitizeStrict(request.Anamnesis)

	for i, id := range request.NurseIDs {
		request.NurseIDs[i] = sanitize.SanitizeUGC(id)
	}

	for i := range request.Treatments {
		request.Treatments[i].TreatmentID = sanitize.SanitizeUGC(request.Treatments[i].TreatmentID)
	}

	for i := range request.Recipes {
		request.Recipes[i].MedicineID = sanitize.SanitizeUGC(request.Recipes[i].MedicineID)
		request.Recipes[i].Dosage = sanitize.SanitizeStrict(request.Recipes[i].Dosage)
		request.Recipes[i].Instruction = sanitize.SanitizeStrict(request.Recipes[i].Instruction)
	}
}

func (h *Handler) AddMedicalRecord(c *gin.Context) {
	var request model.Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeMedicalRecordRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	medicalRecordID, err := h.service.Create(&request)
	if err != nil {
		c.Error(err)
		return
	}

	medicalRecord := &model.DTO{
		MedicalRecordID: medicalRecordID,
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, medicalRecord, nil)
}

func (h *Handler) DeleteMedicalRecord(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}

func (h *Handler) AddNurseAssignment(c *gin.Context) {
	id := c.Param("id")
	var request model.NurseAssignmentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	if err := h.nurseAssignmentService.Create(id, request.NurseID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, nil, nil)
}

func (h *Handler) DeleteNurseAssignment(c *gin.Context) {
	id := c.Param("id")
	nurseID := c.Param("nurseId")
	if err := h.nurseAssignmentService.Delete(id, nurseID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}

func (h *Handler) DeleteTreatmentDetail(c *gin.Context) {
	id := c.Param("id")
	treatmentID := c.Param("treatmentId")
	if err := h.treatmentDetailService.Delete(id, treatmentID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}

func (h *Handler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")
	medicineID := c.Param("medicineId")
	if err := h.recipeService.Delete(id, medicineID); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
