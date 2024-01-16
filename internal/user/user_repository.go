package user

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := "INSERT INTO users(name, email, password, company_id) VALUES (?, ?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, user.Name, user.Email, user.Password, user.Company_ID)
	if err != nil {
		return &User{}, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}
	user.ID = int64(lastInsertId)
	return user, nil

}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}

	query := "SELECT id, name, email, password, company_id FROM users WHERE email = ?"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Company_ID)
	if err != nil {
		return &User{}, nil
	}
	return &u, nil
}
