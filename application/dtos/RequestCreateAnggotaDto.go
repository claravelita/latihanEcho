package dtos

type RequestCreateAnggotaDto struct {
	NamaAnggota   string `json:"nama_anggota"`
	AlamatAnggota string `json:"alamat_anggota"`
	NoTelpAnggota string `json:"no_telp_anggota"`
	Status        string `json:"status"`
}
