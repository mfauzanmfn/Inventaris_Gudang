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
	jumlahDataBarang, jumlahTransaksiIn, jumlahTransaksiOut int
	transaksiIn, transaksiOut tabTransaksi
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
	fmt.Println("6. Urutkan Data (berdasarkan jumlah)")
	fmt.Println("7. Keluar")
	fmt.Print("Masukkan pilihan tindakan: ")
	fmt.Scan(&opsi)
	fmt.Println()
	switch opsi {
		case 1 :
		inputBarang(&inventaris.masterBarang, &inventaris.transaksiIn, &inventaris.jumlahDataBarang, &inventaris.jumlahTransaksiIn)
		case 2 : 
		tampilkanData(*inventaris)
		case 3:
		fmt.Print("Masukkan kode atau nama barang yang akan dicari:")
		fmt.Scan(&cari)
		find = cariBarang(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
		if find != -1 {
			fmt.Println("Data ditemukan:")
			fmt.Printf("%-5s %-10s %-2d %-8g %-8g\n", inventaris.masterBarang[find].kode, inventaris.masterBarang[find].nama, inventaris.masterBarang[find].jumlah, inventaris.masterBarang[find].nilaiSatuan, inventaris.masterBarang[find].nilaiTotal)
		} else {
			fmt.Println("Data tidak terdapat dalam inventaris")
		}
		case 4 :
		fmt.Print("Masukkan kode atau nama barang yang akan dihapus:")
		fmt.Scan(&cari)
		find = cariBarang(inventaris.masterBarang, inventaris.jumlahDataBarang, cari)
		if find != -1 {
			inventaris.transaksiOut[inventaris.jumlahTransaksiOut].keterangan = "Barang " + inventaris.masterBarang[find].nama + " telah dihapus"
			inventaris.transaksiOut[inventaris.jumlahTransaksiOut].kodeT = inventaris.jumlahTransaksiOut + 1
			hapusData(&inventaris.masterBarang, find, &inventaris.jumlahDataBarang)
			inventaris.jumlahTransaksiOut++
		} else {
			fmt.Println("Data sudah tidak ada dalam inventaris")
		}
		case 5:
		fmt.Print("Masukkan kode atau nama barang yang akan diubah:")
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
		fmt.Print("Pilih pengurutan yang diinginkan:")
		fmt.Scan(&opsi)
		switch opsi {
			case 1:
				sortAsc(&inventaris.masterBarang, inventaris.jumlahDataBarang)
			fmt.Println("Data sukses diurutkan menaik berdasarkan jumlah stok!")
			case 2:
				sortDesc(&inventaris.masterBarang, inventaris.jumlahDataBarang)
			fmt.Println("Data sukses diurutkan menurun berdasarkan jumlah stok!")
		}
	}
	fmt.Println()
	if opsi != 7 {
		menu(inventaris)
	}
}

func inputBarang(A *tabBarang, R *tabTransaksi, n, m *int) {
	var i, j int
	fmt.Print("Masukkan jumlah barang yang ingin di input: ")
	fmt.Scan(&j)
	fmt.Println("Silahkan masukkan data")
	for i = *n; i < (*n + j); i++ {
		fmt.Scan(&A[i].kode, &A[i].nama, &A[i].nilaiSatuan, &A[i].jumlah)
		A[i].nilaiTotal = A[i].nilaiSatuan * float64(A[i].jumlah)
		R[i].kodeT = i+1
		R[i].keterangan = "Barang " + A[i].nama + " masuk dengan nilai " + fmt.Sprint(A[i].nilaiTotal) + " dalam jumlah " + fmt.Sprint(A[i].jumlah)
	}
	*m = *m + j
	*n = *n + j
	//Memastikan tidak ada 2 baris barang yang sama
	for i = 0; i < *n; i++ {
		for j = i+1; j < *n; j++ {
			if A[i].nama == A[j].nama {
				A[i].jumlah = A[i].jumlah + A[j].jumlah
				A[i].nilaiTotal = A[i].nilaiTotal + A[j].nilaiTotal
				hapusData(A, j, n)
			}
		}
	}
}


func hapusData(A *tabBarang, m int , n *int) {
	var i int
	for i = m; i < *n; i++ {
		A[i] = A[i + 1]
	}
	*n--
}

func tampilkanData(A gudang) {
	var i, opsi int
	fmt.Println("1. Daftar Barang")
	fmt.Println("2. Riwayat Transaksi Masuk")
	fmt.Println("3. Riwayat Transaksi Keluar")
	fmt.Println("Pilih data yang ingin ditampilkan (dalam angka): ")
	fmt.Scan(&opsi)
	switch opsi {
		case 1:
		fmt.Printf("%-5s %-10s %-2s %-8s %-8s\n", "Kode", "Nama", "Jml", "NilaiSat", "NilaiTot")
		for i = 0; i < A.jumlahDataBarang; i++ {
			fmt.Printf("%-5s %-10s %-2d %-8g %-8g\n", A.masterBarang[i].kode, A.masterBarang[i].nama, A.masterBarang[i].jumlah, A.masterBarang[i].nilaiSatuan, A.masterBarang[i].nilaiTotal)
		}
		fmt.Printf("NIlai total aset adalah %.2f\n", hitungTotalAset(A.masterBarang, A.jumlahDataBarang))
		case 2:
		fmt.Printf("%-5s %-20s\n", "Kode", "Keterangan")
		for i = 0; i < A.jumlahTransaksiIn; i++ {
			fmt.Printf("%-5d %-20s\n", A.transaksiIn[i].kodeT, A.transaksiIn[i].keterangan)
		}
		case 3:
		fmt.Printf("%-5s %-20s\n", "Kode", "Keterangan")
		for i = 0; i < A.jumlahTransaksiOut; i++ {
			fmt.Printf("%-5d %-20s\n", A.transaksiOut[i].kodeT, A.transaksiOut[i].keterangan)
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
	var opsi, temp int
	fmt.Println("1. Ubah Kode")
	fmt.Println("2. Ubah Nama")
	fmt.Println("3. Ubah Nilai Satuan")
	fmt.Println("4. Ubah Jumlah")
	fmt.Print("Apa yang ingin diubah:")
	fmt.Scan(&opsi)
	fmt.Println("Masukkan perubahan")
	switch opsi {
		case 1:
		fmt.Scan(&A.masterBarang[n].kode)
		case 2:
		fmt.Scan(&A.masterBarang[n].nama)
		case 3:
		fmt.Scan(&A.masterBarang[n].nilaiSatuan)
		A.masterBarang[n].nilaiTotal = A.masterBarang[n].nilaiSatuan * float64(A.masterBarang[n].jumlah)
		case 4:
		temp = A.masterBarang[n].jumlah
		fmt.Scan(&A.masterBarang[n].jumlah)
		A.masterBarang[n].nilaiTotal = A.masterBarang[n].nilaiSatuan * float64(A.masterBarang[n].jumlah)
		if (temp - A.masterBarang[n].jumlah) < 0 {
			A.transaksiIn[A.jumlahTransaksiIn].kodeT = A.jumlahTransaksiIn+1
			A.transaksiIn[A.jumlahTransaksiIn].keterangan = "Jumlah " + A.masterBarang[n].nama + " bertambah " + fmt.Sprint(A.masterBarang[n].jumlah - temp) + " dan nilai total bertambah " + fmt.Sprint(float64(A.masterBarang[n].jumlah - temp)*A.masterBarang[n].nilaiSatuan)
			A.jumlahTransaksiIn++
		} else if (temp - A.masterBarang[n].jumlah) > 0 {
			A.transaksiOut[A.jumlahTransaksiOut].kodeT = A.jumlahTransaksiOut+1
			A.transaksiOut[A.jumlahTransaksiOut].keterangan = "Jumlah " + A.masterBarang[n].nama + " berkurang " + fmt.Sprint(temp - A.masterBarang[n].jumlah) + " dan nilai total berkurang " + fmt.Sprint(float64(temp - A.masterBarang[n].jumlah)*A.masterBarang[n].nilaiSatuan)
			A.jumlahTransaksiOut++
		}
	}
}

//Salah satunya pake selection, yang satunya pake insertion, lif.
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
		j = i - 1
		for j >= 0 && A[j].jumlah < key.jumlah {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
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
