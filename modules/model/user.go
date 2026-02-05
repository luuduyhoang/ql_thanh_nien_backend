package model

type NguoiDung struct {
	MaNguoiDung   	string `gorm:"primaryKey" json:"ma_nguoi_dung"`
	TenDangNhap 	string `gorm:"column:ten_dang_nhap;unique;not null" json:"ten_dang_nhap"`
	MatKhau 		string `gorm:"column:mat_khau;not null" json:"mat_khau"`
	TenNguoiDung 	string `gorm:"column:ten_nguoi_dung;not null" json:"ten_nguoi_dung"`
	RoleID   		string `gorm:"column:role_id;not null" json:"role_id"`
	MaCapCoSo   	*string `gorm:"column:ma_cap_co_so;not null" json:"ma_cap_co_so"`
}

type NguoiDungLogin struct {
	TenDangNhap 	string `json:"ten_dang_nhap" binding:"required"`
	MatKhau 		string `json:"mat_khau" binding:"required"`
}
