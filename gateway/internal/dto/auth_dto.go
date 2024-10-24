package dto

type LoginRequestDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponseDTO struct {
	ID           string `json:"id"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterRequestDTO struct {
	Username string  `json:"username" validate:"required"`
	Password string  `json:"password" validate:"required"`
	Email    *string `json:"email" validate:"email, omitempty"`
}

type RegisterResponseDTO struct {
	ID           string `json:"id"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokensRequestDTO struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type RefreshTokensResponseDTO struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LogoutRequestDTO struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

type LogoutResponseDTO struct {
	Ok bool `json:"ok"`
}
