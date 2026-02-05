package model

type ThanhNien struct {
	MaThanhNien string     `json:"ma_thanh_nien"`
	HoVaTen     string     `json:"ho_va_ten"`
	NgaySinh    string      `gorm:"type:date;default:'1990-01-01';column:ngay_sinh" json:"ngay_sinh"`
	DangVien    bool       `json:"dang_vien"`
	DoanVien    bool       `json:"doan_vien"`
	VanHoa	 	*string    `json:"van_hoa"`
	TrinhDoCmkt *string    `json:"trinh_do_cmkt"`
	DanToc	 	*string    `json:"dan_toc"`
	TonGiao	 	*string    `json:"ton_giao"`
	QueQuan     *string    `json:"que_quan"`
	PhanLoai   	*string    `json:"phan_loai"`
	HoTenCha   	*string    `json:"ho_ten_cha"`
	HoTenMe		*string    `json:"ho_ten_me"`
	ThoiGianNhapNgu string   `json:"thoi_gian_nhap_ngu"`
	DonViNhapNgu   *string    `json:"don_vi_nhap_ngu"`
	HoanThanhNVQS bool       `json:"hoan_thanh_nvqs"`
	MaCapCoSo   *string    `gorm:"column:ma_cap_co_so;default:'1'" json:"ma_cap_co_so"`
	MaDonVi     *string    `gorm:"column:ma_don_vi;default:'1'" json:"ma_don_vi"`
}
type ThanhNienFilter struct {
	HoVaTen   *string
	DangVien  *bool
	DoanVien  *bool
	MaCapCoSo *string
	MaDonVi   *string
}