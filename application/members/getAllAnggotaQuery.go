package members

import (
	"latihanEcho/application/dtos"
	"latihanEcho/domain/entities"
)

func (a AnggotaService) GetAllAnggotaQuery() dtos.ResponsesDto {
	responses := dtos.ResponsesDto{}
	var data []entities.Anggota

	err := a.database.Find(&data).Error
	if err != nil {
		responses.Code = 500
		responses.Status = "Internal server error, " + err.Error()
		return responses
	}

	responses.Status = "Succeed"
	responses.Code = 200
	responses.Data = data

	return responses
}
