package tx_providerr

import (
	"context"
	"database/sql"
)

type db interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// Model

type UserModel struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Points string `json:"points"`
}

type PointModel struct {
	ID     string `json:"id"`
	Points int    `json:"points"`
	UserID string `json:"user_id"`
}

// Main Repository

type Repository struct {
	UserRepository  UserRepository
	PointRepository PointRepository
}

func NewRepository(db db) Repository {
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
}

type userRepository struct {
	db db
}

func NewUserRepository(db db) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetById(ctx context.Context, id string) (*UserModel, error) {
	query := "SELECT id, email, points FROM users WHERE id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var user UserModel
	err := row.Scan(&user.ID, &user.Email, &user.Points)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Create(ctx context.Context, data UserModel) (*UserModel, error) {
	query := "INSERT INTO users (id, email, points) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, data.ID, data.Email, data.Points)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *userRepository) Update(ctx context.Context, data UserModel) (*UserModel, error) {
	query := "UPDATE users SET email = $1, points = $2 WHERE id = $3"
	_, err := r.db.ExecContext(ctx, query, data.Email, data.Points, data.ID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Point repository
type PointRepository interface {
	GetByUserId(ctx context.Context, id string) (*PointModel, error)
	Create(ctx context.Context, data PointModel) (*PointModel, error)
	Update(ctx context.Context, data PointModel) (*PointModel, error)
	AddPoint(ctx context.Context, userId string, points int) (*PointModel, error)
}

type pointRepository struct {
	db db
}

func NewPointRepository(db db) PointRepository {
	return &pointRepository{db: db}
}

func (r *pointRepository) GetByUserId(ctx context.Context, id string) (*PointModel, error) {
	query := "SELECT id, points, user_id FROM points WHERE user_id = $1"
	row := r.db.QueryRowContext(ctx, query, id)

	var point PointModel
	err := row.Scan(&point.ID, &point.Points, &point.UserID)
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
	_, err := r.db.ExecContext(ctx, query, data.ID, data.Points, data.UserID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *pointRepository) Update(ctx context.Context, data PointModel) (*PointModel, error) {
	query := "UPDATE points SET points = $1 WHERE id = $2"
	_, err := r.db.ExecContext(ctx, query, data.Points, data.ID)
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
		existingPoint.Points += points
		return r.Update(ctx, *existingPoint)
	}

	newPoint := PointModel{
		ID:     userId + "-points",
		Points: points,
		UserID: userId,
	}
	return r.Create(ctx, newPoint)
}
