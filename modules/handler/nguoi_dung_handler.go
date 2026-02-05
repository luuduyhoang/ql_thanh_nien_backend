package handler

import (
	"net/http"

	"ql_thanh_nien_backend/modules/model"
	"ql_thanh_nien_backend/modules/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type NguoiDungHandler struct {
	Service *service.NguoiDungService
}
func (h *NguoiDungHandler) Create(c *gin.Context) {
	var nd model.NguoiDung
	if err := c.ShouldBindJSON(&nd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lấy cấp cơ sở từ JWT
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)
	nd.MaCapCoSo = &maCap

	if err := h.Service.Create(&nd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func (h *NguoiDungHandler) Update(c *gin.Context, id string) {
	var nd model.NguoiDung
	if err := c.ShouldBindJSON(&nd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Lấy cấp cơ sở từ JWT
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)
	nd.MaCapCoSo = &maCap

	if err := h.Service.Update(&nd, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}	
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *NguoiDungHandler) Delete(c *gin.Context, id string) {
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
func (h *NguoiDungHandler) ListByCapCoSo(c *gin.Context) {
	// Lấy cấp cơ sở từ JWT
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)

	ndList, err := h.Service.ListByCapCoSo(maCap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ndList)
}

func (h *NguoiDungHandler) List(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)

	data, err := h.Service.ListByCapCoSo(maCap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *NguoiDungHandler) GetMe(c *gin.Context) {
    // Lấy claims từ JWT (được set trong middleware)
    claims := c.MustGet("claims").(jwt.MapClaims)
    userID := int(claims["ma_nguoi_dung"].(float64)) // JWT lưu số kiểu float64

    user, err := h.Service.GetUserByID(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    // Trả về chỉ những thông tin cần thiết
    c.JSON(http.StatusOK, gin.H{
        "id":             user.MaNguoiDung,
        "ten_nguoi_dung": user.TenNguoiDung,
    })
}
