package model

import "agro-mart/internal/enum"

type User struct {
	Id   *string        `json:"id,omitempty"`
	Name *string        `json:"name,omitempty"`
	Role *enum.UserRole `json:"role,omitempty"`
}
