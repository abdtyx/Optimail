package grpc

import (
	"context"
	"encoding/hex"
	"errors"

	"github.com/abdtyx/Optimail/micro-user/dto"
	"github.com/abdtyx/Optimail/micro-user/model"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
	"gorm.io/gorm"
)

func argon2Key(password []byte, salt []byte, keyLen uint32) []byte {
	result := argon2.IDKey(password, salt, 1, 64*1024, 4, keyLen)
	return result
}

func (s *Server) UsersCreateUser(ctx context.Context, in *dto.UsersCreateUserRequest) (*dto.UsersCreateUserResponse, error) {
	var user = model.User{Username: in.Username}
	tx := s.db.First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, tx.Error
	}
	if tx.RowsAffected != 0 {
		return nil, errors.New("user exists")
	}

	user.ID = uuid.NewString()
	user.Password = hex.EncodeToString(argon2Key([]byte(in.Password), []byte(user.ID), 255))

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &dto.UsersCreateUserResponse{
		Pk:      user.PK,
		Message: "OK",
	}, nil
}

func (s *Server) UsersLogin(ctx context.Context, in *dto.UsersLoginRequest) (*dto.UsersLoginResponse, error) {
	var user = model.User{Username: in.Username}
	if err := s.db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid username")
		} else {
			return nil, err
		}
	}

	if h := hex.EncodeToString(argon2Key([]byte(in.Password), []byte(user.ID), 255)); h != user.Password {
		return nil, errors.New("invalid password")
	}

	return &dto.UsersLoginResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersLogout(ctx context.Context, in *dto.UsersLogoutRequest) (*dto.UsersLogoutResponse, error) {
	return &dto.UsersLogoutResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersChangePwd(ctx context.Context, in *dto.UsersChangePwdRequest) (*dto.UsersChangePwdResponse, error) {
	var user = model.User{ID: in.Id}
	if err := s.db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid id")
		} else {
			return nil, err
		}
	}

	if h := hex.EncodeToString(argon2Key([]byte(in.OldPassword), []byte(user.ID), 255)); h != user.Password {
		return nil, errors.New("invalid password")
	}

	user.Password = hex.EncodeToString(argon2Key([]byte(in.NewPassword), []byte(user.ID), 255))

	tx := s.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected != 1 {
		return nil, errors.New("failed to change password")
	}
	return &dto.UsersChangePwdResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersGetSettings(ctx context.Context, in *dto.UsersGetSettingsRequest) (*dto.UsersGetSettingsResponse, error) {
	return &dto.UsersGetSettingsResponse{
		Settings: nil,
		Message:  "OK",
	}, nil
}

func (s *Server) UsersUpdateSettings(ctx context.Context, in *dto.UsersUpdateSettingsRequest) (*dto.UsersUpdateSettingsResponse, error) {
	return &dto.UsersUpdateSettingsResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersGetIdByEmail(ctx context.Context, in *dto.UsersGetIdByEmailRequest) (*dto.UsersGetIdByEmailResponse, error) {

	return nil, nil
}

func (s *Server) UsersCreateSummary(ctx context.Context, in *dto.UsersCreateSummaryRequest) (*dto.UsersCreateSummaryResponse, error) {
	return nil, nil
}
