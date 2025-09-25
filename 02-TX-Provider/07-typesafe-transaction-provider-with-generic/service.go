package typesafetransactionproviderwithgeneric

import (
	"context"
	"fmt"
)

type Service[T repoTx] struct {
	tr transactor[T]
}

func NewService[T repoTx](tr transactor[T]) *Service[T] {
	return &Service[T]{tr}
}

func (s *Service[T]) Create(ctx context.Context, name string, items []string) error {
	err := s.tr.InTx(ctx, func(r T) error {
		if err := r.CreateUser(ctx, name); err != nil {
			return fmt.Errorf("create user: %w", err)
		}
		if err := r.CreateOrder(ctx, items); err != nil {
			return fmt.Errorf("create order: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("create user & order: %w", err)
	}
	return nil
}

// type UserService struct {
// 	userRepository  UserRepository
// 	pointRepository PointRepository
// }

// func NewUserService(db *sql.DB) UserService {
// 	return UserService{
// 		userRepository:  NewUserRepository(db),
// 		pointRepository: NewPointRepository(db),
// 	}
// }

// func (u *UserService) AddPoint(ctx context.Context, userId string, points int) error {
// 	user, err := u.userRepository.GetById(ctx, userId)
// 	log.Println("err:", err)
// 	if err != nil {
// 		return err
// 	}

// 	_, err = u.pointRepository.AddPoint(ctx, userId, points)
// 	// err = errors.New("Hee")
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("4")
// 	_, err = u.userRepository.Update(ctx, *user)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
