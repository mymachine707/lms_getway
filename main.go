package main

import (
	"fmt"
	"lms/lms_getway/clients"
	"lms/lms_getway/config"
	"lms/lms_getway/docs"
	"lms/lms_getway/handlars"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/lib/pq"
)

// @license.name	Apache 2.0
func main() {

	cfg := config.Load()
	if cfg.Environment != "development" {
		gin.SetMode(gin.ReleaseMode)
	}

	docs.SwaggerInfo.Title = cfg.App
	docs.SwaggerInfo.Version = cfg.AppVersion

	fmt.Println("----->>")
	fmt.Printf("%+v\n", cfg)
	fmt.Println("---->>")

	r := gin.New()

	if cfg.Environment != "production" {
		r.Use(gin.Logger(), gin.Recovery()) // Later they will be replaced by custom Logger and Recovery
	}

	r.GET("/ping", MyCORSMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	grpcClients, err := clients.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	defer grpcClients.Close()

	h := handlars.NewHandler(cfg, grpcClients)
	v1 := r.Group("v1")
	{
		v1.Use(MyCORSMiddleware())
		v1.POST("/author", h.CreatAuthor)
		v1.GET("/author/:id", h.GetAuthorByID)
		v1.GET("/author", h.GetAuthorList)
		v1.DELETE("/author/:id", h.DeleteAuthor)
		v1.PUT("/author/:id", h.EnabledAuthor)

		v1.POST("/category", h.CreatCategory)
		v1.GET("/category/:id", h.GetCategoryByID)
		v1.GET("/category", h.GetCategoryList)
		v1.PUT("/category", h.CategoryUpdate)
		v1.DELETE("/category/:id", h.DeleteCategory)
		v1.PUT("/category/:id", h.EnabledCategory)

		v1.POST("/location", h.CreatLocation)
		v1.GET("/location/:id", h.GetLocationByID)
		v1.GET("/location", h.GetLocationList)
		v1.PUT("/location", h.LocationUpdate)
		v1.DELETE("/location/:id", h.DeleteLocation)
		v1.PUT("/location/:id", h.EnabledLocation)

		v1.POST("/book", h.CreatBook)
		v1.GET("/book/:id", h.GetBookByID)
		v1.GET("/book", h.GetBookList)
		v1.PUT("/book", h.BookUpdate)
		v1.DELETE("/book/:id", h.DeleteBook)
		v1.PUT("/book/:id", h.EnabledBook)

		v1.POST("/rental", h.CreatRental)
		v1.GET("/rental/:id", h.GetRentalByID)
		v1.GET("/rental", h.GetRentalList)
		v1.PUT("/rental", h.UpdateRental)
		v1.DELETE("/rental/:id", h.DeleteRental)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(cfg.HTTPPort) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func MyCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("MyCORSMiddleware...")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()

	}
}
