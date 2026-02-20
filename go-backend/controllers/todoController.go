package controllers

import (
	"go-backend/initializers" // import ตัวเชื่อม DB เข้ามา
	"go-backend/models"       // import struct เข้ามา
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)

	// เรียกใช้ DB จาก package initializers
	result := initializers.DB.Create(&todo)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{"todo": todo})
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	initializers.DB.Find(&todos)

	// c.JSON(200, gin.H{"todos": todos})
	c.JSON(http.StatusOK, todos)
}

// ... ฟังก์ชัน Update, Delete ก็ทำเหมือนกัน ...
func UpdateTodo(c *gin.Context) {
	// 1. รับ ID จาก URL
	id := c.Param("id")

	// 2. หาข้อมูลเดิมใน DB
	var todo models.Todo
	if err := initializers.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลที่ต้องการอัปเดต"})
		return
	}

	// 3. รับข้อมูลใหม่ที่จะแก้ไข (เช่น completed)
	var input struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "รูปแบบข้อมูลไม่ถูกต้อง"})
		return
	}

	// 4. ทำการอัปเดตเฉพาะ Field ที่ต้องการ
	initializers.DB.Model(&todo).Update("Completed", input.Completed)

	c.JSON(http.StatusOK, todo)
}

// DELETE /todos/:id
func DeleteTodo(c *gin.Context) {
	// 1. รับ ID จาก URL
	id := c.Param("id")

	// 2. ลบข้อมูล (แบบถาวร)
	result := initializers.DB.Delete(&models.Todo{}, id)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลที่ต้องการลบ"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ลบข้อมูลเรียบร้อยแล้ว"})
}


