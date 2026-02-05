package repository

import "database/sql"

type PermissionRepository struct {
	DB *sql.DB
}

func (r *PermissionRepository) HasPermission(roleID, permissionCode string) (bool, error) {
	query := `
		SELECT COUNT(*)
		FROM role_permissions rp
		JOIN permissions p ON rp.permission_id = p.permission_id
		WHERE rp.role_id = ? AND p.permission_code = ?
	`

	var count int
	err := r.DB.QueryRow(query, roleID, permissionCode).Scan(&count)
	return count > 0, err
}
