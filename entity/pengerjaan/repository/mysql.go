package repository

import (
	"be-golang-echo/entity/pengerjaan"
	"database/sql"
)

type PengerjaanRepository interface {
	GetList(filter *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error)
	Add(p *pengerjaan.PengerjaanModel) (bool, error)
	Update(id int, p *pengerjaan.PengerjaanModel) (bool, error)
	Delete(id int) (bool, error)
	GetByID(id int) (*pengerjaan.PengerjaanModel, error)
}

type mysqlPengerjaanRepository struct {
	db *sql.DB
}

func NewMySQLPengerjaanRepository(db *sql.DB) PengerjaanRepository {
	return &mysqlPengerjaanRepository{
		db: db,
	}
}

func (m *mysqlPengerjaanRepository) GetList(filter *pengerjaan.Filter) ([]*pengerjaan.PengerjaanModel, error) {
	rows, err := m.db.Query("SELECT id, status, tanggal, nama_spk, pengaduan, lokasi, detail, estimasi_selesai, penanggung_jawab FROM pengerjaan LIMIT " + filter.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	result := make([]*pengerjaan.PengerjaanModel, 0)
	for rows.Next() {
		r := new(pengerjaan.PengerjaanModel)
		err = rows.Scan(
			&r.ID,
			&r.Status,
			&r.Tanggal,
			&r.NamaSPK,
			&r.Pengaduan,
			&r.Lokasi,
			&r.Detail,
			&r.EstimasiSelesai,
			&r.PenanggungJawab,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, r)
	}
	return result, nil
}

func (m *mysqlPengerjaanRepository) Add(p *pengerjaan.PengerjaanModel) (bool, error) {
	query := "INSERT pengerjaan SET status=?, tanggal=?, nama_spk=?, pengaduan=?, lokasi=?, detail=?, penanggung_jawab=?, estimasi_selesai=?, created_at=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return false, err
	}
	_, err = stmt.Exec(p.Status, p.Tanggal, p.NamaSPK, p.Pengaduan, p.Lokasi, p.Detail, p.PenanggungJawab, p.EstimasiSelesai, p.CreatedAt)
	if err != nil {
		return false, err
	}
	return true, nil

}

func (m *mysqlPengerjaanRepository) Update(id int, p *pengerjaan.PengerjaanModel) (bool, error) {
	query := "UPDATE pengerjaan SET status=?, tanggal=?, nama_spk=?, pengaduan=?, lokasi=?, detail=?, penanggung_jawab=?, estimasi_selesai=?, updated_at=? WHERE id=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return false, err
	}
	res, err := stmt.Exec(p.Status, p.Tanggal, p.NamaSPK, p.Pengaduan, p.Lokasi, p.Detail, p.PenanggungJawab, p.EstimasiSelesai, p.UpdatedAt, id)
	if err != nil {
		return false, err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == 0 || err != nil {
		return false, pengerjaan.UPDATE_FAILED
	}

	return true, nil

}

func (m *mysqlPengerjaanRepository) Delete(id int) (bool, error) {
	query := "DELETE FROM pengerjaan WHERE id=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return false, err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		return false, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return false, pengerjaan.DELETE_FAILED
	}

	return true, nil
}

func (m *mysqlPengerjaanRepository) GetByID(id int) (*pengerjaan.PengerjaanModel, error) {
	userObj := new(pengerjaan.PengerjaanModel)
	row := m.db.QueryRow("SELECT id, status, tanggal, nama_spk, pengaduan, lokasi, detail, penanggung_jawab, estimasi_selesai FROM pengerjaan WHERE id=?", id)
	err := row.Scan(&userObj.ID, &userObj.Status, &userObj.Tanggal, &userObj.NamaSPK, &userObj.Pengaduan, &userObj.Lokasi, &userObj.Detail, &userObj.PenanggungJawab, &userObj.EstimasiSelesai)
	if err != nil {
		return userObj, err
	}
	return userObj, nil
}
