package business

import "time"

type RefreshTokenResult struct {
	ID        uint
	Token     string
	ExpiresAt time.Time
	UserID    uint
	Type      string
}
