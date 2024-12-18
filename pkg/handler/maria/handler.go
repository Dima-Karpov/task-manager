package handler

import (
	"github.com/gin-gonic/gin"
	"task-manager/pkg/service/maria"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.singIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		posts := api.Group("/posts")
		{
			posts.POST("/", h.createPost)
			posts.GET("/", h.getAllPosts)
			posts.DELETE("/:id", h.deletePost)
			posts.GET("/:id", h.getPostByIs)
			posts.PUT("/:id", h.updatePost)
		}
	}

	return router
}
