package user

import "time"

type UserModel struct {
	ID        int       `json:"id"`
	Username  string    `json:"username" validate:"required"`
	Fullname  string    `json:"fullname" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	RoleID    int       `json:"role_id"`
	RoleName  string    `json:"role_name"`
	Active    bool      `json:"active"`
	ImageName string    `json:"image_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginModel struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	RoleID   int    `json:"role_id"`
	Active   bool   `json:"active"`
}

type ResponseModel struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserProfileModel struct {
	Username  string `json:"username" validate:"required"`
	Fullname  string `json:"fullname" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	RoleName  string `json:"role_name" validate:"required"`
	ImageName string `json:"image_name"`
}

type TokenRequestModel struct {
	RefreshToken string `json:"refresh_token"`
}
