package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"latihanEcho/application/books"
	"latihanEcho/application/dtos"
	"strconv"
)

type bukuController struct {
	services *books.BukuServices
}

func NewBukuController(services *books.BukuServices) *bukuController {
	return &bukuController{
		services: services,
	}
}

func (b bukuController) AddBukuController(e *echo.Echo) {
	e.GET("/buku", b.handleGetBuku)
	e.GET("/buku/:id", b.handleGetIdBuku)
	e.POST("/buku", b.handlePostBuku)
	e.DELETE("/buku/:id", b.handleDeleteBuku)
	e.PUT("/buku/:id", b.handlePutBuku)
}

func (b bukuController) handleGetBuku(ctx echo.Context) error {
	responses := b.services.GetAllBukuQuery()
	return ctx.JSON(responses.Code, responses)
}

func (b bukuController) handleGetIdBuku(ctx echo.Context) error {
	id := ctx.Param("id")
	converted, err := strconv.Atoi(id)

	if err != nil {
		return ctx.JSON(400, "Id is not valid format")
	}

	responses := b.services.GetBukuByIdQuery(converted)

	return ctx.JSON(responses.Code, responses)
}

func (b bukuController) handlePostBuku(ctx echo.Context) error {
	request := dtos.RequestCreateBukuDto{}

	if err := ctx.Bind(&request); err != nil {
		fmt.Println(err)
		return ctx.JSON(400, "Beberapa tipe data tidak sesuai!")
	}

	if request.NamaBuku == "" {
		return ctx.JSON(400, "Nama harus diisi!")
	}
	if request.JumlahBuku == 0 {
		return ctx.JSON(400, "Jumlah buku harus diisi!")
	}
	if request.Pengarang == "" {
		return ctx.JSON(400, "Nama pengarang harus diisi!")
	}
	if request.TipeBuku == "" {
		return ctx.JSON(400, "Tipe buku harus diisi!")
	}

	responses := b.services.CreateBukuCommand(request)

	return ctx.JSON(responses.Code, responses)
}

func (b bukuController) handleDeleteBuku(ctx echo.Context) error {
	id := ctx.Param("id")
	converted, err := strconv.Atoi(id)

	if err != nil {
		return ctx.JSON(400, "Id is not valid format")
	}

	responses := b.services.DeleteBukuCommand(converted)

	return ctx.JSON(responses.Code, responses)
}

func (b bukuController) handlePutBuku(ctx echo.Context) error {
	request := dtos.RequestCreateBukuDto{}

	if err := ctx.Bind(&request); err != nil {
		fmt.Println(err)
		return ctx.JSON(400, "Beberapa tipe data tidak sesuai!")
	}

	if request.NamaBuku == "" {
		return ctx.JSON(400, "Nama harus diisi!")
	}
	if request.JumlahBuku == 0 {
		return ctx.JSON(400, "Jumlah buku harus diisi!")
	}
	if request.Pengarang == "" {
		return ctx.JSON(400, "Nama pengarang harus diisi!")
	}
	if request.TipeBuku == "" {
		return ctx.JSON(400, "Tipe buku harus diisi!")
	}

	id := ctx.Param("id")
	converted, err := strconv.Atoi(id)

	if err != nil {
		return ctx.JSON(400, "Id is not valid format")
	}

	responses := b.services.UpdateBukuCommand(request, converted)

	return ctx.JSON(responses.Code, responses)

}
