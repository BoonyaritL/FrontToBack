package initializers

import (
	"log"
	"github.com/joho/godotenv"
)

// ฟังก์ชันนี้มีหน้าที่โหลดค่าจากไฟล์ .env เข้ามาในระบบ
func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}