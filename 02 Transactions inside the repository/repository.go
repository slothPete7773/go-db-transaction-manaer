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
	AddUserPoint(ctx context.Context, id string, points int) (*UserModel, error)
}

type userRepository struct{}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{}
}

func (r *userRepository) AddUserPoint(ctx context.Context, id string, points int) (*UserModel, error) {
	// Do operation to table User and Point here
	return nil, nil
}
