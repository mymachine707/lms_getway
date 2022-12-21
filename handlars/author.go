package handlars

import (
	"lms/lms_getway/models"
	"lms/lms_getway/protogen/book_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatAuthor godoc
//
//	@Summary		Creat Author
//	@Description	Creat a new author
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			author			body		models.CreateAuthorModul	true	"Author body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/author [post]
func (h *handler) CreatAuthor(c *gin.Context) {
	var body models.CreateAuthorModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	author, err := h.grpcClient.Author.CreateAuthor(c.Request.Context(), &book_service.CreateAuthorRequest{
		Name: body.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatAuthor",
		Data:    author,
	})
}

// GetAuthorByID godoc
//
//	@Summary		GetAuthorByID
//	@Description	get an author by id
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Author id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/author/{id} [get]
func (h *handler) GetAuthorByID(c *gin.Context) {

	idStr := c.Param("id")

	author, err := h.grpcClient.Author.GetAuthorById(c.Request.Context(), &book_service.GetAuthorByIDRequest{
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
		Data:    author,
	})
}

// GetAuthorList godoc
//
//	@Summary		List authors
//	@Description	GetAuthorList
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Author}
//	@Router			/v1/author/ [get]
func (h *handler) GetAuthorList(c *gin.Context) {

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

	authorList, err := h.grpcClient.Author.GetAuthorList(c.Request.Context(), &book_service.GetAuthorListRequest{
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
		Message: "GetAuthorList OK",
		Data:    authorList,
	})
}

// DeleteAuthor godoc
//
//	@Summary		Delete Author
//	@Description	get element by id and delete this author
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Author id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/author/{id} [delete]
func (h *handler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")

	author, err := h.grpcClient.Author.DeleteAuthor(c.Request.Context(), &book_service.DeleteAuthorRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author Deleted",
		"data":    author,
	})
}

// EnabledAuthor godoc
//
//	@Summary		Enabled Author
//	@Description	get element by id and enabled this author
//	@Tags			author
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Author id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/author/{id} [delete]
func (h *handler) EnabledAuthor(c *gin.Context) {
	idStr := c.Param("id")

	author, err := h.grpcClient.Author.EnabledAuthor(c.Request.Context(), &book_service.EnabledAuthorRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Author Enabled",
		"data":    author,
	})
}
