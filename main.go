package main

import (
	"log"
	"usertask/database"
	h "usertask/handler"
	"usertask/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectMysql()

	r := gin.Default()
	g := r.Group("/usertask")

	// users
	g.GET("/users", h.GetUsers)
	g.GET("/users/:id", h.GetUserById)
	g.POST("/users", h.AddUser)
	g.PUT("/users/:id", h.ModifyUserById)
	g.DELETE("/users/:id", h.RemoveUserById)

	// login
	g.POST("/login", h.Login)

	// tasks
	g.Use(middleware.JwtAuthMiddleware())
	g.GET("/tasks", h.GetTasks)
	g.GET("/tasks/:id", h.GetTaskById)
	g.POST("/tasks", h.AddTask)
	g.PUT("/tasks/:id", h.ModifyTaskById)
	g.DELETE("/tasks/:id", h.RemoveTaskById)

	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
