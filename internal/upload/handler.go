package upload

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"medical-record-api/internal/constant"
	"medical-record-api/pkg/response"
	"net/http"
	"path/filepath"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.Error(err)
		return
	}

	ext := filepath.Ext(file.Filename)
	newFilename := uuid.New().String() + ext
	dst := "uploads/" + newFilename
	if err = c.SaveUploadedFile(file, dst); err != nil {
		c.Error(err)
		return
	}

	entity := &Entity{
		Filename: newFilename,
		Filepath: dst,
		Size:     file.Size,
		MimeType: file.Header.Get("Content-Type"),
	}

	var uploadID string
	uploadID, err = h.service.Create(entity)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgFileUploaded, gin.H{"upload_id": uploadID}, nil)
}
