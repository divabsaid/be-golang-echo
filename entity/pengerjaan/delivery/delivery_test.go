package delivery_test

import (
	"be-golang-echo/entity/pengerjaan"
	httpDelivery "be-golang-echo/entity/pengerjaan/delivery"
	"be-golang-echo/entity/pengerjaan/mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjo0NzU2ODIzMzEzLCJpZCI6Mn0.Bzz5j-okD6D7obfYomt03kHVmvl4nUB0-ROEEQU1TGA"
)

func TestNewHttpDelivery(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := mock.NewPengerjaanMockUseCase(repo)
	e := echo.New()
	httpDelivery.NewHttpDelivery(e, usecase)

}

func TestAddPengerjaan(t *testing.T) {
	mockPengerjaan := pengerjaanObjSuccess
	jsonPengerjaan, err := json.Marshal(mockPengerjaan)
	assert.NoError(t, err)
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	repo := mock.NewMySQLPengerjaanMockRepository(db)
	usecase := mock.NewPengerjaanMockUseCase(repo)
	// r := redis.InitRedisTest()
	// r.On("Get", "\x02").Return(redisClient.NewStringResult(token, nil))
	// redis.Rdb = r
	e := echo.New()
	req, err := http.NewRequest(echo.POST, "/api/v1/pengerjaans", strings.NewReader(string(jsonPengerjaan)))
	assert.NoError(t, err)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+token)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/api/v1/pengerjaans")
	handler := httpDelivery.PengerjaanHttpDelivery{
		PengerjaanUsecase: usecase,
	}
	err = handler.Add(c)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

}
