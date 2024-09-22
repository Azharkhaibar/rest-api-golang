package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Dryluigi/golang-todos/database"
	"github.com/labstack/echo/v4"
)

// mapping request pake struct
type CreateRequest struct {
	Title       string `json: "title"`
	Desctiption string `json: "description"`
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

	// Cek jika koneksi database hidup
	err = db.Ping()
	if err != nil {
		log.Fatalf("Gagal ping ke database: %v", err)
	}

	// Inisialisasi Echo web server
	e := echo.New()

	// Handler POST untuk route /todos_golang
	e.POST("/todos", func(ctx echo.Context) error {
		// parsing json
		var Request CreateRequest
		json.NewDecoder(ctx.Request().Body).Decode(&Request)
		fmt.Println(Request)

		// INSERT QUERY GOLANG

		return ctx.String(http.StatusOK, "OK")
	})

	// Log ketika server dimulai
	log.Println("Server dimulai di port 8000")
	if err := e.Start(":8000"); err != nil {
		log.Fatalf("Gagal memulai server: %v", err)
	}

	fmt.Println("Go berhasil dijalankan")
}
