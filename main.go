package main
import "fmt"

const NMAX int = 100
type kandidat struct{
	nomorUrut int
	nama string
	visi string
	misi string
	vote int
}
type tabKandidat [NMAX] kandidat

func main() {
	var aksi int
	var nama, status string

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Status Keanggotaan : ")
	fmt.Scan(&status)

	if status == "panitia" {
		fmt.Println("Pilih aksi yang ingin dilakukan")
		fmt.Println("1 Tambah Kandidat\n2 Update Kandidat\n3 Hapus Kandidat\n4 Voting\n5 Tabel Kandidat")
		fmt.Scan(&aksi)
		if aksi == 1 {
			tambah()
		}else if aksi == 2 {
			update()
		}else if aksi == 3 {
			hapus()
		}else if aksi == 4 {
			voting()
		}else if aksi == 5 {
			tabel()
		}else{
			fmt.Println("Mohon maaf aksi tidak ditemukan")
		}
	}else if status == "anggota" {
		fmt.Println("Pilih aksi yang ingin dilakukan")
		fmt.Println("1 Voting\n2 Tabel Kandidat")
		fmt.Scan(&aksi)
		if aksi == 1 {
			voting()
		}else if aksi == 2 {
			tabel()
		}else{
			fmt.Println("Mohon maaf aksi tidak ditemukan")
		}
	}
}

func tambah() {
	fmt.Println("Tambah Kandidat")
}

func update() {
	fmt.Println("Update Kandidat")
}

func hapus() {
	fmt.Println("Hapus Kandidat")
}

func voting() {
	fmt.Println("Voting")
}

func tabel() {
	fmt.Println("Tabel Kandidat")
}

// User panitia dan anggota input biasa, atau mau buat tambah data pengguna?
// dalam tambah kandidat : input nomor urut, visi misi
// didalam tabel kandidat : ada search data berdasarkan nomor urut, mengurutkan data berdasarkan suara terbanyak atau nomor urut, menampikan statistik presentase masing masing kandidat, total pemilihan suara yang sudah masuk