package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	service "task-manager/pkg/service/postgres"

	_ "task-manager/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("api", h.userIdentity)
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.create)
			posts.GET("/", h.getAll)
			posts.GET("/:id", h.getById)
			posts.DELETE("/:id", h.delete)
			posts.PUT("/:id", h.update)
		}
	}

	return router
}
