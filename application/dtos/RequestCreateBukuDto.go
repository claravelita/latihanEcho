package dtos

type RequestCreateBukuDto struct {
	// Data Transfer Object
	NamaBuku   string `json:"nama_buku"`
	TipeBuku   string `json:"tipe_buku"`
	JumlahBuku int    `json:"jumlah_buku"`
	Pengarang  string `json:"pengarang"`
}
