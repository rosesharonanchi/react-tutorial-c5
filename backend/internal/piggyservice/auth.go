package piggyservice

import (
	"context"
	"errors"


	"golang.org/x/crypto/bcrypt"
	"piggy.com/internal/db/sqlc"
	"piggy.com/internal/models"
)


type AuthService struct {
	repo *sqlc.Queries
}
func NewAuthService(repo *sqlc.Queries) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(ctx context.Context, name, email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := s.repo.CreateUser(ctx, sqlc.CreateUserParams{
		Name:         name,
		Email:        email,
		PasswordHash: string(hashedPassword),
	})
	if err != nil {
		return nil, err
	}

	return &models.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}


func (s *AuthService) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

return &models.User{ID: user.ID, Name: user.Name, Email: user.Email}, nil

}
