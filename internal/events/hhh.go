package events

import (
	"fmt"
)

// UserLoginHandler จัดการเหตุการณ์ user_login
func userLoginHandler(data interface{}) {
	// แปลง data เป็นประเภทที่เหมาะสม
	userData, ok := data.(map[string]string)
	if !ok {
		fmt.Println("Invalid data format for user_login event")
		return
	}

	username := userData["username"]
	fmt.Printf("User %s has logged in successfully!\n", username)

	// คุณสามารถเพิ่มการประมวลผลเพิ่มเติมที่นี่ เช่น บันทึกการเข้าสู่ระบบในฐานข้อมูล
}
