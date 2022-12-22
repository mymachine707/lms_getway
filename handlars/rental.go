package handlars

import (
	"fmt"
	"lms/lms_getway/models"
	"lms/lms_getway/protogen/book_service"
	"lms/lms_getway/protogen/rental_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatRental godoc
//
//	@Summary		Creat Rental
//	@Description	Creat a new rental
//	@Tags			rental
//	@Accept			json
//	@Produce		json
//	@Param			rental			body		models.CreateRentalModul	true	"Rental body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Author}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/rental [post]
func (h *handler) CreatRental(c *gin.Context) {
	var body models.CreateRentalModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	rental, err := h.grpcClient.Book.GetBookById(c.Request.Context(), &book_service.GetBookByIDRequest{
		Id: body.BookID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatRental",
		Data:    rental,
	})
}

// GetRentalByID godoc
//
//	@Summary		GetRentalByID
//	@Description	get an rental by id
//	@Tags			rental
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Rental id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Rental}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/rental/{id} [get]
func (h *handler) GetRentalByID(c *gin.Context) {

	idStr := c.Param("id")

	rental, err := h.grpcClient.Rental.GetRentalById(c.Request.Context(), &rental_service.GetRentalByIdRequest{
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
		Data:    rental,
	})
}

// GetRentalList godoc
//
//	@Summary		List rentals
//	@Description	GetRentalList
//	@Tags			rental
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Rental}
//	@Router			/v1/rental/ [get]
func (h *handler) GetRentalList(c *gin.Context) {

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

	rentalList, err := h.grpcClient.Rental.GetRentalList(c.Request.Context(), &rental_service.GetRentalListRequest{
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
		Message: "GetRentalList OK",
		Data:    rentalList,
	})
}

// DeleteRental godoc
//
//	@Summary		Delete Rental
//	@Description	get element by id and delete this rental
//	@Tags			rental
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Rental id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Rental}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/rental/{id} [DELETE]
func (h *handler) DeleteRental(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println(idStr)
	rental, err := h.grpcClient.Rental.DeleteRental(c.Request.Context(), &rental_service.DeleteRentalRequest{
		Id: idStr,
	})
	fmt.Println(rental)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Rental Deleted",
		"data":    rental,
	})
}

// UpdateRental godoc
//
//	@Summary		Update Rental
//	@Description	get element by id and enabled this rental
//	@Tags			rental
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Rental id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Rental}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/rental/{id} [PUT]
func (h *handler) UpdateRental(c *gin.Context) {
	idStr := c.Param("id")
	fmt.Println(idStr)
	rental, err := h.grpcClient.Rental.UpdateRental(c.Request.Context(), &rental_service.UpdateRentalRequest{
		Id: idStr,
	})
	fmt.Println(rental)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Rental Update",
		"data":    rental,
	})
}
