package resolvers

import (
	"context"

	"github.com/FlynntDev/ozon-intern-task/internal/models"
)

func (r *mutationResolver) AddPost(ctx context.Context, title string, content string, allowComments bool) (*models.Post, error) {
	// Создаем новый пост
	post := &models.Post{
		ID:            generateID(), // Убедись, что функция generateID существует
		Title:         title,
		Content:       content,
		AllowComments: allowComments,
		Comments:      []*models.Comment{},
	}

	// Добавляем пост в in-memory хранилище
	r.Posts = append(r.Posts, post)

	// Возвращаем созданный пост
	return post, nil
}
