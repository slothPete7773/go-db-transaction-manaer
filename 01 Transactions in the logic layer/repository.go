package main

import (
	"context"
	"database/sql"
)

// Model

type UserModel struct{}

type PointModel struct{}

// Main Repository

type Repository struct {
	UserRepository UserRepository
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		UserRepository: NewUserRepository(db),
	}
}

// User Repository

type UserRepository interface {
	GetById(ctx context.Context, tx *sql.Tx, id string) (*UserModel, error)
	Create(ctx context.Context, tx *sql.Tx, data UserModel) (*UserModel, error)
	Update(ctx context.Context, tx *sql.Tx, data UserModel) (*UserModel, error)
}

type userRepository struct{}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetById(ctx context.Context, tx *sql.Tx, id string) (*UserModel, error) {
	return nil, nil
}

func (r *userRepository) Create(ctx context.Context, tx *sql.Tx, data UserModel) (*UserModel, error) {
	return nil, nil
}

func (r *userRepository) Update(ctx context.Context, tx *sql.Tx, data UserModel) (*UserModel, error) {
	return nil, nil
}

// Point repository
type PointRepository interface {
	GetByUserId(ctx context.Context, tx *sql.Tx, id string) (*PointModel, error)
	Create(ctx context.Context, tx *sql.Tx, data PointModel) (*PointModel, error)
	Update(ctx context.Context, tx *sql.Tx, data PointModel) (*PointModel, error)
	AddPoint(ctx context.Context, tx *sql.Tx, userId string, points int) (*PointModel, error)
}

type pointRepository struct{}

func NewPointRepository(db *sql.DB) PointRepository {
	return &pointRepository{}
}

func (r *pointRepository) GetByUserId(ctx context.Context, tx *sql.Tx, id string) (*PointModel, error) {
	return nil, nil
}

func (r *pointRepository) Create(ctx context.Context, tx *sql.Tx, data PointModel) (*PointModel, error) {
	return nil, nil
}

func (r *pointRepository) Update(ctx context.Context, tx *sql.Tx, data PointModel) (*PointModel, error) {
	return nil, nil
}

func (r *pointRepository) AddPoint(ctx context.Context, tx *sql.Tx, userId string, points int) (*PointModel, error) {
	return nil, nil
}
