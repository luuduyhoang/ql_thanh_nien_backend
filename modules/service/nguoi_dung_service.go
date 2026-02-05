package service

import (
	"ql_thanh_nien_backend/modules/model"
	"ql_thanh_nien_backend/modules/repository"
)

type NguoiDungService struct {
	Repo *repository.NguoiDungRepository
}

func (s *NguoiDungService) Create(nd *model.NguoiDung) error {
	return s.Repo.Create(nd)
}

func (s *NguoiDungService) Update(nd *model.NguoiDung, id string) error {
	return s.Repo.Update(nd, id)
}

func (s *NguoiDungService) Delete(id string) error {
	return s.Repo.Delete(id)
}

func (s *NguoiDungService) ListByCapCoSo(maCap string) ([]model.NguoiDung, error) {
	return s.Repo.GetByCapCoSo(maCap)
}

func (s *NguoiDungService) GetUserByID(id int) (*model.NguoiDung, error) {
	return s.Repo.GetUserByID(id)
}