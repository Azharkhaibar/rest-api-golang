package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ClientCustomer struct {
	Id          int    `json:"id"`
	NamaClient  string `json:"nama_client"`
	NomerTelpon int    `json:"nomer_telpon"`
	Domisili    string `json:"domisili"`
	OpsiLayanan string `json:"opsi_layanan"`
	TotalHarga  int    `json:"total_harga_bayar"`
	Done        *bool   `json:"done"`
}

func PostClientCustomerController(e *echo.Echo, db *sql.DB) {
    e.POST("/client", func(ctx echo.Context) error {
		var Client ClientCustomer
		if err := json.NewDecoder(ctx.Request().Body).Decode(&Client); err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid REQ Payload")
		}
		fmt.Println("namaKlien:", Client.NamaClient, "NomerTelpon:", Client.NomerTelpon, "Domisili:", Client.Domisili, "OpsiLayanan:", Client.OpsiLayanan, "TotalHargaBayar:", Client.TotalHarga)
		_, err := db.Exec(
			"INSERT INTO client_customer (nama_client, nomer_telpon, domisili, opsi_layanan, total_harga_bayar) VALUES (?, ?, ?, ?, ?)",
			Client.NamaClient,
			Client.NomerTelpon,
			Client.Domisili,
			Client.OpsiLayanan,
			Client.TotalHarga,
		)
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.String(http.StatusOK, "Client Customer data Created!")
	})
} 

func GetAllDataClientCustomerController(e *echo.Echo, db *sql.DB) {
	e.GET("/client", func(ctx echo.Context) error {
		dbRows, err := db.Query("SELECT * FROM client_customer")
		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		defer dbRows.Close()
		var GetClient []ClientCustomer
		for dbRows.Next() {
			var Id int
			var NamaClient, Domisili, OpsiLayanan string
			var NomerTelpon, TotalHarga int
			var Done sql.NullBool
			err := dbRows.Scan(&Id, &NamaClient, &NomerTelpon, &Domisili, &OpsiLayanan, &TotalHarga, &Done)
			if err != nil {
				return ctx.String(http.StatusInternalServerError, err.Error())	
			}
			var DoneValue *bool
			if Done.Valid {
				DoneValue = &Done.Bool
			} else {
				DoneValue = nil
			}
			Client := ClientCustomer{
				Id: Id,
				NamaClient: NamaClient,
				NomerTelpon: NomerTelpon,
				Domisili: Domisili,
				OpsiLayanan: OpsiLayanan,
				TotalHarga: TotalHarga,
				Done: DoneValue,
			}

			GetClient = append(GetClient, Client)
		}
		if len(GetClient) == 0 {
			return ctx.JSON(http.StatusNotFound, echo.Map{"message": "No Client found."})
		}
		return ctx.JSON(http.StatusOK, GetClient)
	})
}

func GetClientCustomerById(e *echo.Echo, db *sql.DB) {
	e.GET("/client/:id", func(ctx echo.Context) error {
		ClientID := ctx.Param("id")
		var ClientGetID ClientCustomer
		err := db.QueryRow(
			"SELECT id, nama_client, nomer_telpon, domisili, opsi_layanan, total_harga_bayar, done FROM client_customer WHERE id = ?", ClientID).
			Scan(&ClientGetID.Id ,&ClientGetID.NamaClient, &ClientGetID.NomerTelpon, &ClientGetID.Domisili, &ClientGetID.OpsiLayanan, &ClientGetID.TotalHarga, &ClientGetID.Done)
		if err != nil {
			if err == sql.ErrNoRows {
				return ctx.String(http.StatusNotFound, "Client ID not found")
			}
			return ctx.String(http.StatusInternalServerError, err.Error())
		}
		return ctx.JSON(http.StatusOK, ClientGetID)
	})
}