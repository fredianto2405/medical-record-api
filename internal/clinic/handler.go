package clinic

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

func (h *Handler) GetClinic(c *gin.Context) {
	clinic, err := h.service.Get()
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, clinic, nil)
}

func sanitizeClinicRequest(request *Request) {
	request.Name = sanitize.SanitizeStrict(request.Name)
	request.Address = sanitize.SanitizeUGC(request.Address)
	request.Logo = sanitize.SanitizeUGC(request.Logo)
	request.SharingFeeType = sanitize.SanitizeStrict(request.SharingFeeType)
}

func (h *Handler) UpdateClinic(c *gin.Context) {
	id := c.Param("id")
	var request Request
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeClinicRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	clinic, err := h.service.Update(id, &request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataUpdated, clinic, nil)
}
