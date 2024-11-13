package models

import "time"

type User struct {
	ID        string    `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`
	Name      string    `gorm:"not null" json:"name,omitempty" valid:"type(string), required~Name is required"`
	Username  string    `json:"username,omitempty" valid:"type(string)"`
	Email     string    `gorm:"not null;unique" json:"email" valid:"email, required~Email is required"`
	Password  string    `gorm:"not null" json:"password,omitempty" valid:"type(string), required~Password is required"`
	CreatedAt time.Time `json:"created_at"  valid:"-"`
	UpdatedAt time.Time `json:"updated_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}

type SignUpRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Name     string `json:"name" validate:"required"`
    Password string `json:"password" validate:"required,min=8"`
}

type SignInRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=8"`
}