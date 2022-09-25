package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	open, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_restful_api?parseTime=true")
	helper.PanicError(err)

	open.SetMaxIdleConns(10)                  // Pengaturan berapa jumlah minimal koneksi yang dibuat / init connection
	open.SetMaxOpenConns(100)                 // Pengaturan berapa jumlah maksimal koneksi yang dibuat
	open.SetConnMaxIdleTime(5 * time.Minute)  // Pengaturan berapa lama koneksi yang tidak digunakan akan dihapus
	open.SetConnMaxLifetime(60 * time.Minute) //Pengaturan berapa lama koneksi boleh digunakan

	return open
}
