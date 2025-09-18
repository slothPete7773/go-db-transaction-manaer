package main

import (
	"context"
	"database/sql"
)

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
	GetById(ctx context.Context, id string) (*UserModel, error)
	Create(ctx context.Context, data UserModel) (*UserModel, error)
	Update(ctx context.Context, data UserModel) (*UserModel, error)
	UpdateByID(ctx context.Context, userID string, updateFn func(user *UserModel) (bool, error)) error
}

type userRepository struct{}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetById(ctx context.Context, id string) (*UserModel, error) {
	return nil, nil
}

func (r *userRepository) Create(ctx context.Context, data UserModel) (*UserModel, error) {
	return nil, nil
}

func (r *userRepository) Update(ctx context.Context, data UserModel) (*UserModel, error) {
	return nil, nil
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

type pointRepository struct{}

func NewPointRepository(db *sql.DB) PointRepository {
	return &pointRepository{}
}

func (r *pointRepository) GetByUserId(ctx context.Context, id string) (*PointModel, error) {
	return nil, nil
}

func (r *pointRepository) Create(ctx context.Context, data PointModel) (*PointModel, error) {
	return nil, nil
}

func (r *pointRepository) Update(ctx context.Context, data PointModel) (*PointModel, error) {
	return nil, nil
}

func (r *pointRepository) AddPoint(ctx context.Context, userId string, points int) (*PointModel, error) {
	return nil, nil
}
