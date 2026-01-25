package models

import "time"

type User struct {
	ID string `gorm:"primarykey; type:uuid; default:uuid_generate_v4()" json:"id,omitempty" valid:"-"`

	RoleID *string `gorm:"type:varchar(50);default:null" json:"role_id,omitempty"`
	Role   *Role   `gorm:"foreignKey:RoleID;references:ID"`

	Name          string    `gorm:"not null" json:"name,omitempty" valid:"type(string), required~Name is required"`
	Username      string    `json:"username,omitempty" valid:"type(string)"`
	Email         string    `gorm:"not null;unique" json:"email" valid:"email, required~Email is required"`
	Password      string    `gorm:"not null" json:"password,omitempty" valid:"type(string), required~Password is required"`
	Image         string    `json:"image,omitempty" valid:"type(string)"`
	PhoneNumber   string    `json:"phone_number,omitempty" valid:"type(string)"`
	OTPCode       string    `json:"otp_code,omitempty" valid:"type(string)"`
	OTPExpiration time.Time `json:"otp_expiration,omitempty" valid:"-"`
	IsVerified    bool      `gorm:"default: false" json:"is_verified,omitempty" valid:"-"`
	IsSuperUser bool      `gorm:"column:is_superuser;default: false" json:"is_superuser,omitempty" valid:"-"`
	CreatedAt     time.Time `json:"created_at"  valid:"-"`
	UpdatedAt     time.Time `json:"updated_at" valid:"-"`
}

type Users []User

func (User) TableName() string {
	return "users"
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type VerifyAccountRequest struct {
	Email   string `json:"email"`
	OTPCode string `json:"otp_code"`
}

type SendNewOTPRequest struct {
	Email string `json:"email"`
}

type UpdateProfileRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Image       string `json:"image"`
	PhoneNumber string `json:"phone_number"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password"`
}

type CreateNewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type ResetPasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}