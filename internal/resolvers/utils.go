package resolvers

import (
	"github.com/google/uuid"
)

// generateID создает уникальный идентификатор
func generateID() string {
	return uuid.New().String()
}
