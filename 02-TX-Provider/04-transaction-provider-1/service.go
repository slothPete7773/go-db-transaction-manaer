package tx_providerr

import (
	"context"
	"database/sql"
)

type UserService struct {
	dbTransactor    *TransactionProvider
	userRepository  UserRepository
	pointRepository PointRepository
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		dbTransactor:    NewTransactionProvider(db),
		userRepository:  NewUserRepository(db),
		pointRepository: NewPointRepository(db),
	}
}

func (u *UserService) AddPoint(ctx context.Context, userId string, points int) error {
	return u.dbTransactor.Transact(func(adapters Repository) error {

		user, err := adapters.UserRepository.GetById(ctx, userId)
		if err != nil {
			return err
		}

		_, err = adapters.PointRepository.AddPoint(ctx, userId, points)
		if err != nil {
			return err
		}

		adapters.UserRepository.Update(ctx, *user)

		return nil
	})
}
