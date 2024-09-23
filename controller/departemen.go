package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Departemen struct {
	Id                     int    `json:"id"`
	Nama_departemen        string `json:"nama_departemen"`
	Nama_kepala_departemen string `json:"nama_kepala_departemen"`
	Kantor_departemen      string `json:"kantor_departemen"`
	Total_karyawan         int    `json:"total_karyawan"`
	Laba_departemen        int    `json:"laba_departemen"`
}

func PostDepartemenController(e *echo.Echo, db *sql.DB) {
	e.POST("/departemen", func(ctx echo.Context) error {
		var PostDepartemen Departemen
		if err := json.NewDecoder(ctx.Request().Body).Decode(&PostDepartemen); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid Request Payload")
		}
		fmt.Println("namaDepartemen:", PostDepartemen.Nama_departemen, "KepalaDepartemen:", PostDepartemen.Nama_kepala_departemen, "kantorDepartmen:", PostDepartemen.Kantor_departemen, "totalKaryawan:", PostDepartemen.Total_karyawan, "labaDepartemen:", PostDepartemen.Laba_departemen)

		_, err := db.Exec(
			"INSERT INTO departemen (nama_departemen, nama_kepala_departemen, kantor_departemen, total_karyawan, laba_departemen) VALUES (?, ?, ?, ?, ?)",
			PostDepartemen.Nama_departemen,
			PostDepartemen.Nama_kepala_departemen,
			PostDepartemen.Kantor_departemen,
			PostDepartemen.Total_karyawan,
			PostDepartemen.Laba_departemen,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "Departemen Created!")
	})
}

// GET ALL DEPARTEMEN

func GetAllDepartemenController(e *echo.Echo, db *sql.DB) error {
    
}

// GET DEPARTEMEN BY ID

// UPDATE DEPARTEMEN

// DELETE DEPARTEMEN
