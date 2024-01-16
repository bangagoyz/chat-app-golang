package user

import "context"

type User struct {
	ID         int64  `json:"id" db:"id"`
	Email      string `json:"email" db:"email"`
	Name       string `json:"name" db:"name"`
	Password   string `json:"password" db:"password"`
	Company_ID string `json:"company_id" db:"company_id"`
}

type CreateUserReq struct {
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password"`
	Company_ID string `json:"company_id" db:"company_id"`
}

type CreateUserRes struct {
	ID         string `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	Company_ID string `json:"company_id" db:"company_id"`
}

type LoginUserReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserRes struct {
	accessToken string
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
}

type UserMessage struct {
	Name       string `json:"name" db:"name"`
	Email      string `json:"email" db:"email"`
	Company_ID string `json:"company_id" db:"company_id"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error)
}
