package users

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string
	Pass      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
