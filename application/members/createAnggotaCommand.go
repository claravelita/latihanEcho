package members

import (
	"latihanEcho/application/dtos"
	"latihanEcho/domain/entities"
)

func (a AnggotaService) CreateAnggotaCommand(requests dtos.RequestCreateAnggotaDto) dtos.ResponsesDto {
	responses := dtos.ResponsesDto{}

	newEntity := entities.Anggota{
		NamaAnggota:   requests.NamaAnggota,
		AlamatAnggota: requests.AlamatAnggota,
		NoTelpAnggota: requests.NoTelpAnggota,
		Status:        requests.Status,
	}

	err := a.database.Save(&newEntity).Error
	if err != nil {
		responses.Code = 500
		responses.Status = "Internal Server Error, " + err.Error()
		return responses
	}

	responses.Code = 200
	responses.Status = "Succeed!"

	return responses
}
