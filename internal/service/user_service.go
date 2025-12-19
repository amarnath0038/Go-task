package service

import (
	"context"
	"time"

	"user-crud/internal/models"
	"user-crud/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func CalculateAge(dob time.Time) int {
	now := time.Now()

	age := now.Year() - dob.Year()

	// Birthday not reached yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob time.Time,
) error {

	_, err := s.repo.CreateUser(ctx, name, dob)
	return err
}

func (s *UserService) GetUser(
	ctx context.Context,
	id int64,
) (*models.UserResponse, error) {

	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	dob := user.Dob.Time
	age := CalculateAge(dob)

	return &models.UserResponse{
		ID:   int64(user.ID),
		Name: user.Name,
		Dob:  dob.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *UserService) ListUsers(
	ctx context.Context,
) ([]models.UserResponse, error) {

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	response := make([]models.UserResponse, 0, len(users))

	for _, user := range users {
		dob := user.Dob.Time
		age := CalculateAge(dob)

		response = append(response, models.UserResponse{
			ID:   int64(user.ID),
			Name: user.Name,
			Dob:  dob.Format("2006-01-02"),
			Age:  age,
		})
	}

	return response, nil
}
