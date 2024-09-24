package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Departemen struct {
	Id                   int    `json:"id"`
	NamaDepartemen       string `json:"nama_departemen"`
	NamaKepalaDepartemen string `json:"nama_kepala_departemen"`
	KantorDepartemen     string `json:"kantor_departemen"`
	TotalKaryawan        int    `json:"total_karyawan"`
	LabaDepartemen       int    `json:"laba_departemen"`
	Done                 bool   `json:"done"`
}

type UpdateDepartemen struct {
	Id                   int    `json:"id"`
	NamaDepartemen       string `json:"nama_departemen"`
	NamaKepalaDepartemen string `json:"nama_kepala_departemen"`
	KantorDepartemen     string `json:"kantor_departemen"`
	TotalKaryawan        int    `json:"total_karyawan"`
	LabaDepartemen       int    `json:"laba_departemen"`
	Done                 bool   `json:"done"`
}

func PostDepartemenController(e *echo.Echo, db *sql.DB) {
	e.POST("/departemen", func(ctx echo.Context) error {
		var PostDepartemen Departemen
		if err := json.NewDecoder(ctx.Request().Body).Decode(&PostDepartemen); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid Request Payload")
		}
		fmt.Println("namaDepartemen:", PostDepartemen.NamaDepartemen, "KepalaDepartemen:", PostDepartemen.NamaKepalaDepartemen, "kantorDepartmen:", PostDepartemen.KantorDepartemen, "totalKaryawan:", PostDepartemen.TotalKaryawan, "labaDepartemen:", PostDepartemen.LabaDepartemen)

		_, err := db.Exec(
			"INSERT INTO departemen (nama_departemen, nama_kepala_departemen, kantor_departemen, total_karyawan, laba_departemen) VALUES (?, ?, ?, ?, ?)",
			PostDepartemen.NamaDepartemen,
			PostDepartemen.NamaKepalaDepartemen,
			PostDepartemen.KantorDepartemen,
			PostDepartemen.TotalKaryawan,
			PostDepartemen.LabaDepartemen,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "Departemen Created!")
	})
}

// GET ALL DEPARTEMEN DATA
func GetAllDepartemenController(e *echo.Echo, db *sql.DB) {
	e.GET("/departemen", func(ctx echo.Context) error {
		dbRows, err := db.Query("SELECT * FROM departemen")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		defer dbRows.Close()
		var departments []Departemen
		for dbRows.Next() {
			var id int
			var namaDepartemen, namaKepalaDepartemen, kantorDepartemen string
			var totalKaryawan, labaDepartemen, done int
			err := dbRows.Scan(&id, &namaDepartemen, &namaKepalaDepartemen, &kantorDepartemen, &totalKaryawan, &labaDepartemen, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			department := Departemen{
				Id:                   id,
				NamaDepartemen:       namaDepartemen,
				NamaKepalaDepartemen: namaKepalaDepartemen,
				KantorDepartemen:     kantorDepartemen,
				TotalKaryawan:        totalKaryawan,
				LabaDepartemen:       labaDepartemen,
				Done:                 done == 1,
			}

			departments = append(departments, department)
		}
		if len(departments) == 0 {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": "No departments found."})
		}
		return ctx.JSON(http.StatusOK, departments)
	})
}

// GET DEPARTEMEN BY ID
func GetDepartemenByIdController(e *echo.Echo, db *sql.DB) {
	e.GET("/departemen/:id", func(ctx echo.Context) error {
		departemenByID := ctx.Param("id")
		var Departemens Departemen
		err := db.QueryRow(
			"SELECT id, nama_departemen, nama_kepala_departemen, kantor_departemen, total_karyawan, laba_departemen, done FROM departemen WHERE id = ? ", departemenByID).
			Scan(&Departemens.Id, &Departemens.NamaDepartemen, &Departemens.NamaKepalaDepartemen, &Departemens.KantorDepartemen, &Departemens.TotalKaryawan, &Departemens.LabaDepartemen, &Departemens.Done)
        if err != nil {
			if err == sql.ErrNoRows {
				return ctx.String(http.StatusNotFound, "Departemen Data Not Found!")
			}
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, Departemens)
	})
}
// UPDATE DEPARTEMEN
func UpdateDepartemenController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/departemen/:id", func(ctx echo.Context) error {
		DepartemenID := ctx.Param("id")
		var UpdateDepartemen UpdateDepartemen
		if err := ctx.Bind(&UpdateDepartemen); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid request payload")
		}

		_, err := db.Exec(
			"UPDATE departemen SET nama_departemen = ?, nama_kepala_departemen = ?, kantor_departemen = ?, total_karyawan = ?, laba_departemen = ? WHERE id = ?",
			UpdateDepartemen.NamaDepartemen,
			UpdateDepartemen.NamaKepalaDepartemen,
			UpdateDepartemen.KantorDepartemen,
			UpdateDepartemen.TotalKaryawan,
			UpdateDepartemen.LabaDepartemen,
			DepartemenID,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "Departemen Updated")
	})
}

// DELETE DEPARTEMEN
func DeleteDepartemenDataController(e *echo.Echo, db *sql.DB) {
    e.DELETE("/departemen/:id", func(ctx echo.Context) error {
		DepartemenID := ctx.Param("id")
		_, err := db.Exec(
			"DELETE FROM departemen WHERE id = ?",
			DepartemenID,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "Departemen ID Deleted")
	})
}
