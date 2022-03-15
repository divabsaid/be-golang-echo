package mock

import (
	"be-golang-echo/entity/pengerjaan"
	"be-golang-echo/entity/pengerjaan/repository"
)

type PengerjaanMockUsecase interface {
	Add(rm *pengerjaan.PengerjaanModel) (bool, error)
	Update(id int, rm *pengerjaan.PengerjaanModel) (bool, error)
	Delete(id int) (bool, error)
	GetList(rf *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error)
	GetSingle(id int) (*pengerjaan.PengerjaanModel, error)
}

type pengerjaanUsecase struct {
	pengerjaanRepository repository.PengerjaanRepository
}

func NewPengerjaanMockUseCase(u repository.PengerjaanRepository) PengerjaanMockUsecase {
	return &pengerjaanUsecase{
		pengerjaanRepository: u,
	}
}

func (r *pengerjaanUsecase) Add(rm *pengerjaan.PengerjaanModel) (bool, error) {
	if rm.ID != 1 {
		return false, pengerjaan.ADD_FAILED
	}
	return true, nil
}

func (r *pengerjaanUsecase) Update(id int, rm *pengerjaan.PengerjaanModel) (bool, error) {
	if id != 1 {
		return false, pengerjaan.UPDATE_FAILED
	}
	return true, nil
}

func (r *pengerjaanUsecase) Delete(id int) (bool, error) {
	if id != 1 {
		return false, pengerjaan.DELETE_FAILED
	}
	return true, nil
}

func (r *pengerjaanUsecase) GetList(rf *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error) {
	pengerjaanObj := make([]*pengerjaan.PengerjaanModel, 0)
	return pengerjaanObj, nil
}

func (r *pengerjaanUsecase) GetSingle(id int) (*pengerjaan.PengerjaanModel, error) {
	pengerjaanObj := new(pengerjaan.PengerjaanModel)
	return pengerjaanObj, nil
}
