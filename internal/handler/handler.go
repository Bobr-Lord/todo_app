package handler

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/petprojects9964409/todo_app/internal/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api")
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.DELETE("/:id", h.deleteList)
			lists.PUT("/:id", h.updateList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.DELETE("/:item_id", h.deleteItem)
				items.PUT("/:item_id", h.updateItem)
			}
		}
	}

	return router
}
