package main
import "fmt"

type barang struct {
	jumlah int
	nama, kode string
	nilaiSatuan, nilaiTotal float64
}

type transaksi struct {
	kodeT int
	keterangan string
}
type tabBarang [10000]barang
type tabTransaksi [10000]transaksi
type gudang struct {
	masterBarang tabBarang
	totalNilaiAset float64
	jumlahDataBarang, jumlahTransaksi int
	riwayatTransaksi tabTransaksi
}

func main() {
	var inventaris gudang
	menu(&inventaris)
}

func menu(inventaris *gudang) {
	var opsi, find int
	var cari string
	fmt.Println("1. Masukkan Data")
	fmt.Println("2. Tampilkan Data")
	fmt.Println("3. Cari Data")
	fmt.Println("4. Hapus Data")
	fmt.Println("5. Ubah Data")
	fmt.Println("6. Urutkan Data (berdasarkan jumlah stok)")
	fmt.Println("7. Transaksi masuk")
	fmt.Println("8. Transaksi keluar")
	fmt.Println("9. Keluar")
	fmt.Print("Masukkan pilihan tindakan: ")
	fmt.Scan(&opsi)
	fmt.Println()
	switch opsi {
		case 1 :
		inputBarang(inventaris, &inventaris.jumlahDataBarang, &inventaris.jumlahTransaksi)
		case 2 : 
		tampilkanData(*inventaris)
		case 3:
		fmt.Print("Masukkan kode atau nama barang yang akan dicari: ")
		fmt.Scan(&cari)
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search (Hanya jika menggunakan nama barang)")
		fmt.Print("Pilih metode pencarian: ")
		fmt.Scan(&opsi)
		switch opsi {
			case 1:
				find = cariBarang(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
				if find != -1 {
					fmt.Println("Data ditemukan:")
					fmt.Printf("%-5s %-10s %-2d %-8g %-8g\n", inventaris.masterBarang[find].kode, inventaris.masterBarang[find].nama, inventaris.masterBarang[find].jumlah, inventaris.masterBarang[find].nilaiSatuan, inventaris.masterBarang[find].nilaiTotal)
				} else {
					fmt.Println("Data tidak terdapat dalam inventaris")
				}
			case 2:
				cariBinary(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
		}
		case 4 :
		fmt.Print("Masukkan kode atau nama barang yang akan dihapus: ")
		fmt.Scan(&cari)
		find = cariBarang(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
		if find != -1 {
			hapusData(&inventaris.masterBarang, find, &inventaris.jumlahDataBarang)
		} else {
			fmt.Println("Data sudah tidak ada dalam inventaris")
		}
		case 5:
		fmt.Print("Masukkan kode atau nama barang yang akan diubah: ")
		fmt.Scan(&cari)
		find = cariBarang(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
		if find != -1 {
			ubahData(inventaris, find)
		} else {
			fmt.Println("Data tidak terdapat dalam inventaris")
		}
		case 6:
		fmt.Println("1. Menaik (Ascending)")
		fmt.Println("2. Menurun (Descending)")
		fmt.Print("Pilih pengurutan yang diinginkan: ")
		fmt.Scan(&opsi)
		switch opsi {
			case 1:
				sortAsc(&inventaris.masterBarang, inventaris.jumlahDataBarang)
			fmt.Println("Data sukses diurutkan menaik berdasarkan jumlah stok!")
			case 2:
				sortDesc(&inventaris.masterBarang, inventaris.jumlahDataBarang)
			fmt.Println("Data sukses diurutkan menurun berdasarkan jumlah stok!")
		}
		case 7:
		fmt.Print("Masukkan kode atau nama barang yang masuk: ")
		fmt.Scan(&cari)
		find = cariBarang(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
		transaksiMasuk(inventaris, find)
		case 8:
		fmt.Print("Masukkan kode atau nama barang yang keluar: ")
		fmt.Scan(&cari)
		find = cariBarang(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
		transaksiKeluar(inventaris, find)
	}
	if opsi > 9 || opsi < 1 {
		fmt.Println("opsi tidak tersedia")
		fmt.Println()
		menu(inventaris)
	} else if opsi != 9 {
		fmt.Println()
		menu(inventaris)
	}
}

func inputBarang(A *gudang, n, m *int) {
	var i, j int
	fmt.Print("Masukkan jumlah barang yang ingin di input: ")
	fmt.Scan(&j)
	fmt.Println("Contoh format input data: SBK001 Minyak_Goreng 14000.0 5")
	fmt.Println("Silahkan masukkan data")
	for i = *n; i < (*n + j); i++ {
		fmt.Scan(&A.masterBarang[i].kode, &A.masterBarang[i].nama, &A.masterBarang[i].nilaiSatuan, &A.masterBarang[i].jumlah)
		A.masterBarang[i].nilaiTotal = A.masterBarang[i].nilaiSatuan * float64(A.masterBarang[i].jumlah)
		A.riwayatTransaksi[i].kodeT = i+1
		A.riwayatTransaksi[i].keterangan = "Barang " + A.masterBarang[i].nama + " masuk dalam jumlah " + fmt.Sprint(A.masterBarang[i].jumlah)
		A.totalNilaiAset = A.totalNilaiAset + A.masterBarang[i].nilaiTotal
	}
	*m = *m + j
	*n = *n + j
}


func hapusData(A *tabBarang, m int , n *int) {
	var i int
	for i = m; i < *n; i++ {
		A[i] = A[i + 1]
	}
	*n--
}

func tampilkanData(A gudang) {
	var i, opsi, minidx int
	fmt.Println("1. Daftar Barang")
	fmt.Println("2. Riwayat Transaksi")
	fmt.Print("Pilih data yang ingin ditampilkan (dalam angka): ")
	fmt.Scan(&opsi)
	switch opsi {
		case 1:
		minidx = 0
		fmt.Printf("%-5s %-10s %-2s %-8s %-8s\n", "Kode", "Nama", "Jml", "NilaiSat", "NilaiTot")
		for i = 0; i < A.jumlahDataBarang; i++ {
			if A.masterBarang[minidx].jumlah > A.masterBarang[i].jumlah {
				minidx = i
			}
			fmt.Printf("%-5s %-15s %-2d %-8g %-8g\n", A.masterBarang[i].kode, A.masterBarang[i].nama, A.masterBarang[i].jumlah, A.masterBarang[i].nilaiSatuan, A.masterBarang[i].nilaiTotal)
		}
		fmt.Printf("NIlai total aset adalah %.2f\n", hitungTotalAset(A.masterBarang, A.jumlahDataBarang))
		fmt.Printf("Barang dengan stok terendah adalah %s dengan jumlah %d\n", A.masterBarang[minidx].nama, A.masterBarang[minidx].jumlah)
		case 2:
		fmt.Printf("%-5s %-20s\n", "Kode", "Keterangan")
		for i = 0; i < A.jumlahTransaksi; i++ {
			fmt.Printf("%-5d %-20s\n", A.riwayatTransaksi[i].kodeT, A.riwayatTransaksi[i].keterangan)
		}
	}
}

func hitungTotalAset(A tabBarang, n int) float64 {
	var i int
	var total float64
	for i = 0; i < n; i++ {
		total = total + A[i].nilaiTotal
	}
	return total
}

func ubahData(A *gudang, n int) {
	var opsi int
	fmt.Println("1. Ubah Nama")
	fmt.Println("2. Ubah Nilai Satuan")
	fmt.Println("3. Ubah Jumlah")
	fmt.Print("Apa yang ingin diubah: ")
	fmt.Scan(&opsi)
	fmt.Println("Masukkan perubahan")
	switch opsi {
		case 1:
		fmt.Scan(&A.masterBarang[n].nama)
		case 2:
		fmt.Scan(&A.masterBarang[n].nilaiSatuan)
		A.totalNilaiAset = A.totalNilaiAset - A.masterBarang[n].nilaiTotal
		A.masterBarang[n].nilaiTotal = A.masterBarang[n].nilaiSatuan * float64(A.masterBarang[n].jumlah)
		A.totalNilaiAset = A.totalNilaiAset + A.masterBarang[n].nilaiTotal
		case 3:
		fmt.Scan(&A.masterBarang[n].jumlah)
		A.totalNilaiAset = A.totalNilaiAset - A.masterBarang[n].nilaiTotal
		A.masterBarang[n].nilaiTotal = A.masterBarang[n].nilaiSatuan * float64(A.masterBarang[n].jumlah)
		A.totalNilaiAset = A.totalNilaiAset + A.masterBarang[n].nilaiTotal
	}
}

func transaksiMasuk(A *gudang, n int) {
	var masuk int
	fmt.Print("Masukkan jumlah barang masuk: ")
	fmt.Scan(&masuk)
	A.masterBarang[n].jumlah = A.masterBarang[n].jumlah + masuk
	A.masterBarang[n].nilaiTotal = A.masterBarang[n].nilaiSatuan * float64(A.masterBarang[n].jumlah)
	A.riwayatTransaksi[A.jumlahTransaksi].kodeT = A.jumlahTransaksi + 1
	A.riwayatTransaksi[A.jumlahTransaksi].keterangan = "Barang " + A.masterBarang[n].nama + " bertambah sebanyak " + fmt.Sprint(masuk) + " dengan nilai " + fmt.Sprint(float64(masuk)*A.masterBarang[n].nilaiSatuan)
	A.totalNilaiAset = A.totalNilaiAset + (float64(masuk)*A.masterBarang[n].nilaiSatuan)
	A.jumlahTransaksi++
}
func transaksiKeluar(A *gudang, n int) {
	var keluar int
	fmt.Print("Masukkan jumlah barang keluar: ")
	fmt.Scan(&keluar)
	if A.masterBarang[n].jumlah - keluar >= 0 {
		A.masterBarang[n].jumlah = A.masterBarang[n].jumlah - keluar 
		A.masterBarang[n].nilaiTotal = A.masterBarang[n].nilaiSatuan * float64(A.masterBarang[n].jumlah)
		A.riwayatTransaksi[A.jumlahTransaksi].kodeT = A.jumlahTransaksi + 1
		A.riwayatTransaksi[A.jumlahTransaksi].keterangan = "Barang " + A.masterBarang[n].nama + " keluar sebanyak " + fmt.Sprint(keluar) + " dengan nilai " + fmt.Sprint(float64(keluar)*A.masterBarang[n].nilaiSatuan)
		A.totalNilaiAset = A.totalNilaiAset - (float64(keluar)*A.masterBarang[n].nilaiSatuan)
	} else {
		fmt.Println("Jumlah data keluar melebihi stok saat ini, anda yakin telah menginput jumlah yang benar?")
	}
	A.jumlahTransaksi++
}

func sortAsc(A *tabBarang, n int) {
	var i, j, minidx int
	for i = 0; i < n-1; i++ {
		minidx = i 
		for j = i + 1; j < n; j++ {
			if A[j].jumlah < A[minidx].jumlah {
				minidx = j 
			}
		}
		A[i], A[minidx] = A[minidx], A[i]
	}
}

func sortDesc(A *tabBarang, n int) {
	var i, j int
	var key barang
	for i = 1; i < n; i++ {
		key = A[i]
		j = i
		for j > 0 && A[j-1].jumlah < key.jumlah {
			A[j] = A[j-1]
			j--
		}
		A[j] = key
	}
}



func cariBarang(A tabBarang, n int, cari string) int {
	var i int
	for i = 0; i < n; i++ {
		if A[i].nama == cari || A[i].kode == cari {
			return i
		}
	}
	return -1
}
 func cariBinary(A tabBarang, n int, cari string) {
	var left, right, mid, found int
	sortAscNama(&A, n,)
	found = -1
	left = 0
	right =	n-1
	mid = (left + right) / 2
	for left <= right && found == -1 {
		if A[mid].nama < cari {
			left = mid +1 
		} else if A[mid].nama > cari {
			right = mid -1
		} else { 
			found = mid
		}
		mid = (left + right) / 2
	}
	if found != -1 {
		fmt.Println("Data ditemukan:")
		fmt.Printf("%-5s %-10s %-2d %-8g %-8g\n", A[found].kode, A[found].nama, A[found].jumlah, A[found].nilaiSatuan, A[found].nilaiTotal)
	} else {
		fmt.Println("Data tidak terdapat dalam inventaris")
	}
 }
func sortAscNama(A *tabBarang, n int) {
	var i, j, minidx int
	for i = 0; i < n-1; i++ {
		minidx = i 
		for j = i + 1; j < n; j++ {
			if A[j].nama < A[minidx].nama {
				minidx = j 
			}
		}
		A[i], A[minidx] = A[minidx], A[i]
	}
}

func output(A tabBarang, n int) {
	var i int
	for i = 0; i < n; i++ {
		fmt.Println(A[i].nama)
	}
}