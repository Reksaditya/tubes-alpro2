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
type anggota struct{
	nama string
	password string
	status string
}
type tabKandidat [NMAX] kandidat
type tabAnggota [NMAX] anggota
var jumlahKandidat int = 0

func isiAnggota(anggotaList *tabAnggota) {
	anggotaList[0] = anggota{"reksa", "reksa1", "panitia"}
	anggotaList[1] = anggota{"amel", "amel1", "anggota"}
	anggotaList[2] = anggota{"aldi", "aldi1", "anggota"}
}

func main() {
	var aksi int
	var nama, password, status string
	var kandidatList tabKandidat
	var anggotaList tabAnggota
	var programKeluar, menuKeluar bool = false, false

	isiAnggota(&anggotaList)
	for !programKeluar {
		fmt.Print("Masukkan Nama : ")
		fmt.Scan(&nama)
		fmt.Print("Masukkan Password : ")
		fmt.Scan(&password)

		status = cekLogin(&anggotaList, nama, password)
		if status == "" {
			fmt.Println("Login gagal: nama atau password tidak cocok")
		}else{
			fmt.Print("\nSelamat datang,", nama)
			for !menuKeluar {
				if status == "panitia" {
					fmt.Println("\nPilih aksi yang ingin dilakukan")
					fmt.Println("1 CRUD Data Kandidat\n2 Voting\n3 Tabel Kandidat\n4 Keluar")
					fmt.Print("Masukkan Aksi : ")
					fmt.Scan(&aksi)
					if aksi == 1 {
						crud(&kandidatList)
					}else if aksi == 2 {
						voting(&kandidatList)
					}else if aksi == 3 {
						tabel(&kandidatList)
					}else if aksi == 4 {
						menuKeluar = true
						fmt.Println("Terima kasih, sampai jumpa!")
					}else{
						fmt.Println("Mohon maaf aksi tidak ditemukan")
					}
				}else if status == "anggota" {
					fmt.Println("\nPilih aksi yang ingin dilakukan")
					fmt.Println("1 Voting\n2 Tabel Kandidat\n3 Keluar")
					fmt.Print("Masukkan Aksi : ")
					fmt.Scan(&aksi)
					if aksi == 1 {
						voting(&kandidatList)
					}else if aksi == 2 {
						tabel(&kandidatList)
					}else if aksi == 3 {
						menuKeluar = true
						fmt.Println("Terima kasih, sampai jumpa!")
					}else{
						fmt.Println("Mohon maaf aksi tidak ditemukan")
					}
				}
			}
			menuKeluar = false
		}
	}
}

func cekLogin(anggotaList *tabAnggota, nama, password string) string {
	var i int
	for i = 0; i < 3; i++ {
		if anggotaList[i].nama == nama && anggotaList[i].password == password {
			return anggotaList[i].status
		}
	}
	return ""
}

func cariSlotKosong(kandidatList *tabKandidat) int {
	var i int
	for i = 0; i < NMAX; i++ {
		if kandidatList[i].nomorUrut == 0 {
			return i
		}
	}
	return -1
}

func crud(kandidatList *tabKandidat) {
	var aksi int
	fmt.Println("\nCRUD Data Kandidat")
	fmt.Println("1 Tambah Kandidat\n2 Update Kandidat\n3 Hapus Kandidat\n4 Kembali")
	fmt.Print("Masukkan Aksi : ")
	fmt.Scan(&aksi)
	if aksi == 1 {
		tambah(kandidatList)
	} else if aksi == 2 {
		update(kandidatList)
	} else if aksi == 3 {
		hapus(kandidatList)
	} else if aksi == 4 {
		return
	} else {
		fmt.Println("Aksi tidak valid!")
	}
}

func tambah(kandidatList *tabKandidat) {
	var slot int
	fmt.Println("Tambah Kandidat")
	slot = cariSlotKosong(kandidatList)
	
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
	var nomor, i int
	fmt.Println("Update Kandidat")
	fmt.Print("Masukkan nomor urut kandidat yang ingin diupdate : ")
	fmt.Scan(&nomor)
	for i = 0; i < NMAX; i++ {
		if kandidatList[i].nomorUrut == nomor {
			fmt.Print("Masukkan Nama : ")
			fmt.Scan(&kandidatList[i].nama)
			fmt.Print("Masukkan Visi : ")
			fmt.Scan(&kandidatList[i].visi)
			fmt.Print("Masukkan Misi : ")
			fmt.Scan(&kandidatList[i].misi)
			fmt.Printf("Kandidat nomor urut %d berhasil diupdate!\n", kandidatList[i].nomorUrut)
			return
		}
	}
	fmt.Println("Kandidat tidak ditemukan!")
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
	var nomorUrut, index int
	fmt.Println("Voting")
	fmt.Print("Masukkan Nomor Urut Yang Ingin di Vote : ")
	fmt.Scan(&nomorUrut)
	index = search(kandidatList, nomorUrut)

	if index != -1 {
		fmt.Printf("Terima kasih telah memilih kandidat nomor urut %d!\n", nomorUrut)
	} else {
		fmt.Println("Kandidat tidak ditemukan!")
	}
}

func tabel(kandidatList *tabKandidat) {
	var i int
	fmt.Println("Tabel Kandidat")
	for i = 0; i < NMAX; i++ {
		if kandidatList[i].nomorUrut != 0 {
			fmt.Println("Nomor Urut | Nama | Visi | Misi | Jumlah Vote")
			fmt.Printf("%d | %s | %s | %s | %d\n\n", kandidatList[i].nomorUrut, kandidatList[i].nama, kandidatList[i].visi, kandidatList[i].misi, kandidatList[i].vote)
		}
	}
}