package models

// ชื่อ Struct ต้องตัวใหญ่ (Todo) ถึงจะเรียกใช้ข้ามไฟล์ได้
type Todo struct {
	ID        uint   `json:"id" gorm:"primaryKey"` // เติม `json:"id"`
	Title     string `json:"title"`                // เติม `json:"title"`
	Completed bool   `json:"completed"`            // เติม `json:"completed"`
}

//// models/hotelModels.go

// type Room struct {
//     ID          uint    `json:"id" gorm:"primaryKey"`
//     RoomNumber  string  `json:"room_number"`
//     RoomType    string  `json:"room_type"` // เช่น Single, Double, Suite
//     Price       float64 `json:"price"`
//     IsAvailable bool    `json:"is_available" gorm:"default:true"`
// }

// type Booking struct {
//     ID          uint      `json:"id" gorm:"primaryKey"`
//     RoomID      uint      `json:"room_id"`
//     Room        Room      `gorm:"foreignKey:RoomID"` // บอก GORM ว่าเชื่อมกับตาราง Room
//     GuestName   string    `json:"guest_name"`
//     CheckIn     time.Time `json:"check_in"`
//     CheckOut    time.Time `json:"check_out"`
//     TotalAmount float64   `json:"total_amount"`
// }