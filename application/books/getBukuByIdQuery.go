package books

import (
	"gorm.io/gorm"
	"latihanEcho/application/dtos"
	"latihanEcho/domain/entities"
)

func (b BukuServices) GetBukuByIdQuery(id int) dtos.ResponsesDto {
	var data entities.Buku
	responses := dtos.ResponsesDto{}

	err := b.database.First(&data, id).Error

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
