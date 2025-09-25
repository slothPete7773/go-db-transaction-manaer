package updatefn_closure

import (
	"context"
	"database/sql"
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

func (u *UserService) AddPoint(ctx context.Context) error {

	return u.userRepository.UpdateByID(ctx, "ID-123", func(user *UserModel) (bool, error) {
		err := user.AddUserPoint(10)
		if err != nil {
			return false, err
		}

		return true, nil
	})
}
