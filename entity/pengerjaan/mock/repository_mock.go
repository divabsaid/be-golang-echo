package mock

import (
	"be-golang-echo/entity/pengerjaan"
	"database/sql"
	"time"
)

type PengerjaanMockRepository interface {
	GetList(filter *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error)
	Add(pm *pengerjaan.PengerjaanModel) (bool, error)
	Update(id int, pm *pengerjaan.PengerjaanModel) (bool, error)
	Delete(id int) (bool, error)
	GetByID(id int) (*pengerjaan.PengerjaanModel, error)
}

type mysqlPengerjaanRepository struct {
	db *sql.DB
}

var pengerjaanObj = &pengerjaan.PengerjaanModel{
	ID:              1,
	Status:          "Belum Ditugaskan",
	Tanggal:         "2022-03-11 00:00:00.000",
	NamaSPK:         "HBL 123",
	Pengaduan:       "Kerusakan Aset",
	Lokasi:          "Blimbing",
	Detail:          "Tekanan Pipa 0,1 bar",
	PenanggungJawab: "Belum ada",
	EstimasiSelesai: "-",
	CreatedAt:       time.Now(),
	UpdatedAt:       time.Now(),
}

func NewMySQLPengerjaanMockRepository(db *sql.DB) PengerjaanMockRepository {
	return &mysqlPengerjaanRepository{db}
}

func (m *mysqlPengerjaanRepository) Add(r *pengerjaan.PengerjaanModel) (bool, error) {
	if r.ID != 1 {
		return false, pengerjaan.ADD_FAILED
	}
	return true, nil

}

func (m *mysqlPengerjaanRepository) Update(id int, r *pengerjaan.PengerjaanModel) (bool, error) {
	if id != 1 {
		return false, pengerjaan.UPDATE_FAILED
	}
	return true, nil

}

func (m *mysqlPengerjaanRepository) Delete(id int) (bool, error) {
	if id != 1 {
		return false, pengerjaan.DELETE_FAILED
	}
	return true, nil
}

func (m *mysqlPengerjaanRepository) GetList(r *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error) {
	pengerjaanObj := make([]*pengerjaan.PengerjaanModel, 0)
	return pengerjaanObj, nil
}

func (m *mysqlPengerjaanRepository) GetByID(id int) (*pengerjaan.PengerjaanModel, error) {
	pengerjaanObj := new(pengerjaan.PengerjaanModel)
	return pengerjaanObj, nil
}
