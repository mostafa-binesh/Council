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
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"type:varchar(100);not null"`
	PhoneNumber  string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password     string `gorm:"type:varchar(100);not null"`
	Role         uint   `gorm:"default:1;not null"` // 1: normal user, 2: moderator, 3: admin
	PersonalCode string `gorm:"type:varchar(10);uniqueIndex"`
	NationalCode string `gorm:"type:varchar(10);uniqueIndex"`
	// Provider  *string    `gorm:"type:varchar(50);default:'local';not null"`
	// Photo     *string    `gorm:"not null;default:'default.png'"`
	Verified  bool      `gorm:"not null;default:false"`
	CreatedAt *time.Time `gorm:"not null;default:now()"`
	UpdatedAt *time.Time `gorm:"not null;default:now()"`
}
type MinUser struct {
	ID           uint   `json:"Id,omitempty"`
	Name         string `json:"Name"`
	PhoneNumber  string `json:"PhoneNumber"`
	PersonalCode string `json:"PersonalCode"`
	NationalCode string `json:"NationalCode"`
}

// ! this model has been used in signup handler
type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	PhoneNumber     string `json:"PhoneNumber" validate:"required"`
	Password        string `json:"password" validate:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8,eqfield=Password"`
	PersonalCode    string `json:"PersonalCode" validate:"required,max=8"`
	NationalCode    string `json:"NationalCode" validate:"required,len=10"`
	// Photo string `json:"photo"`
}

// ! this model has been used in Edit user handler
type EditInput struct {
	Name         string `json:"name" validate:"required"`
	PhoneNumber  string `json:"phoneNumber" validate:"required"`
	PersonalCode string `json:"personalCode" validate:"required,max=10,numeric"`
	NationalCode string `json:"nationalCode" validate:"required,len=10,numeric"`
	Password     string `json:"password"`
	// Photo string `json:"photo"`
}

// ! this model has been used in login handler
type SignInInput struct {
	PersonalCode string `json:"personal_code" validate:"required"`
	Password     string `json:"password" validate:"required"`
}

// ! not been used
type UserResponse struct {
	ID          uint      `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Role        string    `json:"role,omitempty"`
	Photo       string    `json:"photo,omitempty"`
	Provider    string    `json:"provider"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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
