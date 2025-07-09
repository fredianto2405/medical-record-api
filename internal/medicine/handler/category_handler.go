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

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	search := c.DefaultQuery("search", "")

	isPaginated := c.DefaultQuery("pagination", "false") == "true"
	if !isPaginated {
		categories, err := h.service.GetAll(search)
		if err != nil {
			c.Error(err)
			return
		}

		response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, categories, nil)
		return
	}

	var pagination response.Pagination
	if err := c.ShouldBindQuery(&pagination); err != nil {
		response.Respond(c, http.StatusBadRequest, false, constant.MsgInvalidPagination, nil, nil)
		return
	}

	categories, total, err := h.service.GetAllPaginated(pagination.Page, pagination.Limit, search)
	if err != nil {
		c.Error(err)
		return
	}

	meta := response.NewMeta(total, pagination.Page, pagination.Limit)
	response.Respond(c, http.StatusOK, true, constant.MsgDataRetrieved, categories, meta)
}

func sanitizeCategoryRequest(request *model.CategoryRequest) {
	request.Name = sanitize.SanitizeStrict(request.Name)
}

func (h *CategoryHandler) AddCategory(c *gin.Context) {
	var request model.CategoryRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeCategoryRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	category, err := h.service.Create(&request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusCreated, true, constant.MsgDataSaved, category, nil)
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	id := c.Param("id")
	var request model.CategoryRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Error(err)
		return
	}

	sanitizeCategoryRequest(&request)

	if err := errors.Validate.Struct(&request); err != nil {
		c.Error(err)
		return
	}

	category, err := h.service.Update(id, &request)
	if err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataUpdated, category, nil)
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	if err := h.service.Delete(id); err != nil {
		c.Error(err)
		return
	}

	response.Respond(c, http.StatusOK, true, constant.MsgDataDeleted, nil, nil)
}
