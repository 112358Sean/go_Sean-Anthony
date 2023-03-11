package main

type Kendaraan struct {
	totalRoda       int
	kecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

//menambah kecepatanPerJam sebanyak 10 setiap kali dijalankan
func (m *Mobil) TambahKecepatan() {
	m.kecepatanPerJam = m.kecepatanPerJam + 10
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.TambahKecepatan()
	mobilCepat.TambahKecepatan()
	mobilCepat.TambahKecepatan()

	mobilLamban := Mobil{}
	mobilLamban.TambahKecepatan()
}