// modules/handler/auth_handler.go
package handler

import (
	"log"
	"net/http"
	"ql_thanh_nien_backend/modules/model"
	"ql_thanh_nien_backend/modules/service"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input model.NguoiDungLogin

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("input", input)

	token, err := h.AuthService.Login(input.TenDangNhap, input.MatKhau)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
