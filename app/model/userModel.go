package model

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id" validate:"required,uuid"`
	Name     string    `json:"name" validate:"required,lte=255"`
	Email    string    `json:"email" validate:"required,lte=255"`
	Password string    `json:"password" validate:"required,lte=255"`
	Active   bool      `json:"active"`
} //@name User

type Product struct {
	ID          uuid.UUID `json:"id" validate:"required,uuid"`
	Description string    `json:"description" validate:"required,lte=45"`
	Price       float32   `json:"price" validate:"required"`
	Active      bool      `json:"active"`
	Type        int       `json:"type"`
} //@name Product
