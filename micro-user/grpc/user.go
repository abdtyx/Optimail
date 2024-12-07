package grpc

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

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

func (s *Server) getPkById(id string) (int64, error) {
	var user model.User
	tx := s.db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return -1, tx.Error
	}
	return user.PK, nil
}

func (s *Server) UsersCreateUser(ctx context.Context, in *dto.UsersCreateUserRequest) (*dto.UsersCreateUserResponse, error) {
	var user model.User
	tx := s.db.Where("username = ?", in.Username).First(&user)
	if tx.Error != nil && tx.Error != gorm.ErrRecordNotFound {
		return nil, tx.Error
	}
	if tx.RowsAffected != 0 {
		return nil, errors.New("user exists")
	}

	user.ID = uuid.NewString()
	user.Username = in.Username
	user.Password = hex.EncodeToString(argon2Key([]byte(in.Password), []byte(user.ID), 256))

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// create default settings for user
	var setting = model.Setting{
		UserPK:         user.PK,
		Email:          "",
		SummaryLength:  50,
		SummaryPrompt:  "Next I will show you an email. Your job is to analyse it and output a summary of no more than 50 words.",
		EmphasisPrompt: "Next I will show you an email. Your job is to mark important words in bold in the email so as to make it easier for people to read.",
	}

	tx = s.db.Create(&setting)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &dto.UsersCreateUserResponse{
		Pk:      user.PK,
		Message: "OK",
	}, nil
}

func (s *Server) UsersLogin(ctx context.Context, in *dto.UsersLoginRequest) (*dto.UsersLoginResponse, error) {
	var user model.User
	if err := s.db.Where("username = ?", in.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid username")
		} else {
			return nil, err
		}
	}

	if h := hex.EncodeToString(argon2Key([]byte(in.Password), []byte(user.ID), 256)); h != user.Password {
		return nil, errors.New("invalid password")
	}

	return &dto.UsersLoginResponse{
		Id:      user.ID,
		Message: "OK",
	}, nil
}

func (s *Server) UsersLogout(ctx context.Context, in *dto.UsersLogoutRequest) (*dto.UsersLogoutResponse, error) {
	return &dto.UsersLogoutResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersChangePwd(ctx context.Context, in *dto.UsersChangePwdRequest) (*dto.UsersChangePwdResponse, error) {
	var user model.User
	if err := s.db.Where("id = ?", in.Id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("invalid id")
		} else {
			return nil, err
		}
	}

	if h := hex.EncodeToString(argon2Key([]byte(in.OldPassword), []byte(user.ID), 256)); h != user.Password {
		return nil, errors.New("invalid password")
	}

	user.Password = hex.EncodeToString(argon2Key([]byte(in.NewPassword), []byte(user.ID), 256))

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
	pk, err := s.getPkById(in.Id)
	if err != nil {
		return nil, err
	}
	var setting model.Setting
	tx := s.db.Where("user_pk = ?", pk).First(&setting)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dto.UsersGetSettingsResponse{
		Settings: &dto.Settings{
			Email:          setting.Email,
			SummaryLength:  int32(setting.SummaryLength),
			SummaryPrompt:  setting.SummaryPrompt,
			EmphasisPrompt: setting.EmphasisPrompt,
		},
		Message: "OK",
	}, nil
}

func (s *Server) UsersUpdateSettings(ctx context.Context, in *dto.UsersUpdateSettingsRequest) (*dto.UsersUpdateSettingsResponse, error) {
	pk, err := s.getPkById(in.Id)
	if err != nil {
		return nil, err
	}
	var setting model.Setting
	tx := s.db.Where("user_pk = ?", pk).First(&setting)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if in.Settings.SummaryPrompt == "" {
		in.Settings.SummaryPrompt = fmt.Sprintf("Next I will show you an email. Your job is to analyse it and output a summary of no more than %v words.", in.Settings.SummaryLength)
	}
	if in.Settings.EmphasisPrompt == "" {
		in.Settings.EmphasisPrompt = "Next I will show you an email. Your job is to mark important words in bold (using html syntax) in the email so as to make it easier for people to read."
	}
	setting.Email = in.Settings.Email
	setting.SummaryLength = int64(in.Settings.SummaryLength)
	setting.SummaryPrompt = in.Settings.SummaryPrompt
	setting.EmphasisPrompt = in.Settings.EmphasisPrompt
	s.db.Model(&setting).Where("user_pk = ?", pk).Updates(setting)
	return &dto.UsersUpdateSettingsResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersGetIdByEmail(ctx context.Context, in *dto.UsersGetIdByEmailRequest) (*dto.UsersGetIdByEmailResponse, error) {
	var setting model.Setting
	tx := s.db.Where("email = ?", in.Email).First(&setting)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var user = model.User{PK: setting.UserPK}
	tx = s.db.First(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dto.UsersGetIdByEmailResponse{
		Id: user.ID,
	}, nil
}

func (s *Server) UsersCreateSummary(ctx context.Context, in *dto.UsersCreateSummaryRequest) (*dto.UsersCreateSummaryResponse, error) {
	pk, err := s.getPkById(in.Id)
	if err != nil {
		return nil, err
	}
	var summary = model.Summary{
		UserPK:  pk,
		Content: in.Summary,
	}
	tx := s.db.Create(&summary)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dto.UsersCreateSummaryResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersGetSummary(ctx context.Context, in *dto.UsersGetSummaryRequest) (*dto.UsersGetSummaryResponse, error) {
	pk, err := s.getPkById(in.Id)
	if err != nil {
		return nil, err
	}
	var res []model.Summary
	tx := s.db.Where("user_pk = ?", pk).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var contents []string
	for _, r := range res {
		contents = append(contents, r.Content)
	}
	return &dto.UsersGetSummaryResponse{
		Summary: contents,
	}, nil
}

func (s *Server) UsersCreateEmphasis(ctx context.Context, in *dto.UsersCreateEmphasisRequest) (*dto.UsersCreateEmphasisResponse, error) {
	pk, err := s.getPkById(in.Id)
	if err != nil {
		return nil, err
	}
	var emphasis = model.Emphasis{
		UserPK:  pk,
		Content: in.Emphasis,
	}
	tx := s.db.Create(&emphasis)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dto.UsersCreateEmphasisResponse{
		Message: "OK",
	}, nil
}

func (s *Server) UsersGetEmphasis(ctx context.Context, in *dto.UsersGetEmphasisRequest) (*dto.UsersGetEmphasisResponse, error) {
	pk, err := s.getPkById(in.Id)
	if err != nil {
		return nil, err
	}
	var res []model.Emphasis
	tx := s.db.Where("user_pk = ?", pk).Find(&res)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var contents []string
	for _, r := range res {
		contents = append(contents, r.Content)
	}
	return &dto.UsersGetEmphasisResponse{
		Emphasis: contents,
	}, nil
}
