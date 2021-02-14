package entities

type Anggota struct {
	IdAnggota     int    `json:"id_anggota" gorm:"PRIMARY_KEY;AUTO_INCREMENT;size:11"`
	NamaAnggota   string `json:"nama_anggota" gorm:"size:30"`
	AlamatAnggota string `json:"alamat_anggota" gorm:"size:40"`
	NoTelpAnggota string `json:"no_telp_anggota" gorm:"size:15"`
	Status        string `json:"status" gorm:"size:15"`
}
