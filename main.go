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
	isVoted  bool
}
type tabKandidat [NMAX]kandidat
type tabAnggota [NMAX]anggota

func isiAnggota(anggotaList *tabAnggota, jumlahAnggota *int) {
	anggotaList[0] = anggota{"reksa", "reksa1", "panitia", false}
	anggotaList[1] = anggota{"amel", "amel1", "panitia", false}
	anggotaList[2] = anggota{"aldi", "aldi1", "anggota", false}
	anggotaList[3] = anggota{"rafa", "rafa1", "anggota", false}
	*jumlahAnggota = 4
}

func main() {
	var aksi, jumlahAnggota, jumlahKandidat int
	var nama, password, status string
	var kandidatList tabKandidat
	var anggotaList tabAnggota
	var user anggota
	var programKeluar, menuKeluar bool = false, false

	programKeluar = false
	menuKeluar = false
	isiAnggota(&anggotaList, &jumlahAnggota)
	for !programKeluar {
		fmt.Println("\nSelamat datang di Sistem E-Voting!")
		fmt.Print("Masukkan Nama : ")
		fmt.Scan(&nama)

		if nama == "quit" {
			programKeluar = true
		} else {
			fmt.Print("Masukkan Password : ")
			fmt.Scan(&password)

			status = cekLogin(&anggotaList, nama, password, jumlahAnggota)
			if status == "" {
				fmt.Println("Login gagal: nama atau password tidak cocok")
			} else {
				fmt.Print("\nSelamat datang, ", nama)
				for !menuKeluar {
					if status == "panitia" {
						fmt.Println("\nPilih aksi yang ingin dilakukan")
						fmt.Println("1 CRUD Data Kandidat\n2 Voting\n3 Tabel Kandidat\n4 Keluar")
						fmt.Print("Masukkan Aksi : ")
						fmt.Scan(&aksi)
						if aksi == 1 {
							crud(&kandidatList, &jumlahKandidat)
						} else if aksi == 2 {
							voting(&kandidatList, jumlahKandidat, &user)
						} else if aksi == 3 {
							tabel(&kandidatList, jumlahKandidat)
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
							voting(&kandidatList, jumlahKandidat, &user)
						} else if aksi == 2 {
							tabel(&kandidatList, jumlahKandidat)
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
}

func cekLogin(anggotaList *tabAnggota, nama, password string, jumlahAnggota int) string {
	var i int
	for i = 0; i < jumlahAnggota; i++ {
		if anggotaList[i].nama == nama && anggotaList[i].password == password {
			return anggotaList[i].status
		}
	}
	return ""
}

func crud(kandidatList *tabKandidat, jumlahKandidat *int) {
	var aksi int
	fmt.Println("\nCRUD Data Kandidat")
	fmt.Println("1 Tambah Kandidat\n2 Update Kandidat\n3 Hapus Kandidat\n4 Kembali")
	fmt.Print("Masukkan Aksi : ")
	fmt.Scan(&aksi)
	if aksi == 1 {
		tambah(kandidatList, jumlahKandidat)
	} else if aksi == 2 {
		update(kandidatList, *jumlahKandidat)
	} else if aksi == 3 {
		hapus(kandidatList, jumlahKandidat)
	} else if aksi == 4 {
		return
	} else {
		fmt.Println("Aksi tidak valid!")
	}
}

func tambah(kandidatList *tabKandidat, jumlahKandidat *int) {
	var slot, i int
	var isFound bool
	var nomorUrut int

	fmt.Println("\nTambah Kandidat")
	if *jumlahKandidat >= NMAX {
		fmt.Println("Kapasitas kandidat penuh, tidak bisa menambah lagi.")
		return
	}
	fmt.Print("Masukkan Nomor Urut : ")
	fmt.Scan(&nomorUrut)

	if nomorUrut == 0 {
		fmt.Println("Nomor urut tidak boleh 0.")
		return
	}
	for i = 0; i < *jumlahKandidat; i++ {
		if kandidatList[i].nomorUrut == nomorUrut && isFound == false {
			isFound = true
		}
	}

	if isFound {
		fmt.Println("Nomor urut sudah digunakan, silakan coba lagi dengan nomor urut yang berbeda.")
		return
	}

	slot = 0
	for slot < NMAX && kandidatList[slot].nomorUrut != 0 {
		slot++
	}

	kandidatList[slot].nomorUrut = nomorUrut
	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&kandidatList[slot].nama)
	fmt.Print("Masukkan Visi : ")
	fmt.Scan(&kandidatList[slot].visi)
	fmt.Print("Masukkan Misi : ")
	fmt.Scan(&kandidatList[slot].misi)
	kandidatList[slot].vote = 0
	(*jumlahKandidat)++
	fmt.Printf("Kandidat nomor urut %d berhasil ditambahkan!\n", kandidatList[slot].nomorUrut)
	selectionSortByNomorUrut(kandidatList, *jumlahKandidat)
	return
}

func update(kandidatList *tabKandidat, jumlahKandidat int) {
	var nomor, i int
	fmt.Println("\nUpdate Kandidat")
	fmt.Print("Masukkan nomor urut kandidat yang ingin diupdate : ")
	fmt.Scan(&nomor)
	for i = 0; i < jumlahKandidat; i++ {
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

func hapus(kandidatList *tabKandidat, jumlahKandidat *int) {
	var nomorUrut, i, j int
	fmt.Println("\nHapus Kandidat")
	fmt.Print("Masukkan Nomor Urut Kandidat yang akan dihapus : ")
	fmt.Scan(&nomorUrut)
	for i = 0; i < *jumlahKandidat; i++ {
		if kandidatList[i].nomorUrut == nomorUrut {
			for j = i; j < *jumlahKandidat-1; j++ {
				kandidatList[j] = kandidatList[j+1]
			}

			(*kandidatList)[*jumlahKandidat-1] = kandidat{}

			*jumlahKandidat--
			fmt.Printf("Kandidat nomor urut %d berhasil dihapus!\n", nomorUrut)
			return
		}
	}
	fmt.Println("Kandidat tidak ditemukan!")
}

func seqSearch(kandidatList *tabKandidat, nomorUrut int, jumlahKandidat int) int {
	var i int
	for i = 0; i < jumlahKandidat; i++ {
		if kandidatList[i].nomorUrut == nomorUrut {
			kandidatList[i].vote++
			return i
		}
	}
	return -1
}

func voting(kandidatList *tabKandidat, jumlahKandidat int, u *anggota) {
	var nomorUrut, index, i int
	fmt.Println("\nVoting")
	fmt.Printf("%-5s | %-20s\n", "Nomor", "Nama")
	for i = 0; i < jumlahKandidat; i++ {
		if kandidatList[i].nomorUrut != 0 {
			fmt.Printf("%-5d | %-20s\n", kandidatList[i].nomorUrut, kandidatList[i].nama)
		}
	}

	if !u.isVoted {
		fmt.Print("Masukkan Nomor Urut Yang Ingin di Vote : ")
		fmt.Scan(&nomorUrut)
		index = seqSearch(kandidatList, nomorUrut, jumlahKandidat)

		if index != -1 {
			kandidatList[index].vote++
			u.isVoted = true

			fmt.Printf("Terima kasih telah memilih kandidat nomor urut %d!\n", nomorUrut)
		} else {
			fmt.Println("Kandidat tidak ditemukan!")
		}
	} else {
		fmt.Println("Anda telah melakukan voting!")
	}
}

func tabel(kandidatList *tabKandidat, jumlahKandidat int) {
	var i, aksi, nomorUrut, index int
	fmt.Println("\nTabel Kandidat")
	fmt.Printf("%-5s | %-20s | %-40s | %-40s | %-5s\n", "Nomor", "Nama", "Visi", "Misi", "Jumlah Vote")
	for i = 0; i < jumlahKandidat; i++ {
		if kandidatList[i].nomorUrut != -1 {
			fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[i].nomorUrut, kandidatList[i].nama, kandidatList[i].visi, kandidatList[i].misi, kandidatList[i].vote)
		}
	}
	fmt.Println("\nPilih aksi yang ingin dilakukan")
	fmt.Println("1 Urutkan berdasarkan jumlah vote\n2 Urutkan berdasarkan nomor urut\n3 Cari kandidat berdasarkan nomor urut\n4 Tampilkan Statistik\n5 Kembali")
	fmt.Print("Masukkan Aksi : ")
	fmt.Scan(&aksi)
	if aksi == 1 {
		insertionSortByVote(kandidatList, jumlahKandidat)
		fmt.Println("Tabel Kandidat (diurutkan berdasarkan jumlah vote)")
		fmt.Printf("%-5s | %-20s | %-40s | %-40s | %-5s\n", "Nomor", "Nama", "Visi", "Misi", "Jumlah Vote")
		for i = 0; i < jumlahKandidat; i++ {
			if kandidatList[i].nomorUrut != 0 {
				fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[i].nomorUrut, kandidatList[i].nama, kandidatList[i].visi, kandidatList[i].misi, kandidatList[i].vote)
			}
		}
	} else if aksi == 2 {
		selectionSortByNomorUrut(kandidatList, jumlahKandidat)
		fmt.Println("Tabel Kandidat (diurutkan berdasarkan nomor urut)")
		fmt.Printf("%-5s | %-20s | %-40s | %-40s | %-5s\n", "Nomor", "Nama", "Visi", "Misi", "Jumlah Vote")
		for i = 0; i < jumlahKandidat; i++ {
			if kandidatList[i].nomorUrut != 0 {
				fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[i].nomorUrut, kandidatList[i].nama, kandidatList[i].visi, kandidatList[i].misi, kandidatList[i].vote)
			}
		}
	} else if aksi == 3 {
		fmt.Print("Masukkan Nomor Urut Kandidat yang ingin dicari : ")
		fmt.Scan(&nomorUrut)
		selectionSortByNomorUrut(kandidatList, jumlahKandidat)
		index = binarySearch(kandidatList, nomorUrut, jumlahKandidat)
		if index != -1 {
			fmt.Printf("Kandidat ditemukan!\n")
			fmt.Printf("%-5d | %-20s | %-40s | %-40s | %-5d\n", kandidatList[index].nomorUrut, kandidatList[index].nama, kandidatList[index].visi, kandidatList[index].misi, kandidatList[index].vote)
		} else {
			fmt.Println("Kandidat tidak ditemukan!")
		}
	} else if aksi == 4 {
		tampilkanStatistik(kandidatList, jumlahKandidat)
	} else if aksi == 5 {
		return
	} else {
		fmt.Println("Aksi tidak valid!")
	}
}

func insertionSortByVote(kandidatList *tabKandidat, jumlahKandidat int) {
	var i, j int
	var key kandidat
	for i = 1; i < jumlahKandidat; i++ {
		key = kandidatList[i]
		j = i - 1
		for j >= 0 && kandidatList[j].vote < key.vote {
			kandidatList[j+1] = kandidatList[j]
			j--
		}
		kandidatList[j+1] = key
	}
}

func selectionSortByNomorUrut(kandidatList *tabKandidat, jumlahKandidat int) {
	var i, j, minIdx int
	var temp kandidat
	for i = 0; i < jumlahKandidat; i++ {
		minIdx = i
		for j = i + 1; j < jumlahKandidat; j++ {
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

func binarySearch(kandidatList *tabKandidat, nomorUrut, jumlahKandidat int) int {
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

func tampilkanStatistik(kandidatList *tabKandidat, jumlahKandidat int) {
	var totalVotes, i int
	var persentase float64
	fmt.Println("\nStatistik Voting")
	for i = 0; i < jumlahKandidat; i++ {
		totalVotes = totalVotes + kandidatList[i].vote
	}
	fmt.Printf("Total Votes: %d\n", totalVotes)
	if totalVotes > 0 {
		fmt.Printf("Persentase Votes:\n")
		for i = 0; i < jumlahKandidat; i++ {
			if kandidatList[i].nomorUrut != 0 {
				persentase = float64(kandidatList[i].vote) / float64(totalVotes) * 100
				fmt.Printf("- %s: %.2f%%\n", kandidatList[i].nama, persentase)
			}
		}
	}
}
