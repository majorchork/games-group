package server

import (
	"github.com/gin-gonic/gin"
	"github.com/majorchork/tech-crib-africa/internal/controller"
	"github.com/majorchork/tech-crib-africa/internal/middleware"
	"log"
)

func DefineRoutes(handler *controller.Handler) *gin.Engine {
	log.Println("Routes defined")

	router := gin.Default()

	r := router.Group("/api/v1")
	{
		r.GET("/ping", handler.Ping)
		r.POST("/signup", handler.SignUp)
		r.POST("/login", handler.Login)
	}

	authorized := r.Use(middleware.AuthorizeUser(handler))
	{
		authorized.GET("/admin/adminProfile", handler.AdminProfile)
		authorized.GET("/user/group", handler.GetGuestsByGroup)
		authorized.GET("/user/guests", handler.GetGuests)
		authorized.GET("/user/guest/profile", handler.GuestProfile)
		authorized.POST("/user/saveGuests", handler.AssignGroupsAndSaveGuests)

	}

	return router
}

func SetupRouter(h *controller.Handler) *gin.Engine {
	log.Println("Router setup")
	r := DefineRoutes(h)

	return r
}
