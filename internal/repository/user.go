package repository

import (
	"context"
	"github.com/farhanalimohammadi/Go-Commerce-Feri/internal/model"
)

type UserRepository interfacec {
	Create(ctx context.Context , user *model.User) error
	GetByID(ctx context.Context , id int64) (*model.User , error)
	GetByEmail(ctx context.Context , email string) (*model.User , error)
}