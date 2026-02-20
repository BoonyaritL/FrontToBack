package main

import (
	"go-backend/controllers"  // เรียก Controller
	"go-backend/initializers" // เรียก DB
	
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables() 
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Config CORS (เหมือนเดิม)
	// ...
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                     // พอร์ตของ React
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}, // ต้องมี OPTIONS ด้วย!
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	// Routes: เรียกใช้ function จาก package controllers
	r.GET("/todos", controllers.GetTodos)
	r.POST("/todos", controllers.CreateTodo)
	r.PATCH("/todos/:id", controllers.UpdateTodo)
	r.DELETE("/todos/:id", controllers.DeleteTodo)

	r.Run(":3000")
}
