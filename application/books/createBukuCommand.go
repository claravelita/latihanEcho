package books

import (
	"latihanEcho/application/dtos"
	"latihanEcho/domain/entities"
)

func (b BukuServices) CreateBukuCommand(request dtos.RequestCreateBukuDto) dtos.ResponsesDto {
	responses := dtos.ResponsesDto{}

	newEntity := entities.Buku{
		JumlahBuku: request.JumlahBuku,
		NamaBuku:   request.NamaBuku,
		Pengarang:  request.Pengarang,
		TipeBuku:   request.TipeBuku,
	}

	err := b.database.Save(&newEntity).Error
	if err != nil {
		responses.Code = 500
		responses.Status = "Internal Server Error, " + err.Error()
		return responses
	}

	responses.Code = 200
	responses.Status = "Succeed!"

	return responses
}
