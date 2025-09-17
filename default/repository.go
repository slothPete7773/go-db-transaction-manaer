package main

import (
	"context"
	"database/sql"
)

// Model

type UserModel struct{}

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
	Delete(ctx context.Context, id string, deletedBy string) error
	Search(ctx context.Context, keyword string, offset, pageSize int) ([]UserModel, error)
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
func (r *userRepository) Delete(ctx context.Context, id string, deletedBy string) error {
	return nil
}
func (r *userRepository) Search(ctx context.Context, keyword string, offset, pageSize int) ([]UserModel, error) {
	return nil, nil
}
