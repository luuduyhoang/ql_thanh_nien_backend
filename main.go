package main

import (
	"ql_thanh_nien_backend/config"
	"ql_thanh_nien_backend/modules/handler"
	"ql_thanh_nien_backend/modules/middleware"
	"ql_thanh_nien_backend/modules/repository"
	"ql_thanh_nien_backend/modules/service"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1️⃣ Kết nối DB
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	config.ConnectDB()

	// 2️⃣ Khởi tạo router
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// 3️⃣ Khởi tạo repository
	userRepo := &repository.UserRepository{DB: config.DB}
	permissionRepo := &repository.PermissionRepository{DB: config.DB}
	thanhNienRepo := &repository.ThanhNienRepository{DB: config.DB}
	nguoiDungRepo := &repository.NguoiDungRepository{DB: config.DB}

	// 4️⃣ Khởi tạo service
	authService := &service.AuthService{UserRepo: userRepo}
	thanhNienService := &service.ThanhNienService{Repo: thanhNienRepo}
	nguoiDungService := &service.NguoiDungService{Repo: nguoiDungRepo}

	// 5️⃣ Khởi tạo handler
	authHandler := &handler.AuthHandler{AuthService: authService}
	thanhNienHandler := &handler.ThanhNienHandler{Service: thanhNienService}
	nguoiDungHandler := &handler.NguoiDungHandler{Service: nguoiDungService}

	// 6️⃣ Public API
	r.POST("/login", authHandler.Login)

	// 7️⃣ Protected API
	api := r.Group("/api")
	api.Use(middleware.JWTMiddleware())
	{
		// Thanh Niên endpoints
		api.GET(
			"/thanh-nien",
			middleware.RequirePermission("THANHNIEN_VIEW", permissionRepo),
			thanhNienHandler.List,
		)

		api.POST(
			"/thanh-nien",
			middleware.RequirePermission("THANHNIEN_CREATE", permissionRepo),
			thanhNienHandler.Create,
		)
		api.PUT(
			"/thanh-nien/:id",
			middleware.RequirePermission("THANHNIEN_UPDATE", permissionRepo),
			func(c *gin.Context) {
				id := c.Param("id")
				thanhNienHandler.Update(c, id)
			},
		)
		api.DELETE(
			"/thanh-nien/:id",
			middleware.RequirePermission("THANHNIEN_DELETE", permissionRepo),
			func(c *gin.Context) {
				id := c.Param("id")
				thanhNienHandler.Delete(c, id)
			},
		)
		api.GET(
			"/thanh-nien/export",
			middleware.RequirePermission("THANHNIEN_EXPORT", permissionRepo),
			// thanhNienHandler.ExportToExcel,
		)

		// User management endpoints
		api.GET("/nguoi-dung",
			middleware.RequirePermission("USER_VIEW", permissionRepo),
			nguoiDungHandler.List,
		)

		api.GET("/me", nguoiDungHandler.GetMe)

		api.POST("/nguoi-dung",
			middleware.RequirePermission("USER_CREATE", permissionRepo),
			nguoiDungHandler.Create,
		)

		api.PUT("/nguoi-dung/:id",
			middleware.RequirePermission("USER_UPDATE", permissionRepo),
			func(c *gin.Context) {
				id := c.Param("id")
				nguoiDungHandler.Update(c, id)
			},
		)

		api.DELETE("/nguoi-dung/:id",
			middleware.RequirePermission("USER_DELETE", permissionRepo),
			func(c *gin.Context) {
				id := c.Param("id")
				nguoiDungHandler.Delete(c, id)
			},
		)
	}

	// 8️⃣ Run server
	r.Run(":" + port)
}
