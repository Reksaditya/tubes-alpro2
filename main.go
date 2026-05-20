package main

import "fmt"

const NMAX int = 100

type kandidat struct {
	nomorUrut int
	nama      string
	visi      string
	misi      string
	vote      int
}
type anggota struct {
	nama     string
	password string
	status   string
}
type tabKandidat [NMAX]kandidat
type tabAnggota [NMAX]anggota

var jumlahKandidat int = 0
var jumlahAnggota int = 0

func isiAnggota(anggotaList *tabAnggota) {
	anggotaList[0] = anggota{"reksa", "reksa1", "panitia"}
	anggotaList[1] = anggota{"amel", "amel1", "anggota"}
	anggotaList[2] = anggota{"aldi", "aldi1", "anggota"}
	jumlahAnggota = 3
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
					} else if aksi == 2 {
						voting(&kandidatList)
					} else if aksi == 3 {
						tabel(&kandidatList)
					} else if aksi == 4 {
						menuKeluar = true
						fmt.Println("Terima kasih, sampai jumpa!")
					} else {
						fmt.Println("Mohon maaf aksi tidak ditemukan")
					}
				} else if status == "anggota" {
					fmt.Println("\nPilih aksi yang ingin dilakukan")
					fmt.Println("1 Voting\n2 Tabel Kandidat\n3 Keluar")
					fmt.Print("Masukkan Aksi : ")
					fmt.Scan(&aksi)
					if aksi == 1 {
						voting(&kandidatList)
					} else if aksi == 2 {
						tabel(&kandidatList)
					} else if aksi == 3 {
						menuKeluar = true
						fmt.Println("Terima kasih, sampai jumpa!")
					} else {
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
	for i = 0; i < jumlahAnggota; i++ {
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
	var isFound bool

	fmt.Println("Tambah Kandidat")
	slot = cariSlotKosong(kandidatList)
	
	fmt.Print("Masukkan Nomor Urut : ")
	fmt.Scan(&kandidatList[slot].nomorUrut)
	
	isFound = false
	i := 0
	for i < jumlahKandidat && !isFound {
		if kandidatList[i].nomorUrut == kandidatList[slot].nomorUrut {
			isFound = true
		}	
		i++
	}

	if isFound {
		fmt.Println("Nomor urut sudah digunakan, silakan coba lagi dengan nomor urut yang berbeda.")
		return
	}
	
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
	for i = 0; i < len(kandidatList); i++ {
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
	for i = 0; i < len(kandidatList); i++ {
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

func seqSearch(kandidatList *tabKandidat, nomorUrut int) int {
	for i := 0; i < len(kandidatList); i++ {
		if kandidatList[i].nomorUrut == nomorUrut {
			kandidatList[i].vote++
			return i
		}
	}
	return -1
}

func voting(kandidatList *tabKandidat) {
	var nomorUrut, index int
	fmt.Println("Voting")
	fmt.Print("Masukkan Nomor Urut Yang Ingin di Vote : ")
	fmt.Scan(&nomorUrut)
	index = seqSearch(kandidatList, nomorUrut)

	if index != -1 {
		fmt.Printf("Terima kasih telah memilih kandidat nomor urut %d!\n", nomorUrut)
	} else {
		fmt.Println("Kandidat tidak ditemukan!")
	}
}

func tabel(kandidatList *tabKandidat) {
	var i, aksi, nomorUrut, index int
	fmt.Println("Tabel Kandidat")
	fmt.Printf("%-5s | %-20s | %-40s | %-40s | %-5s\n", "Nomor", "Nama", "Visi", "Misi", "Jumlah Vote")
	for i = 0; i < len(kandidatList); i++ {
		if kandidatList[i].nomorUrut != 0 {
			fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[i].nomorUrut, kandidatList[i].nama, kandidatList[i].visi, kandidatList[i].misi, kandidatList[i].vote)
		}
	}
	fmt.Println("Pilih aksi yang ingin dilakukan")
	fmt.Println("1 Urutkan berdasarkan jumlah vote\n2 Urutkan berdasarkan nomor urut\n3 Cari kandidat berdasarkan nomor urut\n4 Kembali")
	fmt.Print("Masukkan Aksi : ")
	fmt.Scan(&aksi)
	if aksi == 1 {
		insertionSortByVote(kandidatList, jumlahKandidat)
		fmt.Println("Tabel Kandidat (diurutkan berdasarkan jumlah vote)")
		fmt.Printf("%-5s | %-20s | %-40s | %-40s | %-5s\n", "Nomor", "Nama", "Visi", "Misi", "Jumlah Vote")
		for i = 0; i < len(kandidatList); i++ {
			if kandidatList[i].nomorUrut != 0 {
				fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[i].nomorUrut, kandidatList[i].nama, kandidatList[i].visi, kandidatList[i].misi, kandidatList[i].vote)
			}
		}
	} else if aksi == 2 {
		selectionSortByNomorUrut(kandidatList, jumlahKandidat)
		fmt.Println("Tabel Kandidat (diurutkan berdasarkan nomor urut)")
		fmt.Printf("%-5s | %-20s | %-40s | %-40s | %-5s\n", "Nomor", "Nama", "Visi", "Misi", "Jumlah Vote")
		for i = 0; i < len(kandidatList); i++ {
			if kandidatList[i].nomorUrut != 0 {
				fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[i].nomorUrut, kandidatList[i].nama, kandidatList[i].visi, kandidatList[i].misi, kandidatList[i].vote)
			}
		}
	} else if aksi == 3 {
		fmt.Print("Masukkan Nomor Urut Kandidat yang ingin dicari : ")
		fmt.Scan(&nomorUrut)
		selectionSortByNomorUrut(kandidatList, jumlahKandidat)
		index = binarySearch(kandidatList, nomorUrut)
		if index != -1 {
			fmt.Printf("Kandidat ditemukan!\n")
			fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[index].nomorUrut, kandidatList[index].nama, kandidatList[index].visi, kandidatList[index].misi, kandidatList[index].vote)
		} else {
			fmt.Println("Kandidat tidak ditemukan!")
		}
	} else if aksi == 4 {
		return
	} else {
		fmt.Println("Aksi tidak valid!")
	}
}

func insertionSortByVote(kandidatList *tabKandidat, n int) {
	var i, j int
	var key kandidat
	for i = 1; i < n; i++ {
		key = kandidatList[i]
		j = i - 1
		for j >= 0 && kandidatList[j].vote < key.vote {
			kandidatList[j+1] = kandidatList[j]
			j--
		}
		kandidatList[j+1] = key
	}
}

func selectionSortByNomorUrut(kandidatList *tabKandidat, n int) {
	var i, j, minIdx int
	var temp kandidat
	for i = 0; i < n; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if kandidatList[j].nomorUrut < kandidatList[minIdx].nomorUrut && kandidatList[j].nomorUrut != 0 {
				minIdx = j
			}
		}
		if minIdx != i {
			temp = kandidatList[i]
			kandidatList[i] = kandidatList[minIdx]
			kandidatList[minIdx] = temp
		}
	}
}

func binarySearch(kandidatList *tabKandidat, nomorUrut int) int {
	var left, right, mid int
	left = 0
	right = jumlahKandidat - 1

	for left <= right {
		mid = (left + right) / 2
		if kandidatList[mid].nomorUrut == nomorUrut {
			return mid
		} else if kandidatList[mid].nomorUrut < nomorUrut {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}