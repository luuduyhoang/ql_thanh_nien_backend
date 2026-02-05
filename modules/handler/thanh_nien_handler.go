package handler

import (
	"bytes"
	"net/http"

	"ql_thanh_nien_backend/modules/model"
	"ql_thanh_nien_backend/modules/service"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ThanhNienHandler struct {
	Service *service.ThanhNienService
}

func (h *ThanhNienHandler) Create(c *gin.Context) {
	var tn model.ThanhNien
	if err := c.ShouldBindJSON(&tn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lấy cấp cơ sở từ JWT
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)
	tn.MaCapCoSo = &maCap

	if err := h.Service.Create(&tn); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

func (h *ThanhNienHandler) Update(c *gin.Context, id string) {
	var tn model.ThanhNien
	if err := c.ShouldBindJSON(&tn); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Lấy cấp cơ sở từ JWT
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)
	tn.MaCapCoSo = &maCap

	if err := h.Service.Update(&tn, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (h *ThanhNienHandler) Delete(c *gin.Context, id string) {
	if err := h.Service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (h *ThanhNienHandler) List(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)

	data, err := h.Service.ListByCapCoSo(maCap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
func (h *ThanhNienHandler) ListByFilter(c *gin.Context) {
	var filter model.ThanhNienFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.Service.ListByFilter(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}
func (h *ThanhNienHandler) ExportToExcel(c *gin.Context) (*bytes.Buffer, error) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	maCap := claims["ma_cap_co_so"].(string)

	data, err := h.Service.ExportToExcel(maCap)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return nil, err
	}

	// Trả về dữ liệu dưới dạng buffer
	var buffer bytes.Buffer
	// Giả sử bạn có một hàm để ghi dữ liệu vào buffer dưới dạng Excel
	// writeDataToExcelBuffer(&buffer, data)
	log.Printf("Exported %d records to Excel", len(data)) // Thêm log để kiểm tra số bản ghi

	return &buffer, nil
}