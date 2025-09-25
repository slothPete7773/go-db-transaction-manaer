package txdemo

import (
	"context"
	"database/sql"
	"log"
)

type UserService struct {
	txManager       TxManager
	userRepository  UserRepository
	pointRepository PointRepository
}

func NewUserService(db *sql.DB) UserService {
	return UserService{
		txManager:       NewTxManager(db),
		userRepository:  NewUserRepository(db),
		pointRepository: NewPointRepository(db),
	}
}

func (u *UserService) AddPoint(ctx context.Context, userId string, points int) error {
	return u.txManager.Run(ctx, func(ctx context.Context, tx *sql.Tx) error {
		log.Println("1")
		userTx := u.userRepository.WithTransaction(tx)
		pointTx := u.pointRepository.WithTransaction(tx)

		log.Println("2")
		user, err := userTx.GetById(ctx, userId)
		log.Println("err:", err)
		if err != nil {
			return err
		}

		log.Println("3")
		point, err := pointTx.AddPoint(ctx, userId, points)
		// err = errors.New("Hee")
		if err != nil {
			return err
		}

		log.Println("4")
		_, err = userTx.Update(ctx, *user)
		if err != nil {
			return err
		}

		log.Println("Point: ", point)

		return nil
	})
}
