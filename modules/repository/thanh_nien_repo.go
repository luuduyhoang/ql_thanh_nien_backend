package repository

import (
	"database/sql"
	"fmt"
	"ql_thanh_nien_backend/modules/model"

	"github.com/google/uuid"
)

type ThanhNienRepository struct {
	DB *sql.DB
}

func (r *ThanhNienRepository) Create(tn *model.ThanhNien) error {
	id := uuid.New() // UUID v4
	query := `
		INSERT INTO thanh_nien
		(ma_thanh_nien, ho_va_ten, ngay_sinh, dang_vien, doan_vien, van_hoa, trinh_do_cmkt, dan_toc, ton_giao, que_quan, phan_loai, ho_ten_cha, ho_ten_me, thoi_gian_nhap_ngu, don_vi_nhap_ngu, hoan_thanh_nvqs, ma_cap_co_so, ma_don_vi)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.DB.Exec(
		query,
		id,
		tn.HoVaTen,
		tn.NgaySinh,
		tn.DangVien,
		tn.DoanVien,
		tn.VanHoa,
		tn.TrinhDoCmkt,
		tn.DanToc,
		tn.TonGiao,
		tn.QueQuan,
		tn.PhanLoai,
		tn.HoTenCha,
		tn.HoTenMe,
		tn.ThoiGianNhapNgu,
		tn.DonViNhapNgu,
		tn.HoanThanhNVQS,
		tn.MaCapCoSo,
		tn.MaDonVi,
	)
	return err
}
func (r *ThanhNienRepository) Update(tn *model.ThanhNien, id string) error {
	fmt.Printf("Updating thanh_nien with ID: %s\n", id)
	query := `
		UPDATE thanh_nien
		SET ho_va_ten = ?, ngay_sinh = ?, dang_vien = ?, doan_vien = ?, van_hoa = ?, trinh_do_cmkt = ?, dan_toc = ?, ton_giao = ?, que_quan = ?, phan_loai = ?, ho_ten_cha = ?, ho_ten_me = ?, thoi_gian_nhap_ngu = ?, don_vi_nhap_ngu = ?, hoan_thanh_nvqs = ?, ma_cap_co_so = ?, ma_don_vi = ?
		WHERE ma_thanh_nien = ?
	`
	_, err := r.DB.Exec(
		query,
		tn.HoVaTen,
		tn.NgaySinh,
		tn.DangVien,
		tn.DoanVien,
		tn.VanHoa,
		tn.TrinhDoCmkt,
		tn.DanToc,
		tn.TonGiao,
		tn.QueQuan,
		tn.PhanLoai,
		tn.HoTenCha,
		tn.HoTenMe,
		tn.ThoiGianNhapNgu,
		tn.DonViNhapNgu,
		tn.HoanThanhNVQS,
		tn.MaCapCoSo,
		tn.MaDonVi,
		id,
	)
	return err
}

func (r *ThanhNienRepository) Delete(id string) error {
	query := `
		DELETE FROM thanh_nien
		WHERE ma_thanh_nien = ?
	`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *ThanhNienRepository) GetByCapCoSo(maCap string) ([]model.ThanhNien, error) {
	query := `
		SELECT ma_thanh_nien, ho_va_ten, ngay_sinh, dang_vien, doan_vien,
		       van_hoa, trinh_do_cmkt, dan_toc, ton_giao, que_quan,
		       phan_loai, ho_ten_cha, ho_ten_me, thoi_gian_nhap_ngu, don_vi_nhap_ngu, hoan_thanh_nvqs, ma_cap_co_so, ma_don_vi
		FROM thanh_nien
		WHERE ma_cap_co_so = ?
	`

	rows, err := r.DB.Query(query, maCap)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.ThanhNien

	for rows.Next() {
		var tn model.ThanhNien
		var ngaySinh sql.NullString
		var thoiGianNhapNgu sql.NullString

		err := rows.Scan(
			&tn.MaThanhNien,
			&tn.HoVaTen,
			&ngaySinh,
			&tn.DangVien,
			&tn.DoanVien,
			&tn.VanHoa,
			&tn.TrinhDoCmkt,
			&tn.DanToc,
			&tn.TonGiao,
			&tn.QueQuan,
			&tn.PhanLoai,
			&tn.HoTenCha,
			&tn.HoTenMe,
			&thoiGianNhapNgu,
			&tn.DonViNhapNgu,
			&tn.HoanThanhNVQS,
			&tn.MaCapCoSo,
			&tn.MaDonVi,
		)
		if err != nil {
			return nil, err
		}

		// GÁN STRING (KHÔNG DÙNG time.Time NỮA)
		if ngaySinh.Valid {
			tn.NgaySinh = ngaySinh.String
		} else {
			tn.NgaySinh = "" // string rỗng
		}
		if thoiGianNhapNgu.Valid {
			tn.ThoiGianNhapNgu = thoiGianNhapNgu.String
		} else {
			tn.ThoiGianNhapNgu = "" // string rỗng
		}

		list = append(list, tn)
	}

	return list, nil
}


func (r *ThanhNienRepository) GetByFilter(filter model.ThanhNienFilter) ([]model.ThanhNien, error) {
	query := `
		SELECT ma_thanh_nien, ho_va_ten, ngay_sinh, dang_vien, doan_vien, van_hoa, trinh_do_cmkt, dan_toc, ton_giao, que_quan, phan_loai, ho_ten_cha, ho_ten_me, ma_cap_co_so, ma_don_vi
		FROM thanh_nien
		WHERE 1=1
	`
	var args []interface{}

	if filter.HoVaTen != nil {
		query += " AND ho_va_ten LIKE ?"
		args = append(args, "%"+*filter.HoVaTen+"%")
	}
	if filter.DangVien != nil {
		query += " AND dang_vien = ?"
		args = append(args, *filter.DangVien)
	}
	if filter.DoanVien != nil {
		query += " AND doan_vien = ?"
		args = append(args, *filter.DoanVien)
	}
	if filter.MaCapCoSo != nil {
		query += " AND ma_cap_co_so = ?"
		args = append(args, *filter.MaCapCoSo)
	}
	if filter.MaDonVi != nil {
		query += " AND ma_don_vi = ?"
		args = append(args, *filter.MaDonVi)
	}

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.ThanhNien
	for rows.Next() {
		var tn model.ThanhNien
		err := rows.Scan(
			&tn.MaThanhNien,
			&tn.HoVaTen,
			&tn.NgaySinh,
			&tn.DangVien,
			&tn.DoanVien,
			&tn.MaCapCoSo,
			&tn.MaDonVi,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, tn)
	}
	return list, nil
}

func (r *ThanhNienRepository) ExportToExcel(maCap string) ([]model.ThanhNien, error) {
	return r.ExportToExcel(maCap)
}