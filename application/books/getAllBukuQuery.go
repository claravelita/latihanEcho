package books

import (
	"latihanEcho/application/dtos"
	"latihanEcho/domain/entities"
)

func (b BukuServices) GetAllBukuQuery() dtos.ResponsesDto {
	responses := dtos.ResponsesDto{}
	var dataBuku []entities.Buku

	err := b.database.Find(&dataBuku).Error
	if err != nil {
		responses.Code = 500
		responses.Status = "Internal server error, " + err.Error()
		return responses
	}

	responses.Status = "Succeed"
	responses.Code = 200
	responses.Data = dataBuku

	return responses
}
