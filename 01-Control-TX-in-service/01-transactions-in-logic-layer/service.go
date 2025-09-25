package main

import (
	"context"
	"database/sql"
	"log"
)

type UserService struct {
	userRepository  UserRepository
	pointRepository PointRepository
	db              *sql.DB
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		userRepository:  NewUserRepository(db),
		pointRepository: NewPointRepository(db),
		db:              db,
	}
}

func (u *UserService) Create(ctx context.Context) error {
	return runInTx(u.db, func(tx *sql.Tx) error {
		user, err := u.userRepository.GetById(ctx, tx, "ID-123")
		if err != nil {
			return err
		}

		point, err := u.pointRepository.AddPoint(ctx, tx, "ID-123", 10)
		if err != nil {
			return err
		}

		u.userRepository.Update(ctx, tx, *user)

		log.Println("Point: ", point)

		return nil

	})

}
