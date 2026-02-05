package repository

import (
	"database/sql"
	"fmt"
	"ql_thanh_nien_backend/modules/model"
)

type NguoiDungRepository struct {
	DB *sql.DB
}

func (r *NguoiDungRepository) Create(nd *model.NguoiDung) error {
	query := `
		INSERT INTO nguoi_dung
		(ma_nguoi_dung, ten_dang_nhap, mat_khau, ho_va_ten, vai_tro, ma_cap_co_so)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.DB.Exec(
		query,
		nd.MaNguoiDung,
		nd.TenDangNhap,
		nd.MatKhau,
		nd.TenNguoiDung,
		nd.RoleID,
		nd.MaCapCoSo,
	)
	return err
}

func (r *NguoiDungRepository) Update(nd *model.NguoiDung, id string) error {
	fmt.Printf("Updating nguoi_dung with ID: %s\n", id)
	query := `
		UPDATE nguoi_dung
		SET ten_dang_nhap = ?, mat_khau = ?, ho_va_ten = ?, vai_tro = ?, ma_cap_co_so = ?
		WHERE ma_nguoi_dung = ?
	`
	_, err := r.DB.Exec(
		query,
		nd.TenDangNhap,
		nd.MatKhau,
		nd.TenNguoiDung,
		nd.RoleID,
		nd.MaCapCoSo,
		id,
	)
	return err
}

func (r *NguoiDungRepository) Delete(id string) error {
	query := `
		DELETE FROM nguoi_dung
		WHERE ma_nguoi_dung = ?
	`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *NguoiDungRepository) GetByCapCoSo(maCap string) ([]model.NguoiDung, error) {
	query := `
		SELECT ma_nguoi_dung, ten_dang_nhap, mat_khau, ten_nguoi_dung, vai_tro, ma_cap_co_so
		FROM nguoi_dung
		WHERE ma_cap_co_so = ?
	`
	rows, err := r.DB.Query(query, maCap)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var nguoiDungs []model.NguoiDung
	for rows.Next() {
		var nd model.NguoiDung
		if err := rows.Scan(
			&nd.MaNguoiDung,
			&nd.TenDangNhap,
			&nd.MatKhau,
			&nd.TenNguoiDung,
			&nd.RoleID,
			&nd.MaCapCoSo,
		); err != nil {
			return nil, err
		}
		nguoiDungs = append(nguoiDungs, nd)
	}
	return nguoiDungs, nil
}

func (r *NguoiDungRepository) GetUserByID(id int) (*model.NguoiDung, error) {
    var user model.NguoiDung
    query := `SELECT ma_nguoi_dung, ten_nguoi_dung FROM nguoi_dung WHERE ma_nguoi_dung = ?`
    err := r.DB.QueryRow(query, id).Scan(&user.MaNguoiDung, &user.TenNguoiDung)
    if err != nil {
        return nil, err
    }
    return &user, nil
}