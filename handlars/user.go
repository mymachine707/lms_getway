package handlars

import (
	"mymachine707/models"
	"mymachine707/protogen/eCommerce"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware ...
//
//	@param	Authorization	header	string	false	"Authorization"
func (h *handler) AuthMiddleware(userType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		hasAccessTokenResponse, err := h.grpcClient.Client.HasAccess(c.Request.Context(), &eCommerce.TokenRequest{
			Token: token,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
				Error: err.Error(),
			})
			c.Abort()
			return
		}

		if !hasAccessTokenResponse.HasAccess {
			c.JSON(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		if userType != "*" {
			if hasAccessTokenResponse.User.Type != userType {
				c.JSON(http.StatusUnauthorized, "permission denied")
				c.Abort()
			}
		}

		c.Set("client_username", hasAccessTokenResponse.User.Username)
		c.Set("client_userID", hasAccessTokenResponse.User.Id)

		c.Next()
		//
	}
}

// Login godoc
//
//	@Summary		Login
//	@Description	Login
//	@Tags			client
//	@Accept			json
//	@Produce		json
//	@Param			users	body		models.LoginModul	true	"Login body"
//	@Success		201		{object}	models.JSONResult{data=models.TokenResponse}
//	@Failure		400		{object}	models.JSONErrorResponse
//	@Router			/v1/login [post]
func (h *handler) Login(c *gin.Context) {

	var body models.LoginModul
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	tokenResponse, err := h.grpcClient.Client.Login(c.Request.Context(), &eCommerce.LoginAuthRequest{
		Username: body.Username,
		Password: body.Password,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "Login",
		Data:    tokenResponse,
	})
}
