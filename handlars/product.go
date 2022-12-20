package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/eCommerce"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreatProduct godoc
//
//	@Summary		Creat Product
//	@Description	Creat a new product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			product			body		models.CreateProductModul	true	"Product body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Product}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/product [post]
func (h *handler) CreatProduct(c *gin.Context) {

	var body models.CreateProductModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	product, err := h.grpcClient.Product.CreateProduct(c.Request.Context(), &eCommerce.CreateProductRequest{
		CategoryId:  body.Category_id,
		ProductName: body.Product_name,
		Description: body.Description,
		Price:       body.Price,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "Product Created",
		Data:    product,
	})
}

// GetProductByID godoc
//
//	@Summary		GetProductByID
//	@Description	get an product by id
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Product id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Product}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/product/{id} [get]
func (h *handler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")

	product, err := h.grpcClient.Product.GetProductById(c.Request.Context(), &eCommerce.GetProductByIDRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "GetProductByID",
		Data:    product,
	})
}

// GetProductList godoc
//
//	@Summary		List products
//	@Description	GetProductList
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			offset			query		int		false	"0"			default()
//	@Param			limit			query		int		false	"100"		default()
//	@Param			search			query		string	false	"search"	default()
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Product}
//	@Router			/v1/product/ [get]
func (h *handler) GetProductList(c *gin.Context) {

	offsetStr := c.DefaultQuery("offset", h.cfg.Default_Offset)
	limitStr := c.DefaultQuery("limit", h.cfg.Default_Limit)
	searchStr := c.DefaultQuery("search", "")

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	productList, err := h.grpcClient.Product.GetProductList(c.Request.Context(), &eCommerce.GetProductListRequest{
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
		Data:    productList,
		Message: "GetList OK",
	})
}

// ProductUpdate godoc
//
//	@Summary		Update Product
//	@Description	Update Product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			product			body		models.UpdateProductModul	true	"Product body"
//	@Param			Authorization	header		string						false	"Authorization"
//	@Success		200				{object}	models.JSONResult{data=[]models.Product}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/product/ [put]
func (h *handler) ProductUpdate(c *gin.Context) {

	var body models.UpdateProductModul

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	product, err := h.grpcClient.Product.UpdateProduct(c.Request.Context(), &eCommerce.UpdateProductRequest{
		Id:          body.ID,
		Description: body.Description,
		Price:       body.Price,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Update",
		"data":    product,
	})

}

// DeleteProduct godoc
//
//	@Summary		Delete Product
//	@Description	get element by id and delete this product
//	@Tags			product
//	@Accept			json
//	@Produce		json
//	@Param			id				path		string	true	"Product id"
//	@Param			Authorization	header		string	false	"Authorization"
//	@Success		201				{object}	models.JSONResult{data=models.Product}
//	@Failure		400				{object}	models.JSONErrorResponse
//	@Router			/v1/product/{id} [delete]
func (h *handler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")

	product, err := h.grpcClient.Product.DeleteProduct(c.Request.Context(), &eCommerce.DeleteProductRequest{
		Id: idStr,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Deleted",
		"data":    product,
	})
}
