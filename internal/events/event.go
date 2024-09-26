package events

import "go-ibooking/internal/infrastructure/mailer"

type Event struct {
	Name string
	Data interface{}
}

var mail *mailer.Mailer

// EventListener เป็น channel ที่ใช้รับเหตุการณ์
type EventListener chan Event

// NewEventListener สร้าง event listener ใหม่
func NewEventListener(mailer *mailer.Mailer) EventListener {
	mail = mailer
	return make(EventListener)
}

// Listen ฟังก์ชันที่รอรับเหตุการณ์
func (el EventListener) Listen() {
	for event := range el {
		switch event.Name {
		case "create_user_email":
			createUserEmail(event.Data)
		default:
			// fmt.Printf("Received event: %s with data: %v\n", event.Name, event.Data)
		}
	}
}

// Emit ส่งเหตุการณ์ไปยัง listener
func Emit(el EventListener, name string, data interface{}) {
	el <- Event{Name: name, Data: data}
}
