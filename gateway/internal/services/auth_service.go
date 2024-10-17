package services

import (
	"context"
	"errors"

	"github.com/avran02/decplan/gateway/pb"
)

type AuthService interface {
	Login(context.Context, string, string) (string, string, error)
	Register(context.Context, string, string, string) error
	RefreshTokens(context.Context, string) (string, string, error)
	Logout(context.Context, string) error
}

type authService struct {
	client pb.AuthServiceClient
}

func (a *authService) Login(ctx context.Context, username, password string) (accesToken, refreshToken string, err error) {
	req := &pb.LoginRequest{
		Username: username,
		Password: password,
	}

	resp, err := a.client.Login(ctx, req)
	if err != nil {
		return "", "", err
	}

	accesToken = resp.AccessToken
	refreshToken = resp.RefreshToken
	return accesToken, refreshToken, nil
}

func (a *authService) Register(ctx context.Context, username, password, email string) error {
	req := &pb.RegisterRequest{
		Username: username,
		Password: password,
		Email:    email,
	}

	resp, err := a.client.Register(ctx, req)
	if err != nil {
		return err
	}

	if !resp.Success {
		return errors.New("failed to register")
	}

	return nil
}

func (a *authService) RefreshTokens(ctx context.Context, refreshToken string) (string, string, error) {
	req := &pb.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}

	resp, err := a.client.RefreshTokens(ctx, req)
	if err != nil {
		return "", "", err
	}

	return resp.AccessToken, resp.RefreshToken, nil
}

func (a *authService) Logout(ctx context.Context, accessToken string) error {
	req := &pb.LogoutRequest{
		AccessToken: accessToken,
	}

	resp, err := a.client.Logout(ctx, req)
	if err != nil {
		return err
	}

	if !resp.Success {
		return errors.New("failed to logout")
	}

	return nil
}

func NewAuthService(client pb.AuthServiceClient) AuthService {
	return &authService{
		client: client,
	}
}
