package models

import "database/sql"

type JasaWeb struct {
	Id          int    `json:"id"`
	NamaWeb     string `json:"nama_web"`
	KategoriWeb string `json:"kategori_web"`
	HargaWeb    int    `json:"harga_web"`
	NamaClient  string `json:"nama_client"`
	PageWeb     int    `json:"page_web"`
	Done        *bool  `json:"done"`
}

func GetAllJasaWebDataModels(db *sql.DB) ([]JasaWeb, error) {
	RowsJasaWeb, err := db.Query("SELECT id, nama_web, kategori_web, harga_web, nama_client, page_web, done")
	if err != nil {
		return nil, err
	}
	defer RowsJasaWeb.Close()
	var jasaweb []JasaWeb
	for RowsJasaWeb.Next() {
		var getjasaweb JasaWeb
		err := RowsJasaWeb.Scan(&getjasaweb.Id, &getjasaweb.NamaWeb, &getjasaweb.KategoriWeb, &getjasaweb.HargaWeb, &getjasaweb.NamaClient, &getjasaweb.PageWeb, &getjasaweb.Done)
		if err != nil {
			return nil, err
		}
		jasaweb = append(jasaweb, getjasaweb)
	}
	return jasaweb, nil
}

func GetJasaWebDataModels(db *sql.DB, id int) (JasaWeb, error) {
	var GetIdJasaWeb JasaWeb
	err := db.QueryRow("SELECT id, nama_web, kategori_web, harga_web, nama_client, page_web, done", id).
		Scan(&GetIdJasaWeb.Id, GetIdJasaWeb.NamaWeb, &GetIdJasaWeb.KategoriWeb, &GetIdJasaWeb.HargaWeb, &GetIdJasaWeb.NamaClient, &GetIdJasaWeb.PageWeb, &GetIdJasaWeb.Done)
    if err != nil {
		return GetIdJasaWeb, err
	}
	return GetIdJasaWeb, nil
}
