package main

import (
	"context"
	"database/sql"
	"log"
)

type UserService struct {
	userRepository  UserRepository
	pointRepository PointRepository
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		userRepository:  NewUserRepository(db),
		pointRepository: NewPointRepository(db),
	}
}

func (u *UserService) Create(ctx context.Context) error {

	user, err := u.userRepository.GetById(ctx, "ID-123")
	if err != nil {
		return err
	}

	point, err := u.pointRepository.AddPoint(ctx, "ID-123", 10)
	if err != nil {
		return err
	}

	u.userRepository.Update(ctx, *user)

	log.Println("Point: ", point)

	return nil
}
