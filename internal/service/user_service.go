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
	page int,
	limit int,
) ([]models.UserResponse, error) {

	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit

	users, err := s.repo.ListUsers(
		ctx,
		int32(limit),
		int32(offset),
	)
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

func (s *UserService) UpdateUser(
	ctx context.Context,
	id int64,
	name string,
	dob time.Time,
) (*models.UserResponse, error) {

	user, err := s.repo.UpdateUser(ctx, id, name, dob)
	if err != nil {
		return nil, err
	}

	age := CalculateAge(user.Dob.Time)

	return &models.UserResponse{
		ID:   int64(user.ID),
		Name: user.Name,
		Dob:  user.Dob.Time.Format("2006-01-02"),
		Age:  age,
	}, nil
}

func (s *UserService) DeleteUser(
	ctx context.Context,
	id int64,
) error {
	return s.repo.DeleteUser(ctx, id)
}
