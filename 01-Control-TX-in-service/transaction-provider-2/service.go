package tx_injection

import (
	"context"
	"database/sql"
	"log"
)

type UserService struct {
	dbTransactor    DBTransactor
	userRepository  UserRepository
	pointRepository PointRepository
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		dbTransactor:    NewTransactor(db),
		userRepository:  NewUserRepository(db),
		pointRepository: NewPointRepository(db),
	}
}

func (u *UserService) AddPoint(ctx context.Context, userId string, points int) error {

	tx := u.dbTransactor.Begin()
	defer tx.SafeRollback(recover())

	user, err := u.userRepository.GetById(tx, ctx, userId)
	if err != nil {
		return err
	}

	point, err := u.pointRepository.AddPoint(tx, ctx, userId, points)
	if err != nil {
		return err
	}

	u.userRepository.Update(tx, ctx, *user)
	log.Println("Point: ", point)

	tx.Commit()

	return nil
}
