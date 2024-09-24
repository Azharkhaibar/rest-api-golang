package controller

import (
	"database/sql"
	"net/http"
	"strconv"
	"github.com/Dryluigi/golang-todos/models"
	"github.com/labstack/echo/v4"
)

func PostPegawaiController(e *echo.Echo, db *sql.DB) {
	e.POST("/pegawai", func(ctx echo.Context) error {
		var postPegawai models.Pegawai
		if err := ctx.Bind(&postPegawai); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid Request")
		}
		if err := models.CreatePegawai(db, postPegawai); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "Successfully Post Pegawai")
	})
}

func GetAllPegawaiDataController(e *echo.Echo, db *sql.DB) {
	e.GET("/pegawai", func(ctx echo.Context) error {
		getPegawai, err := models.GetAllPegawaiDataModels(db)
		if err != nil {
			return ctx.String(http.StatusBadRequest, err.Error())
		}
		if len(getPegawai) == 0 {
			return ctx.JSON(http.StatusNotFound, echo.Map{
				"message": "Pegawai Data Not Found",
			})
		}
		return ctx.JSON(http.StatusOK, getPegawai)
	})
}

func GetPegawaiByIdController(e *echo.Echo, db *sql.DB) {
	e.GET("/pegawai/:id", func(ctx echo.Context) error {
		getPegawaiID := ctx.Param("id")
		id, err := strconv.Atoi(getPegawaiID)
		if err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid Pegawai ID!")
		}
		pegawai, err := models.GetPegawaiById(db, id)
		if err != nil {
			if err == sql.ErrNoRows {
				return ctx.String(http.StatusNotFound, "Id Pegawai not found")
			}
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, pegawai)
	})
}