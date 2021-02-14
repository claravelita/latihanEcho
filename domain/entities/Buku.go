package entities

type Buku struct {
	IdBuku     int    `json:"id_buku" gorm:"PRIMARY_KEY;AUTO_INCREMENT;size:11"`
	NamaBuku   string `json:"nama_buku" gorm:"size:30"`
	TipeBuku   string `json:"tipe_buku" gorm:"size:30"`
	JumlahBuku int    `json:"jumlah_buku" gorm:"size:11"`
	Pengarang  string `json:"pengarang" gorm:"size:50"`
}
