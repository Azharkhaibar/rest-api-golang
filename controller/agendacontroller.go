package controller

import (
	"database/sql"
	"net/http"
    "encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
)

type AgendaResponse struct {
	Id              int    `json:"id"`
	NamaAgenda      string `json:"nama_agenda"`
	HariPelaksanaan string `json:"hari_pelaksanaan"`
	NamaPelaksana   string `json:"nama_pelaksana"`
	Done            bool   `json:"done"`
}

type AgendaRequest struct {
	NamaAgenda      string `json:"nama_agenda"`
	HariPelaksanaan string `json:"hari_pelaksanaan"`
	NamaPelaksana   string `json:"nama_pelaksana"`
}

type AgendaUpdate struct {
	NamaAgenda      string `json:"nama_agenda"`
	HariPelaksanaan string `json:"hari_pelaksanaan"`
	NamaPelaksana   string `json:"nama_pelaksana"`
}

func NewGetAllAgendaController(e *echo.Echo, db *sql.DB) {
	e.GET("/agenda", func(ctx echo.Context) error {
		dbRows, err := db.Query(
			"SELECT * FROM agenda")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		var res []AgendaResponse

		// menggunakan loops unutk cetak rows
		for dbRows.Next() {
			var id int
			var nama_agenda string
			var hari_pelaksanaan string
			var nama_pelaksana string
			var done int

			err := dbRows.Scan(&id, &nama_agenda, &hari_pelaksanaan, &nama_pelaksana, &done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())
			}

			// array response
			var Agenda AgendaResponse
			Agenda.Id = id
			Agenda.NamaAgenda = nama_agenda
			Agenda.HariPelaksanaan = hari_pelaksanaan
			Agenda.NamaPelaksana = nama_pelaksana
			if done == 1 {
				Agenda.Done = true
			}
			res = append(res, Agenda)
		}
		return ctx.JSON(http.StatusOK, res)
	})
}

func GetDataAgendaByIdController(e *echo.Echo, db *sql.DB) {
	e.GET("/agenda/:id", func(ctx echo.Context) error {
		agendaID := ctx.Param("id")
		var agenda AgendaResponse
		err := db.QueryRow(
			"SELECT id, nama_agenda, hari_pelaksanaan, nama_pelaksana, done FROM agenda WHERE id = ?", agendaID).
			Scan(&agenda.Id, &agenda.NamaAgenda, &agenda.HariPelaksanaan, &agenda.NamaPelaksana, &agenda.Done)

		if err != nil {
			if err == sql.ErrNoRows {
				return ctx.String(http.StatusNotFound, "Agenda Not Found")
			}
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, agenda)
	})
}

func PostAgendaController(e *echo.Echo, db *sql.DB) {
	e.POST("/agenda", func(ctx echo.Context) error {
		var requestAgenda AgendaRequest
		if err := json.NewDecoder(ctx.Request().Body).Decode(&requestAgenda); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid request payload")
		}
		fmt.Println("NamaAgenda:", requestAgenda.NamaAgenda, "HariPelaksanaan:", requestAgenda.HariPelaksanaan, "NamaPelaksana:", requestAgenda.NamaPelaksana)

		// INSERT query ke database
		_, err := db.Exec(
			"INSERT INTO agenda (nama_agenda, hari_pelaksanaan, nama_pelaksana) VALUES (?, ?, ?)",
			requestAgenda.NamaAgenda,
			requestAgenda.HariPelaksanaan,
			requestAgenda.NamaPelaksana,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "Agenda created")
	})
}

func UpdateAgendaByIdController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/agenda/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		var UpdateAgenda AgendaUpdate
		if err := ctx.Bind(&UpdateAgenda); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid request body")
		}

		// INSERT query ke database
		_, err := db.Exec(
			"UPDATE agenda SET nama_agenda = ?, hari_pelaksanaan = ?, nama_pelaksana = ? WHERE id = ?",
			UpdateAgenda.NamaAgenda,
			UpdateAgenda.HariPelaksanaan,
			UpdateAgenda.NamaPelaksana,
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "Agenda Updated")
	})
}

func DeleteAgendaByIdController(e *echo.Echo, db *sql.DB) {
	e.DELETE("/agenda/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		// INSERT query ke database
		_, err := db.Exec(
			"DELETE FROM agenda where id = ?",
			id,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "Agenda Id Deleted")
	})
}
