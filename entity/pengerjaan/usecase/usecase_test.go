package usecase_test

import (
	"be-golang-echo/entity/pengerjaan"
	"be-golang-echo/entity/pengerjaan/mock"
	"be-golang-echo/entity/pengerjaan/usecase"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v2"
)

var (
	pengerjaanObjSuccess = &pengerjaan.PengerjaanModel{
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

	pengerjaanObjFail = &pengerjaan.PengerjaanModel{
		ID:              2,
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

	pengerjaanObjWrong = &pengerjaan.PengerjaanModel{
		ID:              2,
		Status:          "",
		Tanggal:         "",
		NamaSPK:         "",
		Pengaduan:       "",
		Lokasi:          "",
		Detail:          "",
		PenanggungJawab: "",
		EstimasiSelesai: "",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	filter = &pengerjaan.Filter{}

	wrongFilter = &pengerjaan.Filter{}
)

func TestAdd(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.Add(pengerjaanObjSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestAddError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.Add(pengerjaanObjFail)
	assert.Error(t, err)
	assert.False(t, res)

}

func TestAddBadRequest(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.Add(pengerjaanObjWrong)
	assert.Error(t, err)
	assert.False(t, res)

}

func TestUpdate(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.Update(pengerjaanObjSuccess.ID, pengerjaanObjSuccess)
	assert.NoError(t, err)
	assert.NotNil(t, res)
}

func TestUpdateError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.Update(pengerjaanObjFail.ID, pengerjaanObjFail)
	assert.Error(t, err)
	assert.False(t, res)
}

func TestUpdateBadRequest(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.Update(pengerjaanObjWrong.ID, pengerjaanObjWrong)
	assert.Error(t, err)
	assert.False(t, res)
}

func TestDelete(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.Delete(pengerjaanObjSuccess.ID)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestDeleteError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	_, err = usecase.Delete(pengerjaanObjFail.ID)
	assert.Error(t, err)
}

func TestGetList(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	res, err := usecase.GetList(filter)
	assert.NoError(t, err)
	assert.NotNil(t, res)

}

func TestGetListError(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := usecase.NewPengerjaanUseCase(repo)
	_, err = usecase.GetList(wrongFilter)
	assert.Error(t, err)

}
