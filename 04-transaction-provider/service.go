package main

import (
	"context"
	"database/sql"
	"log"
)

type UserService struct {
	txProvider      txProvider
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
	return u.txProvider.Transact(func(adapters Adapters) error {
		user, err := adapters.UserRepository.GetById(ctx, "ID-123")
		if err != nil {
			return err
		}

		point, err := adapters.PointRepository.AddPoint(ctx, "ID-123", 10)
		if err != nil {
			return err
		}

		adapters.UserRepository.Update(ctx, *user)

		log.Println("Point: ", point)
		return nil
	})
}
