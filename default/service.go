package main

import (
	"context"
	"database/sql"
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

	u.userRepository.Create(ctx, UserModel{})

	return nil
}
