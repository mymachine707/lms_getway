package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/eCommerce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatCategory godoc
//
//	@Summary		Creat Category
//	@Description	Creat a new category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category		body		models.CreateCategoryModul	true	"Category body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Category}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category [post]
func (h *handler) CreatCategory(c *gin.Context) {
	var body models.CreateCategoryModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	category, err := h.grpcClient.Category.CreateCategory(c.Request.Context(), &eCommerce.CreateCategoryRequest{
		CategoryName: body.Category_name,
		Description:  body.Description,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatCategory",
		Data:    category,
	})
}

// GetCategoryByID godoc
//
//	@Summary		GetCategoryByID
//	@Description	get an category by id
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Category id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Category}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category/{id} [get]
func (h *handler) GetCategoryByID(c *gin.Context) {

	idStr := c.Param("id")

	category, err := h.grpcClient.Category.GetCategoryById(c.Request.Context(), &eCommerce.GetCategoryByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "OK",
		Data:    category,
	})
}

// GetCategoryList godoc
//
//	@Summary		List categorys
//	@Description	GetCategoryList
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Category}
//	@Router			/v1/category/ [get]
func (h *handler) GetCategoryList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", h.cfg.Default_Offset)
	limitStr := c.DefaultQuery("limit", h.cfg.Default_Limit)
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	categoryList, err := h.grpcClient.Category.GetCategoryList(c.Request.Context(), &eCommerce.GetCategoryListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
		Search: searchStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "GetCategoryList OK",
		Data:    categoryList,
	})
}

// CategoryUpdate godoc
//
//	@Summary		Update Category
//	@Description	Update Category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			category		body		models.UpdateCategoryModul	true	"Category body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=[]models.Category}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category/ [put]
func (h *handler) CategoryUpdate(c *gin.Context) {

	var body models.UpdateCategoryModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	category, err := h.grpcClient.Category.UpdateCategory(c.Request.Context(), &eCommerce.UpdateCategoryRequest{
		Id:           body.ID,
		CategoryName: body.Category_name,
		Description:  body.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category Updated",
		"data":    category,
	})

}

// DeleteCategory godoc
//
//	@Summary		Delete Category
//	@Description	get element by id and delete this category
//	@Tags			category
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Category id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Category}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/category/{id} [delete]
func (h *handler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")

	category, err := h.grpcClient.Category.DeleteCategory(c.Request.Context(), &eCommerce.DeleteCategoryRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Category Deleted",
		"data":    category,
	})

}
