package main

import (
	"context"
	"database/sql"
	"log"
)

type UserService struct {
	userRepository UserRepository
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		userRepository: NewUserRepository(db),
	}
}

func (u *UserService) Create(ctx context.Context) error {

	user, err := u.userRepository.AddUserPoint(ctx, "ID-1234", 10)
	if err != nil {
		return err
	}

	log.Println(user)

	return nil
}
