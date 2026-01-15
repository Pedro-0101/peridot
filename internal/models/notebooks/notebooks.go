package notebooks

import (
	"time"

	"github.com/google/uuid"
)

type Notebook struct {
	ID        uuid.UUID
	Name      string
	UserID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}
