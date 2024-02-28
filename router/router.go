package router

import (
	"net/http"

	"github.com/aungkoko1234/gin-crud/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(tagController *controller.TagController) *gin.Engine {
	service := gin.Default()

	service.GET("",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,"Welcome Home")
	})

	service.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	router := service.Group("/api")
	tagRouter := router.Group("/tags")
	tagRouter.GET("",tagController.FindAll)
	tagRouter.GET("/:tagId", tagController.FindById)
	tagRouter.POST("", tagController.Create)
	tagRouter.PATCH("/:tagId", tagController.Update)
	tagRouter.DELETE("/:tagId", tagController.Delete)

	return service
}