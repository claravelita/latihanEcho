package books

import (
	"latihanEcho/application/dtos"
	"latihanEcho/domain/entities"
)

func (b BukuServices) UpdateBukuCommand(request dtos.RequestCreateBukuDto, id int) dtos.ResponsesDto {
	responses := dtos.ResponsesDto{}
	/* Update Pake -Save-

	var data entities.Buku
	err := b.database.First(&data, id).Error

	if err == gorm.ErrRecordNotFound {
		responses.Code = 404
		responses.Status = "Buku is not found!"
		return responses
	}

	data.NamaBuku = request.NamaBuku
	data.TipeBuku = request.TipeBuku
	data.Pengarang = request.Pengarang
	data.JumlahBuku = request.JumlahBuku

	err = b.database.Save(&data).Error
	*/

	newEntity := entities.Buku{
		JumlahBuku: request.JumlahBuku,
		NamaBuku:   request.NamaBuku,
		Pengarang:  request.Pengarang,
		TipeBuku:   request.TipeBuku,
	}
	// Update Pake -Update-
	err := b.database.Where("id_buku = ?", id).Updates(&newEntity).Error
	if err != nil {
		responses.Code = 500
		responses.Status = "Internal Server Error, " + err.Error()
		return responses
	}

	responses.Code = 200
	responses.Status = "Succeed!"

	return responses
}
