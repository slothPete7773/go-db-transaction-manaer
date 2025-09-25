package tx_injection

import (
	"context"
	"database/sql"
)

type DB interface {
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
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

func NewRepository(db *sql.DB) Repository {
	return Repository{
		UserRepository:  NewUserRepository(db),
		PointRepository: NewPointRepository(db),
	}
}

// User Repository

type UserRepository interface {
	GetById(tx *Tx, ctx context.Context, id string) (*UserModel, error)
	Create(tx *Tx, ctx context.Context, data UserModel) (*UserModel, error)
	Update(tx *Tx, ctx context.Context, data UserModel) (*UserModel, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) dbOrTx(tx *Tx) DB {
	if tx != nil && tx.Instance() != nil {
		return tx.Instance()
	}
	return r.db
}

func (r *userRepository) GetById(tx *Tx, ctx context.Context, id string) (*UserModel, error) {
	db := r.dbOrTx(tx)

	query := "SELECT id, email, points FROM users WHERE id = $1"
	row := db.QueryRowContext(ctx, query, id)

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

func (r *userRepository) Create(tx *Tx, ctx context.Context, data UserModel) (*UserModel, error) {
	db := r.dbOrTx(tx)

	query := "INSERT INTO users (id, email, points) VALUES ($1, $2, $3)"
	_, err := db.ExecContext(ctx, query, data.ID, data.Email, data.Points)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *userRepository) Update(tx *Tx, ctx context.Context, data UserModel) (*UserModel, error) {
	db := r.dbOrTx(tx)

	query := "UPDATE users SET email = $1, points = $2 WHERE id = $3"
	_, err := db.ExecContext(ctx, query, data.Email, data.Points, data.ID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// Point repository
type PointRepository interface {
	GetByUserId(tx *Tx, ctx context.Context, id string) (*PointModel, error)
	Create(tx *Tx, ctx context.Context, data PointModel) (*PointModel, error)
	Update(tx *Tx, ctx context.Context, data PointModel) (*PointModel, error)
	AddPoint(tx *Tx, ctx context.Context, userId string, points int) (*PointModel, error)
}

type pointRepository struct {
	db *sql.DB
}

func NewPointRepository(db *sql.DB) PointRepository {
	return &pointRepository{db: db}
}

func (r *pointRepository) dbOrTx(tx *Tx) DB {
	if tx != nil && tx.Instance() != nil {
		return tx.Instance()
	}
	return r.db
}

func (r *pointRepository) GetByUserId(tx *Tx, ctx context.Context, id string) (*PointModel, error) {
	db := r.dbOrTx(tx)

	query := "SELECT id, points, user_id FROM points WHERE user_id = $1"
	row := db.QueryRowContext(ctx, query, id)

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

func (r *pointRepository) Create(tx *Tx, ctx context.Context, data PointModel) (*PointModel, error) {
	db := r.dbOrTx(tx)

	query := "INSERT INTO points (id, points, user_id) VALUES ($1, $2, $3)"
	_, err := db.ExecContext(ctx, query, data.ID, data.Points, data.UserID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *pointRepository) Update(tx *Tx, ctx context.Context, data PointModel) (*PointModel, error) {
	db := r.dbOrTx(tx)

	query := "UPDATE points SET points = $1 WHERE id = $2"
	_, err := db.ExecContext(ctx, query, data.Points, data.ID)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *pointRepository) AddPoint(tx *Tx, ctx context.Context, userId string, points int) (*PointModel, error) {
	// db := r.dbOrTx(tx)

	existingPoint, err := r.GetByUserId(tx, ctx, userId)
	if err != nil {
		return nil, err
	}

	if existingPoint != nil {
		existingPoint.Points += points
		return r.Update(tx, ctx, *existingPoint)
	}

	newPoint := PointModel{
		ID:     userId + "-points",
		Points: points,
		UserID: userId,
	}
	return r.Create(tx, ctx, newPoint)
}
