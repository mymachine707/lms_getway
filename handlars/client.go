package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/eCommerce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateSuperUser godoc
//
//	@Summary		CreateSuperUser
//	@Description	CreateSuperUser
//	@Tags			CreateSuperUser
//	@Accept			json
//	@Produce		json
//	@Param			client	body		models.CreateClientModul	true	"Client body"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201		{object}	models.JSONResult{data=models.Client}
//	@Failure		400		{object}	models.JSONErrorResponse
//	@Router			/v1/sudo [post]
func (h *handler) CreateSuperUser(c *gin.Context) {
	var body models.CreateClientModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	client, err := h.grpcClient.Client.CreateSuperUser(c.Request.Context(), &eCommerce.CreateSudoRequest{
		Firstname:   body.Firstname,
		Lastname:    body.Lastname,
		Username:    body.Username,
		PhoneNumber: body.PhoneNumber,
		Address:     body.Address,
		Type:        body.Type,
		Password:    body.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreateSuperUser",
		Data:    client,
	})
}

// GetClientByID godoc
//
//	@Summary		GetClientByID
//	@Description	get an client by id
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Client id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Client}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/client/{id} [get]
func (h *handler) GetClientByID(c *gin.Context) {

	idStr := c.Param("id")

	client, err := h.grpcClient.Client.GetClientById(c.Request.Context(), &eCommerce.GetClientByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "GetClientById OK",
		Data:    client,
	})
}

// GetClientList godoc
//
//	@Summary		List clients
//	@Description	GetClientList
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"
//	@Param			limit			query		int		false	"100"
//	@Param			search			query		string	false	"search exapmle"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Client}
//	@Router			/v1/client/ [get]
func (h *handler) GetClientList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", h.cfg.Default_Offset)
	limitStr := c.DefaultQuery("limit", h.cfg.Default_Limit)

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

	clientList, err := h.grpcClient.Client.GetClientList(c.Request.Context(), &eCommerce.GetClientListRequest{
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "GetClientList OK",
		Data:    clientList,
	})
}

// ClientUpdate godoc
//
//	@Summary		Update Client
//	@Description	Update Client
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			client			body		models.UpdateClientModul	true	"Client body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=[]models.Client}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/client/ [put]
func (h *handler) ClientUpdate(c *gin.Context) {

	var body models.UpdateClientModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	client, err := h.grpcClient.Client.UpdateClient(c.Request.Context(), &eCommerce.UpdateClientRequest{
		Id:          body.Id,
		PhoneNumber: body.PhoneNumber,
		Address:     body.Address,
		Password:    body.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Client Updated",
		"data":    client,
	})

}

// DeleteClient godoc
//
//	@Summary		Delete Client
//	@Description	get element by id and delete this client
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Client id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Client}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/client/{id} [delete]
func (h *handler) DeleteClient(c *gin.Context) {
	idStr := c.Param("id")

	client, err := h.grpcClient.Client.DeleteClient(c.Request.Context(), &eCommerce.DeleteClientRequest{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Client Deleted",
		"data":    client,
	})

}

// CreatClient godoc
//
//	@Summary		Creat Client
//	@Description	Creat a new client
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			client	body		models.CreateClientModul	true	"Client body"
//	@Success		201		{object}	models.JSONResult{data=models.Client}
//	@Failure		400		{object}	models.JSONErrorResponse
//	@Router			/v1/client [post]
func (h *handler) CreatClient(c *gin.Context) {
	var body models.CreateClientModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	client, err := h.grpcClient.Client.CreateClient(c.Request.Context(), &eCommerce.CreateClientRequest{
		Firstname:   body.Firstname,
		Lastname:    body.Lastname,
		Username:    body.Username,
		PhoneNumber: body.PhoneNumber,
		Address:     body.Address,
		Type:        body.Type,
		Password:    body.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "CreatClient",
		Data:    client,
	})
}
