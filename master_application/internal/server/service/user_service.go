package server

import (
	"context"
	"database/sql"

	"github.com/pubudu2003060/proxy_project/internal/db/repository"
	models "github.com/pubudu2003060/proxy_project/internal/server/models"
)

type UserService interface {
	CreateUser(ctx context.Context, params models.CreateUserRequest) (*models.CreateUserResponse, error)
}

type userService struct {
	queries repository.Querier
	db      *sql.DB
}

func NewUserService(q repository.Querier, db *sql.DB) UserService {
	return &userService{
		queries: q,
		db:      db,
	}
}

func (s *userService) CreateUser(ctx context.Context, params models.CreateUserRequest) (*models.CreateUserResponse, error) {
	user, err := s.queries.CreateUser(ctx, repository.CreateUserParams{
		Email:    params.Email,
		Username: params.Username,
		Password: params.Password,
	})
	if err != nil {
		return nil, err
	}

	return &models.CreateUserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}
