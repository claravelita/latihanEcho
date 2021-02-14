package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"latihanEcho/application/dtos"
	"latihanEcho/application/members"
	"strconv"
)

type anggotaController struct {
	services *members.AnggotaService
}

func NewAnggotaController(services *members.AnggotaService) *anggotaController {
	return &anggotaController{
		services: services,
	}
}

func (a anggotaController) AddAnggotaController(e *echo.Echo) {

	e.GET("/anggota", a.handleGetAnggota)
	e.GET("/anggota/:id", a.handleGetIdAnggota)
	e.POST("/anggota", a.handlePostAnggota)
	e.PUT("/anggota/:id", a.handlePutAnggota)
	e.DELETE("/anggota/:id", a.handleDeleteAnggota)
}

func (a anggotaController) handleGetAnggota(ctx echo.Context) error {
	responses := a.services.GetAllAnggotaQuery()

	return ctx.JSON(responses.Code, responses)
}

func (a anggotaController) handleGetIdAnggota(ctx echo.Context) error {
	id := ctx.Param("id")
	converted, err := strconv.Atoi(id)

	if err != nil {
		return ctx.JSON(400, "Id is not valid format")
	}

	responses := a.services.GetAnggotaByIdQuery(converted)

	return ctx.JSON(responses.Code, responses)
}

func (a anggotaController) handlePostAnggota(ctx echo.Context) error {
	requests := dtos.RequestCreateAnggotaDto{}

	if err := ctx.Bind(&requests); err != nil {
		fmt.Println(err)
		return ctx.JSON(400, "Beberapa tipe data tidak sesuai!")
	}

	if requests.NamaAnggota == "" {
		return ctx.JSON(400, "Nama Anggota harus diisi!")
	}
	if requests.AlamatAnggota == "" {
		return ctx.JSON(400, "Alamat Anggota harus diisi!")
	}
	if requests.NoTelpAnggota == "" {
		return ctx.JSON(400, "No. Telp Anggota harus diisi!")
	}

	if requests.Status == "" {
		return ctx.JSON(400, "Status harus diisi!")
	} else if requests.Status != "aktif" && requests.Status != "tidak aktif" {
		return ctx.JSON(400, "Status hanya dapat diisi aktif atau tidak aktif")
	}

	responses := a.services.CreateAnggotaCommand(requests)

	return ctx.JSON(responses.Code, responses)
}

func (a anggotaController) handlePutAnggota(ctx echo.Context) error {
	requests := dtos.RequestCreateAnggotaDto{}

	if err := ctx.Bind(&requests); err != nil {
		fmt.Println(err)
		return ctx.JSON(400, "Beberapa tipe data tidak sesuai!")
	}

	if requests.NamaAnggota == "" {
		return ctx.JSON(400, "Nama Anggota harus diisi!")
	}
	if requests.AlamatAnggota == "" {
		return ctx.JSON(400, "Alamat Anggota harus diisi!")
	}
	if requests.NoTelpAnggota == "" {
		return ctx.JSON(400, "No. Telp Anggota harus diisi!")
	}

	if requests.Status == "" {
		return ctx.JSON(400, "Status harus diisi!")
	} else if requests.Status != "aktif" && requests.Status != "tidak aktif" {
		return ctx.JSON(400, "Status hanya dapat diisi aktif atau tidak aktif")
	}

	id := ctx.Param("id")
	converted, err := strconv.Atoi(id)

	if err != nil {
		return ctx.JSON(400, "Id is not valid format")
	}

	responses := a.services.UpdateAnggotaCommand(requests, converted)

	return ctx.JSON(responses.Code, responses)

}

func (a anggotaController) handleDeleteAnggota(ctx echo.Context) error {
	id := ctx.Param("id")
	converted, err := strconv.Atoi(id)

	if err != nil {
		return ctx.JSON(400, "Id is not valid format")
	}

	responses := a.services.DeleteAnggotaCommand(converted)

	return ctx.JSON(responses.Code, responses)
}
