// modules/service/auth_service.go
package service

import (
	"errors"
	"log"
	"ql_thanh_nien_backend/modules/repository"
	"ql_thanh_nien_backend/modules/utils"
)

type AuthService struct {
	UserRepo *repository.UserRepository
}

func (s *AuthService) Login(ten_dang_nhap, mat_khau string) (string, error) {
	// log.Printf("Login attempt for user: %s", ten_dang_nhap)
	user, err := s.UserRepo.FindByUsername(ten_dang_nhap)
	if err != nil {
		return "", errors.New("user not found")
	}

	// log.Printf("User found: %+v", user)
	log.Printf("Comparing password hash: %s with password: %s", user.MatKhau, mat_khau)

	if utils.CheckPassword(user.MatKhau, mat_khau) {
		return "", errors.New("invalid password")
	}

	return utils.GenerateToken(
		*user,
	)
}

