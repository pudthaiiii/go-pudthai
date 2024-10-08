package events

import (
	"go-pudthai/internal/infrastructure/cache"
	"go-pudthai/internal/infrastructure/mailer"

	"gorm.io/gorm"
)

var (
	mail *mailer.Mailer
	// db    *gorm.DB
//  cacheManager *cache.CacheManager
)

type EventListener chan Event

type Event struct {
	Name string
	Data interface{}
}

func NewEventListener(mailer *mailer.Mailer, dbConn *gorm.DB, cache *cache.CacheManager) EventListener {
	mail = mailer
	// cacheManager = cache
	// db = dbConn

	return make(EventListener)
}

func (el EventListener) Listen() {
	for event := range el {
		switch event.Name {
		case "user.created":
		case "merchant.created":
		case "admin.created":
			userCreated(event.Data)
		}
	}
}

func Emit(el EventListener, name string, data interface{}) {
	el <- Event{Name: name, Data: data}
}
