package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	// "github.com/google/uuid"
)

// ! the model that been used for migration and retrieve and add data to the database
type User struct {
	// ID        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	// gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"type:varchar(100);not null"`
	Role     uint   `gorm:"default:1;not null"`
	// Provider  *string    `gorm:"type:varchar(50);default:'local';not null"`
	// Photo     *string    `gorm:"not null;default:'default.png'"`
	// Verified  *bool      `gorm:"not null;default:false"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}

// ! this model has been used in signup handler
type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required,email,gunique=users.email"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8,eqfield=Password"`
	Photo           string `json:"photo"`
}

// ! this model has been used in login handler
type SignInInput struct {
	Email    string `json:"email"  validate:"required,email"`
	Password string `json:"password"  validate:"required"`
}

// ! not been used
type UserResponse struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Provider  string    `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var validate = validator.New()

// ! not been used
type ErrorResponse struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`
}

// ! not been used
func ValidateStruct[T any](payload T) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(payload)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.Field = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
