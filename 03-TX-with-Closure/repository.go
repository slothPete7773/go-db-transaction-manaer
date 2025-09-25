package updatefn_closure

import (
	"context"
	"database/sql"
)

// Main Repository

type Repository struct {
	UserRepository  UserRepository
	PointRepository PointRepository
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		UserRepository:  NewUserRepository(db),
		PointRepository: NewPointRepository(db),
	}
}

// User Repository

type UserRepository interface {
	GetById(ctx context.Context, id string) (*UserModel, error)
	Create(ctx context.Context, data UserModel) (*UserModel, error)
	Update(ctx context.Context, data UserModel) (*UserModel, error)
	UpdateByID(ctx context.Context, userID string, updateFn func(user *UserModel) (bool, error)) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetById(ctx context.Context, id string) (*UserModel, error) {
	query := "SELECT id, email, points FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var user UserModel
	var points int
	err := row.Scan(&user.id, &user.email, &points)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	user.point = PointModel{
		points: points,
		userID: user.id,
	}

	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, data UserModel) (*UserModel, error) {
	query := "INSERT INTO users (id, email, points) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, data.id, data.email, data.point.points)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *userRepository) Update(ctx context.Context, data UserModel) (*UserModel, error) {
	query := "UPDATE users SET email = $1, points = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, data.email, data.point.points, data.id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *userRepository) UpdateByID(ctx context.Context, userID string, updateFn func(user *UserModel) (bool, error)) error {
	// DO: 1
	user, err := r.GetById(ctx, userID)
	if err != nil {
		return err
	}

	// Update user field.

	ok, err := updateFn(user)
	if !ok && err != nil {
		return err
	}

	_, err = r.Update(ctx, *user)
	if err != nil {
		return err
	}

	// DO: 3

	return nil
}

// Point repository
type PointRepository interface {
	GetByUserId(ctx context.Context, id string) (*PointModel, error)
	Create(ctx context.Context, data PointModel) (*PointModel, error)
	Update(ctx context.Context, data PointModel) (*PointModel, error)
	AddPoint(ctx context.Context, userId string, points int) (*PointModel, error)
}

type pointRepository struct {
	db *sql.DB
}

func NewPointRepository(db *sql.DB) PointRepository {
	return &pointRepository{db: db}
}

func (r *pointRepository) GetByUserId(ctx context.Context, id string) (*PointModel, error) {
	query := "SELECT id, points, user_id FROM points WHERE user_id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var point PointModel
	err := row.Scan(&point.id, &point.points, &point.userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &point, nil
}

func (r *pointRepository) Create(ctx context.Context, data PointModel) (*PointModel, error) {
	query := "INSERT INTO points (id, points, user_id) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, data.id, data.points, data.userID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *pointRepository) Update(ctx context.Context, data PointModel) (*PointModel, error) {
	query := "UPDATE points SET points = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, data.points, data.id)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *pointRepository) AddPoint(ctx context.Context, userId string, points int) (*PointModel, error) {
	existingPoint, err := r.GetByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	if existingPoint != nil {
		existingPoint.points += points
		return r.Update(ctx, *existingPoint)
	}

	newPoint := PointModel{
		id:     userId + "-points",
		points: points,
		userID: userId,
	}
	return r.Create(ctx, newPoint)
}
