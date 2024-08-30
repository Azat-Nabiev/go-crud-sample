package services

import (
	"awesomeProject/internal/models"
	"awesomeProject/internal/repositories"
	"context"
	"go.uber.org/zap"
	"time"
)

type UserService struct {
	userRepository repositories.UserRepository
	logger         *zap.SugaredLogger
}

func NewUserService(userRepository repositories.UserRepository, logger *zap.SugaredLogger) *UserService {
	return &UserService{
		userRepository: userRepository,
		logger:         logger,
	}
}

func (h *UserService) GetAllUserBooks(ctx context.Context) ([]models.User, error) {
	users, err := h.userRepository.GetAllUserBooks(ctx)

	if err != nil {
		h.logger.Errorw("Error on service layer", "error", err.Error())
		return nil, err
	}

	return users, nil
}

func (h *UserService) GetAll(ctx context.Context) ([]models.User, error) {
	users, err := h.userRepository.GetAllUsers(ctx)

	if err != nil {
		h.logger.Errorw("Error on service layer",
			"error", err.Error())
		return nil, err
	}

	return users, nil
}

func (h *UserService) GetById(ctx context.Context, id int) (*models.User, error) {
	user, err := h.userRepository.GetUserByID(ctx, id)

	if err != nil {
		h.logger.Errorw("Error on service layer",
			"error", err.Error(),
			"userId", id)
		return nil, err
	}

	return user, nil
}

func (h *UserService) Add(ctx context.Context, user *models.User) (*models.User, error) {
	user.Since = time.Now()
	createdUser, err := h.userRepository.CreateUser(ctx, user)

	if err != nil {
		h.logger.Errorw("Error on service layer",
			"error", err.Error(),
			"user", user.Name)
	}
	return createdUser, nil
}

func (h *UserService) DeleteByID(ctx context.Context, id int) (*models.User, error) {
	user, err := h.userRepository.DeleteUser(ctx, id)

	if err != nil {
		h.logger.Errorw("Error on service layer",
			"error", err.Error(),
			"userId", id)
		return nil, err
	}

	return user, nil
}

func (h *UserService) Update(ctx context.Context, id int, user *models.User) (*models.User, error) {
	user, err := h.userRepository.UpdateUser(ctx, id, user)

	if err != nil {
		h.logger.Errorw("Error on service layer",
			"error", err.Error(),
			"userId", id)
		return nil, err
	}

	return user, nil
}
