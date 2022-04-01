package repository_test

import (
	"be-golang-echo/entity/pengerjaan"
	"be-golang-echo/entity/pengerjaan/repository"
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
)

func TestAdd(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// query := "INSERT role SET role=\\?, created_at=\\?"

	query := "INSERT pengerjaan SET status=\\?, tanggal=\\?, nama_spk=\\?, pengaduan=\\?, lokasi=\\?, detail=\\?, penanggung_jawab=\\?, estimasi_selesai=\\?, created_at=\\?"

	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(pengerjaanObjSuccess.Status, pengerjaanObjSuccess.Tanggal, pengerjaanObjSuccess.NamaSPK, pengerjaanObjSuccess.Pengaduan, pengerjaanObjSuccess.Lokasi, pengerjaanObjSuccess.Detail, pengerjaanObjSuccess.PenanggungJawab, pengerjaanObjSuccess.EstimasiSelesai, pengerjaanObjSuccess.CreatedAt).WillReturnResult(sqlmock.NewResult(1, 1))

	repo := repository.NewMySQLPengerjaanRepository(db)

	res, err := repo.Add(pengerjaanObjSuccess)
	assert.NoError(t, err)
	assert.True(t, res)
}
