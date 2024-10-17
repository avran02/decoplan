package dto

type LoginRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RegisterRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type RegisterResponseDTO struct {
	Success bool `json:"success"`
}

type RefreshTokensRequestDTO struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokensResponseDTO struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LogoutRequestDTO struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type LogoutResponseDTO struct {
	Success bool `json:"success"`
}
