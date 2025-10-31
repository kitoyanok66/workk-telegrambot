package dto

import "github.com/google/uuid"

type AuthRequest struct {
	UserID     uuid.UUID `json:"user_id,omitempty"`
	Provider   string    `json:"provider"`
	ExternalID string    `json:"external_id"`
	Username   string    `json:"username"`
}

type AuthResponse struct {
	Token string   `json:"token"`
	User  *UserDTO `json:"user"`
}
