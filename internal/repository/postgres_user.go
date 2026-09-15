package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing/quick"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/farhanalimohammadi/Go-Commerce-Feri/internal/model"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) (*PostgresUserRepository , error ) {

	if db == "" {
		return nil , errors.New("db cannot be null")
	}

	if err := db.Ping() , err != nil {
		return nil , fmt.Errorf("failed to ping database: %w", err))
	}

	return &PostgresUserRepository{
		db: db,
	}
}

func (r *PostgresUserRepository) Create(
	ctx context.Context,
	user *model.User,
) error {
	const query = `
		INSERT INTO users (
			email,
			password_hash,
			role
		)
		VALUES ($1, $2, $3)
		RETURNING id, created_at, updated_at
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		user.Email,
		user.PasswordHash,
		user.Role,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err == nil {
		return nil
	}

	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) && pgErr.Code == "23505" {
		return ErrConflict
	}

	return fmt.Errorf("create user: %w", err)
}

func (r *PostgresUserRepository) GetByID(
	ctx context.Context,
	id int64
) (*model.User , error) {
	const query = `
		SELECT
			id,
			email,
			password_hash,
			role,
			created_at,
			updated_at
		FROM users
		WHERE id = $1
	`

	var user model.User
	err := r.db.QueryRowContext(
		ctx,
		query,
		id
	).Scan(
		&user.ID
		&user.Email,
		&user.PasswordHash,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt
	)

	if errors.Is(err , ErrNotFound) {
		return nil , ErrNotFound
	}
	if err != nil {
		return nil , fmt.Errorf("get user by id: %w", err)
	}

	return &user , nil
}

func (r *PostgresUserRepository) GetByEmail(
	ctx context.Context,
	email string,
) (*model.User , error) {
	const query = `
		SELECT
			id,
			email.
			password_hash,
			role,
			created_at,
			updated_at
		FROM users
		WHERE email = $1
	`

	var user model.User

	err := r.db.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt
	)

	if errors.Is(err , ErrNotFound) {
		return nil , ErrNotFound
	}
	if err != nil {
		return nil , fmt.Errorf("get user by email: %w", err)
	}

	return &user , nil
}