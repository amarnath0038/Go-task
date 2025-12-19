package service

import (
	"context"
	"time"

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
