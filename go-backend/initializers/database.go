package initializers

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ประกาศตัวแปร Global ให้คนอื่นใช้ (ตัว D ใหญ่ = Public)
var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=todo_db port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("เชื่อมต่อ Database ไม่สำเร็จ")
	}
}