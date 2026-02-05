package middleware

import (
	"log"
	"net/http"

	"ql_thanh_nien_backend/modules/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequirePermission(permissionCode string, repo *repository.PermissionRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		claimsAny, exists := c.Get("claims")
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims := claimsAny.(jwt.MapClaims)
		roleID := claims["role_id"].(string)

		log.Printf("Checking permission '%s' for role ID '%s'\n", permissionCode, roleID)

		ok, err := repo.HasPermission(roleID, permissionCode)
		if err != nil || !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{	
				"error": "permission denied",
			})
			return
		}

		c.Next()
	}
}
