package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"workk-telegrambot/dto"
)

type AuthService struct {
	BaseURL string
}

func NewAuthService(baseURL string) *AuthService {
	return &AuthService{BaseURL: baseURL}
}

func (s *AuthService) Auth(req dto.AuthRequest) (*dto.AuthResponse, error) {
	data, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/auth", s.BaseURL), "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var authResp dto.AuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&authResp); err != nil {
		return nil, err
	}
	return &authResp, nil
}
