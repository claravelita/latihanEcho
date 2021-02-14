package main

import (
	"github.com/labstack/echo/v4"
	"latihanEcho/api/controller"
	"latihanEcho/application/books"
	"latihanEcho/application/members"
	"latihanEcho/infrastructure/persistence"
)

func main() {
	echoServer := echo.New()

	initGorm := persistence.NewDbConnection()

	newBukuService := books.NewBukuService(initGorm)
	controllerBuku := controller.NewBukuController(newBukuService)
	controllerBuku.AddBukuController(echoServer)

	newAnggotaService := members.NewAnggotaService(initGorm)
	controllerAnggota := controller.NewAnggotaController(newAnggotaService)
	controllerAnggota.AddAnggotaController(echoServer)

	_ = echoServer.Start(":5300")
}
