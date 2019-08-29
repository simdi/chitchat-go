package structs

import "time"

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserId    int
	CreatedAt time.Time
}
