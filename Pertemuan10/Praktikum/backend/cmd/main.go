package cmd

import (
	"context"
	"w4backend/model"
	"w4backend/repository"
)

func main() {
	n := 0
	for {
		if n > 1000 {
			return
		}
		repository.InsertMahasiswa(context.Background(), model.Mahasiswa{
			Nama:     "adit ganteng",
			NPM:      n,
			Prodi:    "zakar",
			Fakultas: "reproduksi",
			Alamat: model.Alamat{
				Jalan:     "Jl. cepat kemari No.1",
				Kelurahan: "balik papan",
				Kota:      "bawah tanah",
			},
			Minat: []string{"beranak"},
			MataKuliah: []model.MataKuliah{
				{Kode: "UT101", Nama: "Dasar Testing", Nilai: 100},
			},
		})
		n++
	}
}
