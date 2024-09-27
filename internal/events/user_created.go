package events

import (
	"go-ibooking/internal/entities"

	"github.com/jinzhu/copier"
)

func userCreated(data interface{}) {
	var (
		user = entities.User{}
	)

	copier.Copy(&user, data)

	mail.Send("ยินดีตอนรับเข้าสู่ระบบ Demo Golang", "user_created", user, user.Email)
}
