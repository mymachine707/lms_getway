package handlars

import (
	"lms/lms_getway/models"
	"lms/lms_getway/protogen/book_service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatLocation godoc
//
//	@Summary		Creat Location
//	@Description	Creat a new location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Param			location		body		models.CreateLocationModul	true	"Location body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Location}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/location [post]
func (h *handler) CreatLocation(c *gin.Context) {
	var body models.CreateLocationModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	location, err := h.grpcClient.Location.CreateLocation(c.Request.Context(), &book_service.CreateLocationRequest{
		Name: body.Name,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatLocation",
		Data:    location,
	})
}

// GetLocationByID godoc
//
//	@Summary		GetLocationByID
//	@Description	get an location by id
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Location id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Location}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/location/{id} [get]
func (h *handler) GetLocationByID(c *gin.Context) {

	idStr := c.Param("id")

	location, err := h.grpcClient.Location.GetLocationById(c.Request.Context(), &book_service.GetLocationByIdRequest{
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
		Data:    location,
	})
}

// GetLocationList godoc
//
//	@Summary		List locations
//	@Description	GetLocationList
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Location}
//	@Router			/v1/location/ [get]
func (h *handler) GetLocationList(c *gin.Context) {

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

	locationList, err := h.grpcClient.Location.GetLocationList(c.Request.Context(), &book_service.GetLocationListRequest{
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
		Message: "GetLocationList OK",
		Data:    locationList,
	})
}

// LocationUpdate godoc
//
//	@Summary		Update Location
//	@Description	Update Location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Param			location		body		models.UpdateLocationModul	true	"Location body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=[]models.Location}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/location/ [put]
func (h *handler) LocationUpdate(c *gin.Context) {

	var body models.UpdateLocationModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	location, err := h.grpcClient.Location.UpdateLocation(c.Request.Context(), &book_service.UpdateLocationRequest{
		Id:   body.ID,
		Name: body.Name,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Location Updated",
		"data":    location,
	})

}

// DeleteLocation godoc
//
//	@Summary		Delete Location
//	@Description	get element by id and delete this location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Location id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Location}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/location/{id} [delete]
func (h *handler) DeleteLocation(c *gin.Context) {
	idStr := c.Param("id")

	location, err := h.grpcClient.Location.DeleteLocation(c.Request.Context(), &book_service.DeleteLocationRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Location Deleted",
		"data":    location,
	})

}

// EnabledLocation godoc
//
//	@Summary		Enabled Location
//	@Description	get element by id and delete this location
//	@Tags			location
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Location id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Location}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/location/{id} [delete]
func (h *handler) EnabledLocation(c *gin.Context) {
	idStr := c.Param("id")

	location, err := h.grpcClient.Location.EnabledLocation(c.Request.Context(), &book_service.EnabledLocationRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Location Enabled",
		"data":    location,
	})

}
