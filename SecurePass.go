package main

import (
	"fmt"
)

// Struktur data untuk Akun
type Account struct {
	ServiceName string
	Email       string
	Password    string
	LastUpdate  string
}

// Global variable untuk menyimpan daftar akun
var accounts []Account

func main() {
	var choice int
	for {
		fmt.Println("\n--- SecurePass Menu ---")
		fmt.Println("1. Tambah Akun")
		fmt.Println("2. Edit Akun")
		fmt.Println("3. Hapus Akun")
		fmt.Println("4. Cari Akun (Sequential/Binary)")
		fmt.Println("5. Urutkan Akun (Selection/Insertion)")
		fmt.Println("6. Tampilkan Data & Statistik")
		fmt.Println("7. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addAccount()
		case 2:
			editAccount()
		case 3:
			deleteAccount()
		case 4:
			searchMenu()
		case 5:
			sortMenu()
		case 6:
			displayAccounts()
		case 7:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

// Fitur CRUD (Create, Read, Update, Delete)
func addAccount() {
	var acc Account
	fmt.Print("Nama Layanan: ")
	fmt.Scan(&acc.ServiceName)
	fmt.Print("Email: ")
	fmt.Scan(&acc.Email)
	fmt.Print("Password: ")
	fmt.Scan(&acc.Password)
	fmt.Print("Tanggal (DD-MM-YYYY): ")
	fmt.Scan(&acc.LastUpdate)
	accounts = append(accounts, acc)
	fmt.Println("Akun berhasil ditambahkan!")
}

func editAccount() {
	var name string
	fmt.Print("Masukkan nama layanan yang ingin diedit: ")
	fmt.Scan(&name)
	for i := range accounts {
		if accounts[i].ServiceName == name {
			fmt.Print("Email baru: ")
			fmt.Scan(&accounts[i].Email)
			fmt.Print("Password baru: ")
			fmt.Scan(&accounts[i].Password)
			fmt.Print("Tanggal baru: ")
			fmt.Scan(&accounts[i].LastUpdate)
			fmt.Println("Data berhasil diubah!")
			return
		}
	}
	fmt.Println("Akun tidak ditemukan.")
}

func deleteAccount() {
	var name string
	fmt.Print("Masukkan nama layanan yang ingin dihapus: ")
	fmt.Scan(&name)
	for i := range accounts {
		if accounts[i].ServiceName == name {
			accounts = append(accounts[:i], accounts[i+1:]...)
			fmt.Println("Akun berhasil dihapus!")
			return
		}
	}
	fmt.Println("Akun tidak ditemukan.")
}

// Fitur Pencarian berdasarkan Nama Layanan
func searchMenu() {
	var choice int
	fmt.Println("1. Sequential Search | 2. Binary Search")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&choice)

	var target string
	fmt.Print("Cari nama layanan: ")
	fmt.Scan(&target)

	found := false
	if choice == 1 {
		for _, acc := range accounts {
			if acc.ServiceName == target {
				statusKekuatan := hitungKekuatanKataSandi(acc.Password)
				fmt.Printf("Ditemukan -> Layanan: %s | Email: %s | Status Sandi: %s\n", acc.ServiceName, acc.Email, statusKekuatan)
				found = true
				break
			}
		}
	} else if choice == 2 {
		// Catatan: Binary Search membutuhkan data yang sudah terurut
		Sorted := true
		for i := 0; i < len(accounts)-1; i++ {
			if accounts[i].ServiceName > accounts[i+1].ServiceName {
				Sorted = false
				break
		}
	}
	if !Sorted {
		fmt.Println("Error: Data belum diurutkan berdasarkan Nama Layanan. Binary Search tidak dapat dilakukan.")
		return
	}

		low, high := 0, len(accounts)-1
		for low <= high {
			mid := (low + high) / 2
			if accounts[mid].ServiceName == target {
				statusKekuatan := hitungKekuatanKataSandi(accounts[mid].Password)
				fmt.Printf("Ditemukan -> Layanan: %s | Email: %s | Status Sandi: %s\n", accounts[mid].ServiceName, accounts[mid].Email, statusKekuatan)
				found = true
				break
			} else if accounts[mid].ServiceName < target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	if !found {
		fmt.Println("Akun tidak ditemukan.")
	}
}

// Fitur Pengurutan berdasarkan Nama Layanan
func sortMenu() {
	var choice int
	fmt.Println("1. Selection Sort | 2. Insertion Sort")
	fmt.Print("Pilih metode: ")
	fmt.Scan(&choice)

	n := len(accounts)
	if choice == 1 {
		for i := 0; i < n-1; i++ {
			minIdx := i
			for j := i + 1; j < n; j++ {
				if accounts[j].ServiceName < accounts[minIdx].ServiceName {
					minIdx = j
				}
			}
			accounts[i], accounts[minIdx] = accounts[minIdx], accounts[i]
		}
	} else if choice == 2 {
		for i := 1; i < n; i++ {
			key := accounts[i]
			j := i - 1
			for j >= 0 && accounts[j].ServiceName > key.ServiceName {
				accounts[j+1] = accounts[j]
				j--
			}
			accounts[j+1] = key
		}
	}
	fmt.Println("Data berhasil diurutkan berdasarkan Nama Layanan secara alfabetis.")
}

// Fungsi untuk menghitung kekuatan kata sandi yang Anda berikan
func hitungKekuatanKataSandi(kataSandi string) string {
	panjang := len(kataSandi)
	hurufBesar := 0
	hurufKecil := 0
	angka := 0
	simbol := 0

	for _, c := range kataSandi {
		switch {
		case c >= 'A' && c <= 'Z':
			hurufBesar++
		case c >= 'a' && c <= 'z':
			hurufKecil++
		case c >= '0' && c <= '9':
			angka++
		default:
			simbol++
		}
	}

	skor := 0
	if panjang >= 8 {
		skor++
	}
	if panjang >= 12 {
		skor++
	}
	if hurufBesar > 0 {
		skor++
	}
	if hurufKecil > 0 {
		skor++
	}
	if angka > 0 {
		skor++
	}
	if simbol > 0 {
		skor++
	}

	switch {
	case skor <= 2:
		return "Lemah"
	case skor <= 4:
		return "Sedang"
	default:
		return "Kuat"
	}
}

// Menampilkan Data dan Statistik Klasifikasi Kekuatan Kata Sandi
func displayAccounts() {
	fmt.Println("\n--- Daftar Akun & Kekuatan Kata Sandi ---")
	
	lemahCount := 0
	sedangCount := 0
	kuatCount := 0

	for _, acc := range accounts {
		statusKekuatan := hitungKekuatanKataSandi(acc.Password)
		
		fmt.Printf("Layanan: %s | Email: %s | Tgl Update: %s | Kekuatan Sandi: [%s]\n", 
			acc.ServiceName, acc.Email, acc.LastUpdate, statusKekuatan)
		
		// Klasifikasi statistik berdasarkan fungsi yang ditambahkan
		switch statusKekuatan {
		case "Lemah":
			lemahCount++
		case "Sedang":
			sedangCount++
		case "Kuat":
			kuatCount++
		}
	}

	fmt.Printf("\n=== Statistik Total Akun: %d ===\n", len(accounts))
	fmt.Printf("Jumlah Sandi Lemah  : %d\n", lemahCount)
	fmt.Printf("Jumlah Sandi Sedang : %d\n", sedangCount)
	fmt.Printf("Jumlah Sandi Kuat   : %d\n", kuatCount)
}