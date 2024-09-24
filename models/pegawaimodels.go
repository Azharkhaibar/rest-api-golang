package models

import "database/sql"

type Pegawai struct {
	Id                int    `json:"id"`
	NamaPegawai       string `json:"nama_pegawai"`
	JabatanPegawai    string `json:"jabatan_pegawai"`
	GajiPegawai       int    `json:"gaji_pegawai"`
	StatusPegawai     string `json:"status_pegawai"`
	DepartemenPegawai string `json:"departemen_pegawai"`
	EmailPegawai      string `json:"email_pegawai"`
	Done              *bool  `json:"done"`
}

func CreatePegawai(db *sql.DB, pegawai Pegawai) error {
	_, err := db.Exec(
		"INSERT INTO pegawai_departemen (nama_pegawai, jabatan_pegawai, gaji_pegawai, status_pegawai, departemen_pegawai, email_pegawai) VALUES (?, ?, ?, ?, ?, ?)",
		pegawai.NamaPegawai,
		pegawai.JabatanPegawai,
		pegawai.GajiPegawai,
		pegawai.StatusPegawai,
		pegawai.DepartemenPegawai,
		pegawai.EmailPegawai,
	)
	return err
}

// []Pegawai
func GetAllPegawaiDataModels(db *sql.DB) ([]Pegawai, error) {
	RowsPegawai, err := db.Query("SELECT id, nama_pegawai, jabatan_pegawai, gaji_pegawai, status_pegawai, departemen_pegawai, email_pegawai, done FROM pegawai")
	if err != nil {
		return nil, err
	}
	defer RowsPegawai.Close()

	var getPegawai []Pegawai
	for RowsPegawai.Next() {
		var DataPegawai Pegawai
		err := RowsPegawai.Scan(&DataPegawai.Id, &DataPegawai.NamaPegawai, &DataPegawai.JabatanPegawai, &DataPegawai.GajiPegawai, &DataPegawai.StatusPegawai, &DataPegawai.DepartemenPegawai, &DataPegawai.EmailPegawai, &DataPegawai.Done)
		if err != nil {
			return nil, err
		}
		getPegawai = append(getPegawai, DataPegawai)
	}
	return getPegawai, nil
}

// id int
func GetPegawaiById(db *sql.DB, id int) (Pegawai, error) {
	var GetIdPegawai Pegawai
	err := db.QueryRow("SELECT id, nama_pegawai, jabatan_pegawai, gaji_pegawai, status_pegawai, departemen_pegawai, email_pegawai, done FROM pegawai WHERE id = ?", id).
		Scan(&GetIdPegawai.Id, &GetIdPegawai.NamaPegawai, &GetIdPegawai.JabatanPegawai, &GetIdPegawai.GajiPegawai, &GetIdPegawai.StatusPegawai, &GetIdPegawai.DepartemenPegawai, &GetIdPegawai.EmailPegawai, &GetIdPegawai.Done)
	if err != nil {
		return GetIdPegawai, err
	}
	return GetIdPegawai, nil
}
