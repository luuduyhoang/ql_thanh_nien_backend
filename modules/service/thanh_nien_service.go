package service

import (
	"ql_thanh_nien_backend/modules/model"
	"ql_thanh_nien_backend/modules/repository"
)

type ThanhNienService struct {
	Repo *repository.ThanhNienRepository
}

func (s *ThanhNienService) Create(tn *model.ThanhNien) error {
	return s.Repo.Create(tn)
}

func (s *ThanhNienService) Update(tn *model.ThanhNien, id string) error {
	return s.Repo.Update(tn, id)
}

func (s *ThanhNienService) Delete(id string) error {
	return s.Repo.Delete(id)
}

func (s *ThanhNienService) ListByCapCoSo(maCap string) ([]model.ThanhNien, error) {
	return s.Repo.GetByCapCoSo(maCap)
}

func (s *ThanhNienService) ListByFilter(filter model.ThanhNienFilter) ([]model.ThanhNien, error) {
	return s.Repo.GetByFilter(filter)
}
func (s *ThanhNienService) ExportToExcel(maCap string) ([]model.ThanhNien, error) {
	return s.Repo.ExportToExcel(maCap)
}