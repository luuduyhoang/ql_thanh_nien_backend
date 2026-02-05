package utils

import (
	"time"

	"ql_thanh_nien_backend/modules/model"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("SECRET_KEY")

func GenerateToken(user model.NguoiDung) (string, error) {
	claims := jwt.MapClaims{
		"ma_nguoi_dung": user.MaNguoiDung,
		"role_id":       user.RoleID,
		"ma_cap_co_so":  user.MaCapCoSo,
		"exp":           time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
func ParseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}