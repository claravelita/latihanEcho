package members

import "gorm.io/gorm"

type AnggotaService struct {
	database *gorm.DB
}

func NewAnggotaService(gorm *gorm.DB) *AnggotaService {
	return &AnggotaService{
		database: gorm,
	}

}
