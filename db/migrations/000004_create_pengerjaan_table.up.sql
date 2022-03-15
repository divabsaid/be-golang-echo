CREATE TABLE IF NOT EXISTS pengerjaan (
    id INT AUTO_INCREMENT PRIMARY KEY,
    status VARCHAR(255) NOT NULL,
    tanggal TIMESTAMP,
    nama_spk VARCHAR(255) NOT NULL,
    pengaduan VARCHAR(255) NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    detail VARCHAR(255) NOT NULL,
    penanggung_jawab VARCHAR(255) NOT NULL,
    estimasi_selesai VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
)