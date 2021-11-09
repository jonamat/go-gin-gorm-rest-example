package main

import (
	"fmt"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/controllers"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/db"
	"github.com/jonamat/go-gin-gorm-rest-example/pkg/middlewares"
)

func main() {
	// Load envs
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// Create http server
	r := gin.Default()

	// Connect to Postgres & init migrations
	if _, err := db.Connect(); err != nil {
		panic(err)
	}

	// Connect to redis
	store, err := db.CreateSessionStorage()
	if err != nil {
		panic(err)
	}

	// Define routes
	r.Use(sessions.Sessions("auth_session", store))
	r.GET("/ping", controllers.Ping)

	// Auth
	r.POST("/auth/login", controllers.Login)
	r.GET("/auth/logout", controllers.Logout)

	// APIsV1 group setup
	apiV1 := r.Group("/api/v1")
	apiV1.Use(middlewares.Auth)

	// Users
	apiV1.GET("/users/:id", controllers.FindUser)
	apiV1.PUT("/users/:id", controllers.UpdateUser)

	// Bookings
	apiV1.GET("/bookings", controllers.FindBookings)
	apiV1.GET("/bookings/:id", controllers.FindBooking)
	apiV1.POST("/bookings", controllers.CreateBooking)
	apiV1.PUT("/bookings/:id", controllers.UpdateBooking)
	apiV1.DELETE("/bookings/:id", controllers.DeleteBooking)

	// Services
	apiV1.GET("/services", controllers.FindServices)
	apiV1.GET("/services/:id", controllers.FindService)
	apiV1.POST("/services", controllers.CreateService)
	apiV1.PUT("/services/:id", controllers.UpdateService)
	apiV1.DELETE("/services/:id", controllers.DeleteService)

	// Rooms
	apiV1.GET("/rooms", controllers.FindRooms)
	apiV1.GET("/rooms/:id", controllers.FindRoom)
	apiV1.POST("/rooms", controllers.CreateRoom)
	apiV1.PUT("/rooms/:id", controllers.UpdateRoom)
	apiV1.DELETE("/rooms/:id", controllers.DeleteRoom)

	// Start server
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}
