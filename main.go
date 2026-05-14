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
var jumlahKandidat int = 0

func main() {
	var aksi int
	var nama, status string
	var kandidatList tabKandidat
	var keluar bool = false

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&nama)
	fmt.Print("Masukkan Status Keanggotaan : ")
	fmt.Scan(&status)

	for !keluar {
		if status == "panitia" {
			fmt.Println("Pilih aksi yang ingin dilakukan")
			fmt.Println("0 Keluar\n1 Tambah Kandidat\n2 Update Kandidat\n3 Hapus Kandidat\n4 Voting\n5 Tabel Kandidat")
			fmt.Print("Masukkan Aksi : ")
			fmt.Scan(&aksi)
			if aksi == 0 {
				keluar = true
				fmt.Println("Terima kasih, sampai jumpa!")
			}else if aksi == 1 {
				tambah(&kandidatList)
			}else if aksi == 2 {
				update(&kandidatList)
			}else if aksi == 3 {
				hapus(&kandidatList)
			}else if aksi == 4 {
				voting(&kandidatList)
			}else if aksi == 5 {
				tabel(&kandidatList)
			}else{
				fmt.Println("Mohon maaf aksi tidak ditemukan")
			}
		}else if status == "anggota" {
			fmt.Println("Pilih aksi yang ingin dilakukan")
			fmt.Println("0 Keluar\n1 Voting\n2 Tabel Kandidat")
			fmt.Print("Masukkan Aksi : ")
			fmt.Scan(&aksi)
			if aksi == 0 {
				keluar = true
				fmt.Println("Terima kasih, sampai jumpa!")
			}else if aksi == 1 {
				voting(&kandidatList)
			}else if aksi == 2 {
				tabel(&kandidatList)
			}else{
				fmt.Println("Mohon maaf aksi tidak ditemukan")
			}
		}else{
			fmt.Println("Status tidak dikenali")
			keluar = true
		}
	}
}

func cariSlotKosong(kandidatList *tabKandidat) int {
	for i := 0; i < NMAX; i++ {
		if kandidatList[i].nomorUrut == 0 {
			return i
		}
	}
	return -1
}

func tambah(kandidatList *tabKandidat) {
	var slot int
	fmt.Println("Tambah Kandidat")
	slot = cariSlotKosong(kandidatList)
	if slot == -1 {
		fmt.Println("Maaf, array penuh!")
		return
	}
	
	fmt.Print("Masukkan Nomor Urut : ")
	fmt.Scan(&kandidatList[slot].nomorUrut)
	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&kandidatList[slot].nama)
	fmt.Print("Masukkan Visi : ")
	fmt.Scan(&kandidatList[slot].visi)
	fmt.Print("Masukkan Misi : ")
	fmt.Scan(&kandidatList[slot].misi)
	kandidatList[slot].vote = 0
	jumlahKandidat++
	fmt.Printf("Kandidat nomor urut %d berhasil ditambahkan!\n", kandidatList[slot].nomorUrut)
	return
}

func update(kandidatList *tabKandidat) {
	fmt.Println("Update Kandidat")
}

func hapus(kandidatList *tabKandidat) {
	var nomorUrut, i int
	fmt.Println("Hapus Kandidat")
	fmt.Print("Masukkan Nomor Urut Kandidat yang akan dihapus : ")
	fmt.Scan(&nomorUrut)
	
	for i = 0; i < NMAX; i++ {
		if kandidatList[i].nomorUrut == nomorUrut {
			kandidatList[i].nomorUrut = 0
			kandidatList[i].nama = ""
			kandidatList[i].visi = ""
			kandidatList[i].misi = ""
			kandidatList[i].vote = 0
			jumlahKandidat--
			fmt.Printf("Kandidat nomor urut %d berhasil dihapus!\n", nomorUrut)
			return
		}
	}
	fmt.Println("Kandidat tidak ditemukan!")
}

func voting(kandidatList *tabKandidat) {
	fmt.Println("Voting")
}

func tabel(kandidatList *tabKandidat) {
	fmt.Println("Tabel Kandidat")
}

// User panitia dan anggota input biasa, atau mau buat tambah data pengguna? tapi klo ada data pengguna terus log out, data votingnya ilang dong?
// dalam tambah kandidat : input nomor urut, visi misi
// didalam tabel kandidat : ada search data berdasarkan nomor urut, mengurutkan data berdasarkan suara terbanyak atau nomor urut, menampikan statistik presentase masing masing kandidat, total pemilihan suara yang sudah masuk