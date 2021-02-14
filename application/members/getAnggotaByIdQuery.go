package members

import (
	"gorm.io/gorm"
	"latihanEcho/application/dtos"
	"latihanEcho/domain/entities"
)

func (a AnggotaService) GetAnggotaByIdQuery(id int) dtos.ResponsesDto {
	var data entities.Anggota
	responses := dtos.ResponsesDto{}

	err := a.database.First(&data, id).Error
	if err == gorm.ErrRecordNotFound {
		responses.Code = 404
		responses.Status = "Buku is not found!"
		return responses
	}

	responses.Code = 200
	responses.Status = "Succeed!"
	responses.Data = data

	return responses
}
