package usecase

import (
	"be-golang-echo/entity/pengerjaan"
	"be-golang-echo/entity/pengerjaan/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type PengerjaanUsecase interface {
	GetList(filter *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error)
	Add(pm *pengerjaan.PengerjaanModel) (bool, error)
	Update(id int, pm *pengerjaan.PengerjaanModel) (bool, error)
	Delete(id int) (bool, error)
	GetSingle(id int) (*pengerjaan.PengerjaanModel, error)
}

type pengerjaanUsecase struct {
	pengerjaanRepository repository.PengerjaanRepository
}

func NewPengerjaanUseCase(u repository.PengerjaanRepository) PengerjaanUsecase {
	return &pengerjaanUsecase{
		pengerjaanRepository: u,
	}
}

func (p *pengerjaanUsecase) GetList(filter *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error) {
	res, err := p.pengerjaanRepository.GetList(filter)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (p *pengerjaanUsecase) Add(pm *pengerjaan.PengerjaanModel) (bool, error) {
	err := validator.New().Struct(pm)
	if err != nil {
		return false, pengerjaan.REQUEST_BODY_NOT_VALID
	}

	pm.CreatedAt = time.Now()
	res, err := p.pengerjaanRepository.Add(pm)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (p *pengerjaanUsecase) Update(id int, pm *pengerjaan.PengerjaanModel) (bool, error) {
	pm.UpdatedAt = time.Now()
	err := validator.New().Struct(pm)
	if err != nil {
		return false, pengerjaan.REQUEST_BODY_NOT_VALID
	}
	res, err := p.pengerjaanRepository.Update(id, pm)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (p *pengerjaanUsecase) Delete(id int) (bool, error) {
	res, err := p.pengerjaanRepository.Delete(id)
	if err != nil {
		return false, err
	}
	return res, nil
}

func (p *pengerjaanUsecase) GetSingle(id int) (*pengerjaan.PengerjaanModel, error) {
	res, err := p.pengerjaanRepository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
