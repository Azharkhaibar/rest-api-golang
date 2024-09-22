package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dryluigi/golang-todos/database"
	"github.com/labstack/echo/v4"
)

// Struct untuk mapping request
type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type AgendaRequest struct {
	NamaAgenda      string `json:"nama_agenda"`
	HariPelaksanaan string `json:"hari_pelaksanaan"`
	NamaPelaksana   string `json:"nama_pelaksana"`
}

type AgendaResponse struct {
	id              int `json:"id"`
	namaAgenda      string `json:"nama_agenda"`
	HariPelaksanaan string `json:"hari_pelaksanaan"`
	NamaPelaksana   string `json:"nama_pelaksana"`
	Done            bool `json:"done"`
}

func main() {
	// Inisialisasi database dan tangani error yang mungkin terjadi
	db, err := database.InitDb()
	if err != nil {
		log.Fatalf("Gagal menginisialisasi database: %v", err)
	}

	// Pastikan koneksi database ditutup setelah aplikasi selesai
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("Gagal menutup koneksi database: %v", err)
		}
	}()

	// Cek koneksi database
	err = db.Ping()
	if err != nil {
		log.Fatalf("Gagal ping ke database: %v", err)
	}

	// Inisialisasi Echo web server
	e := echo.New()

	// Handler POST untuk route /todos
	e.POST("/todos", func(ctx echo.Context) error {
		// Parsing JSON dari request body
		var request CreateRequest
		if err := json.NewDecoder(ctx.Request().Body).Decode(&request); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid request payload")
		}
		fmt.Println("Title:", request.Title, "Description:", request.Description)

		// INSERT query ke database
		_, err := db.Exec(
			"INSERT INTO todos (title, description) VALUES (?, ?)",
			request.Title,
			request.Description,
		)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "Todo created")
	})

	// Handler POST untuk route /agenda
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
			Agenda.id = id
			Agenda.namaAgenda = nama_agenda
			Agenda.HariPelaksanaan = hari_pelaksanaan
			Agenda.NamaPelaksana = nama_pelaksana
			if done == 1 {
				Agenda.Done = true
			}
			res = append(res, Agenda)
		}
		return ctx.JSON(http.StatusOK, res)
	})

	// Jalankan server di port 8000
	log.Println("Server dimulai di port 8000")
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}
}
