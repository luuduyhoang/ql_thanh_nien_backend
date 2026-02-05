// modules/repository/user_repo.go
package repository

import (
	"database/sql"
	"ql_thanh_nien_backend/modules/model"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) FindByUsername(ten_dang_nhap string) (*model.NguoiDung, error) {
	query := `SELECT ma_nguoi_dung, ten_dang_nhap, ten_nguoi_dung, mat_khau, ma_cap_co_so, role_id FROM nguoi_dung WHERE ten_dang_nhap=?`
	row := r.DB.QueryRow(query, ten_dang_nhap)

	var u model.NguoiDung
	err := row.Scan(&u.MaNguoiDung, &u.TenDangNhap, &u.TenNguoiDung, &u.MatKhau, &u.MaCapCoSo, &u.RoleID)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
