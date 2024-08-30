package repositories

import (
	"awesomeProject/internal/models"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]models.User, error)
	UpdateUser(ctx context.Context, id int, user *models.User) (*models.User, error)
	DeleteUser(ctx context.Context, id int) (*models.User, error)
	GetAllUserBooks(ctx context.Context) ([]models.User, error)
}
