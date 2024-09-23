package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/Dryluigi/golang-todos/controller"
	"github.com/Dryluigi/golang-todos/database"
	"github.com/labstack/echo/v4"
)

// Struct untuk mapping request
type CreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
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

	// agenda route
    controller.NewGetAllAgendaController(e, db)
	controller.GetDataAgendaByIdController(e, db)
	controller.PostAgendaController(e, db)
	controller.UpdateAgendaByIdController(e, db)
	controller.DeleteAgendaByIdController(e, db)

	// departemen route
	controller.PostDepartemenController(e, db)
	controller.GetAllDepartemenController(e, db)
	controller.GetDepartemenByIdController(e, db)
	controller.UpdateDepartemenController(e, db)
	controller.DeleteDepartemenDataController(e, db)

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

	// Jalankan server di port 8000
	log.Println("Server dimulai di port 8000")
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}

}
