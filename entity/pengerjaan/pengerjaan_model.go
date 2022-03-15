package pengerjaan

import "time"

type PengerjaanModel struct {
	ID              int       `json:"id"`
	Status          string    `json:"status" validate:"required"`
	Tanggal         string    `json:"tanggal" validate:"required"`
	NamaSPK         string    `json:"nama_spk" validate:"required"`
	Pengaduan       string    `json:"pengaduan" validate:"required"`
	Lokasi          string    `json:"lokasi" validate:"required"`
	Detail          string    `json:"detail" validate:"required"`
	EstimasiSelesai string    `json:"estimasi_selesai"`
	PenanggungJawab string    `json:"penanggung_jawab" validate:"required"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type ResponseModel struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Filter struct {
	Limit string
}
