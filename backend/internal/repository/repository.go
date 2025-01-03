package repository

import (
	"context"
	"github.com/Serbroda/bookmark-manager/internal/models"
)

type Repository interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	GetUserByUsername(ctx context.Context, username string) (models.User, error)

	CreateBookmark(ctx context.Context, bookmark models.Bookmark) (models.Bookmark, error)
	GetAllBookmarks(ctx context.Context) ([]models.Bookmark, error)
	GetBookmarkByID(ctx context.Context, id string) (models.Bookmark, error)

	CreateSpace(ctx context.Context, bookmark models.Space) (models.Space, error)
}
