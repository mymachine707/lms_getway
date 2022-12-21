package handlars

import (
	"lms/lms_getway/models"
	"lms/lms_getway/protogen/book_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatBook godoc
//
//	@Summary		Creat Book
//	@Description	Creat a new book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			book			body		models.CreateBookModul	true	"Book body"
//	@Param			Authorization	header		string					false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Book}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/book [post]
func (h *handler) CreatBook(c *gin.Context) {
	var body models.CreateBookModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	book, err := h.grpcClient.Book.CreateBook(c.Request.Context(), &book_service.CreateBookRequest{
		Name:       body.Name,
		AuthorId:   body.AuthorID,
		CategoryId: body.CategoryID,
		LocationId: body.LocationID,
		ISBN:       body.ISBN,
		Quantity:   body.Quantity,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatBook",
		Data:    book,
	})
}

// GetBookByID godoc
//
//	@Summary		GetBookByID
//	@Description	get an book by id
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Book id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Book}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/book/{id} [get]
func (h *handler) GetBookByID(c *gin.Context) {

	idStr := c.Param("id")

	book, err := h.grpcClient.Book.GetBookById(c.Request.Context(), &book_service.GetBookByIDRequest{
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
		Data:    book,
	})
}

// GetBookList godoc
//
//	@Summary		List books
//	@Description	GetBookList
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Book}
//	@Router			/v1/book/ [get]
func (h *handler) GetBookList(c *gin.Context) {

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

	bookList, err := h.grpcClient.Book.GetBookList(c.Request.Context(), &book_service.GetBookListRequest{
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
		Message: "GetBookList OK",
		Data:    bookList,
	})
}

// BookUpdate godoc
//
//	@Summary		Update Book
//	@Description	Update Book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			book			body		models.UpdateBookModul	true	"Book body"
//	@Param			Authorization	header		string					false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=[]models.Book}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/book/ [put]
func (h *handler) BookUpdate(c *gin.Context) {

	var body models.UpdateBookModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	book, err := h.grpcClient.Book.UpdateBook(c.Request.Context(), &book_service.UpdateBookRequest{
		Id:         body.Id,
		LocationId: body.LocationID,
		Quantity:   body.Quantity,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book Updated",
		"data":    book,
	})

}

// DeleteBook godoc
//
//	@Summary		Delete Book
//	@Description	get element by id and delete this book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Book id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Book}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/book/{id} [delete]
func (h *handler) DeleteBook(c *gin.Context) {
	idStr := c.Param("id")

	book, err := h.grpcClient.Book.DeleteBook(c.Request.Context(), &book_service.DeleteBookRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book Deleted",
		"data":    book,
	})

}

// EnabledBook godoc
//
//	@Summary		Enabled Book
//	@Description	get element by id and delete this book
//	@Tags			book
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Book id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Book}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/book/{id} [PUT]
func (h *handler) EnabledBook(c *gin.Context) {
	idStr := c.Param("id")

	book, err := h.grpcClient.Book.EnabledBook(c.Request.Context(), &book_service.EnabledBookRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book Enabled",
		"data":    book,
	})

}
