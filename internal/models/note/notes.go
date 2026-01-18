package notes

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID         uuid.UUID
	Title      string
	Content    string
	NotebookID uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
