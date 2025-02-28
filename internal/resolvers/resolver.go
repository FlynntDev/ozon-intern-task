package resolvers

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import "github.com/FlynntDev/ozon-intern-task/internal/models"

type Resolver struct {
	Posts    []*models.Post
	Comments []*models.Comment
}
