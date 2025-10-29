// I.S. Merupakan inputan sejumlah tim yang mengikuti turnamen dengan status Skor,Menang dan Kalah adalah 0

// F.S. Meruapakan tim yang juara dari turnamen serta klasemen tim yang mengikuti 
// Memberikan hasil status Skor,Menang dan Kalah sesuai dengan input Hasil Pertandingan
// Memberikan tim yang menjuarai turnamen

//Bukti buat WGTIK. 103032400055 IT-48-05.
package main
import "fmt"


type Peserta struct {
	Nama   string
	Skor   int
	Menang int
	Kalah  int
}

type Waktu struct {
	bulan int
	tanggal int
	jam string
}

type Pertandingan struct {
	IDTim1   int
	IDTim2   int
	SkorTim1 int
	SkorTim2 int
	waktu  Waktu
}

const NMAX int = 50

var peserta [NMAX]Peserta
var jumlahPeserta int
var totalPeserta int

var pertandingan [NMAX]Pertandingan
var jumlahPertandingan int

func tambahPeserta() {
	var nama string
	var found bool
	
	if jumlahPeserta >= totalPeserta {
		fmt.Println("Jumlah Peserta Sudah memenuhi kapasitas.")
		return
	}
	
	fmt.Print("(Ketik end untuk batal.)Masukkan nama tim: ")
	fmt.Scan(&nama)
	for nama != "end" && jumlahPeserta < totalPeserta {
		found = false
		for i := 0; i < jumlahPeserta; i++ {
			if nama == peserta[i].Nama {
				fmt.Println("nama tim sudah ada.")
				found = true
				break
			}
		}
		
		if !found {
		peserta[jumlahPeserta] = Peserta{Nama: nama, Skor: 0, Menang: 0, Kalah: 0}
		fmt.Println("Peserta Sukses Ditambahkan")
		jumlahPeserta++
		}
		
		if jumlahPeserta < totalPeserta {
			fmt.Print("(Ketik end untuk batal.)Masukkan nama tim: ")
			fmt.Scan(&nama)
		}
	}
}

func jadwalkanPertandingan(){
	var id1, id2 int
	var tgl, bln int
	var jm string
	
	if jumlahPeserta < 2 {
		fmt.Println("Belum cukup tim untuk menjadwalkan pertandingan.")
		return
	}
	
	fmt.Println("Daftar Tim:")
	for i := 0; i < jumlahPeserta; i++ {
		fmt.Printf("ID: %d - %s\n", i, peserta[i].Nama)
	}

	
	fmt.Print("ID Tim 1: ")
	fmt.Scan(&id1)
	fmt.Print("ID Tim 2: ")
	fmt.Scan(&id2)
	
	if id1 < 0 || id1 >= jumlahPeserta || id2 < 0 || id2 >= jumlahPeserta {
		fmt.Println("Salah satu ID tim tidak valid.")
		return
	}
	
	if id1 == id2 {
		fmt.Println("ngga bisa gitu!")
		return
	}

	fmt.Print("Tentuin Waktunya (Jam Tanggal Bulan): ")
	fmt.Scan(&jm,&tgl,&bln)
	
	pertandingan[jumlahPertandingan] = Pertandingan{IDTim1: id1, IDTim2: id2, waktu : Waktu { jam : jm,tanggal : tgl,bulan : bln,},}
	jumlahPertandingan++
	fmt.Println("Pertandingan sukses dijadwalkan!")
}

func jadwalPertandingan() {
	var p Pertandingan
	var pilihan int
	
	if jumlahPertandingan == 0 {
		fmt.Println("Belum ada pertandingan yang dijadwalkan.")
		return
	}

	fmt.Println("\n===== Jadwal Pertandingan =====")
	for i := 0; i < jumlahPertandingan; i++ {
		p = pertandingan[i]
		namaTim1 := peserta[p.IDTim1].Nama
		namaTim2 := peserta[p.IDTim2].Nama
		fmt.Printf("%d. %s vs %s | Waktu Match Dilakukan: %s %d %d", i , namaTim1, namaTim2, p.waktu.jam, p.waktu.tanggal, p.waktu.bulan)
		
		
		if p.SkorTim1 != 0 || p.SkorTim2 != 0 {
			fmt.Printf(" | Skor: %d - %d", p.SkorTim1, p.SkorTim2)
		} else {
			fmt.Print(" | Belum dimainkan")
		}
		fmt.Println()
	}
	
	for {
		fmt.Println("\n=== Menu Jadwal ===")
		fmt.Println("1. Hapus Jadwal")
		fmt.Println("2. FilterTanggal")
		fmt.Println("3. FilterBulan")
		fmt.Println("4. Kembali")
		fmt.Println("=================")
		fmt.Print("Pilihan mu: ")
		
		fmt.Scan(&pilihan)
		switch pilihan {
			case 1:
				hapusJadwal()
			case 2:
				filterTanggal()
			case 3:
				filterBulan()
			case 4:
				return
			default :
				fmt.Println("pilihan mu ga valid.")
		}
	}
}



func filterTanggal() {
	var tanggalCari int
	var ada bool
	
	ada = false
	fmt.Println("mau cari tanggal berapa?")
	fmt.Scan(&tanggalCari)
	for i := 0 ; i < jumlahPertandingan; i++  {
		if pertandingan[i].waktu.tanggal == tanggalCari {
				fmt.Printf("%s vs %s | Jam: %s\n", peserta[pertandingan[i].IDTim1].Nama,peserta[pertandingan[i].IDTim2].Nama, pertandingan[i].waktu.jam)
				ada = true
		}
	}
	
	if !ada {
		fmt.Println("tanggal itu ga ada pertandingan")
	}
}

func filterBulan() {
	var bulanCari int
	var ada bool
	
	ada = false
	fmt.Println("mau cari bulan berapa? ")
	fmt.Scan(&bulanCari)
	for i := 0; i < jumlahPertandingan; i++ {
		if pertandingan[i].waktu.bulan == bulanCari {
			fmt.Printf("%s vs %s | Jam: %s\n", peserta[pertandingan[i].IDTim1].Nama,peserta[pertandingan[i].IDTim2].Nama, pertandingan[i].waktu.jam)
			ada = true
		}  	
	}
	
	if !ada {
		fmt.Println("bulan itu ga ada pertandingan")
	}
}

func hapusJadwal() {
	var index int
    fmt.Print("Masukkan index pertandingan yang ingin dihapus: ")
    fmt.Scan(&index)

    if index < 0 || index >= jumlahPertandingan {
        fmt.Println("Index tidak valid.")
        return
    }

    for i := index; i < jumlahPertandingan-1; i++ {
        pertandingan[i] = pertandingan[i+1]
    }
    jumlahPertandingan--

    fmt.Println("Pertandingan berhasil dihapus.")
}

func menuTurnamen() {
	if totalPeserta == 0 {
		fmt.Println("kok belum ada peserta? balik deh ke start menu")
		return
	}
	
	for {
		fmt.Println("\n=== Menu Turnamen E-Sports ===")
		fmt.Println("1. Tambah Peserta")
		fmt.Println("2. Jadwalkan Pertandingan")
		fmt.Println("3. Lihat Jadwal Pertandingan")
		fmt.Println("4. Input Hasil Pertandingan")
		fmt.Println("5. Tampilkan Klasemen")
		fmt.Println("6. Kembali ke Start Menu")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahPeserta()
		case 2:
			jadwalkanPertandingan()
		case 3:
			jadwalPertandingan()
		case 4:
			inputHasilPertandingan()
		case 5:
			tampilkanKlasemen()
		case 6:
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func inputHasilPertandingan() {
	var index int
	fmt.Print("Masukkan indeks pertandingan: ")
	fmt.Scan(&index)
	
	if index < 0 || index >= jumlahPeserta {
		fmt.Println("Index tidak valid.")
		return
	}
	
	if pertandingan[index].SkorTim1 != 0 || pertandingan[index].SkorTim2 != 0 {
		fmt.Println("Pertandingan ini sudah selesai.")
		return
	}

	fmt.Print("Skor Tim 1: ")
	fmt.Scan(&pertandingan[index].SkorTim1)
	fmt.Print("Skor Tim 2: ")
	fmt.Scan(&pertandingan[index].SkorTim2)
	
	id1 := pertandingan[index].IDTim1
	id2 := pertandingan[index].IDTim2

	
	if pertandingan[index].SkorTim1 > pertandingan[index].SkorTim2 {
		peserta[id1].Skor += 3
		peserta[id1].Menang++
		peserta[id2].Kalah++
	} else if pertandingan[index].SkorTim2 > pertandingan[index].SkorTim1 {
		peserta[id2].Skor += 3
		peserta[id2].Menang++
		peserta[id1].Kalah++
	} else {
		peserta[id1].Skor += 1
		peserta[id2].Skor += 1
	}
	fmt.Println("Hasil pertandingan sukses disimpan!")
}

func tampilkanKlasemen() { 
	var p Peserta
	selectionSort()
	
	fmt.Println("\n===== Klasemen =====")
	for i := 0; i < jumlahPeserta; i++ {
		p = peserta[i]
		fmt.Printf("%d. %s | Poin: %d | Menang: %d | Kalah: %d\n", i + 1, p.Nama, p.Skor, p.Menang, p.Kalah)
	}
}

func selectionSort() {
	var pass, i, idx int
	var temp Peserta
	
	pass = 1
	n := jumlahPeserta
	for pass < n {
		idx = pass - 1
		i = pass
		for i < n {
			if peserta[idx].Skor < peserta[i].Skor {
				idx = i
			}
			i++
		}
		temp = peserta[pass - 1]
		peserta[pass - 1] = peserta[idx]
		peserta[idx] = temp
		pass++
	}
}



func main() {
	for {
		fmt.Println("\n=== START MENU ===")
		fmt.Println("1. Tentukan Jumlah Peserta")
		if totalPeserta > 0 {
			fmt.Println("2. Kembali ke Menu Turnamen")
		}
		fmt.Println("3. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)
		
		switch pilihan {
		case 1:
			if totalPeserta != 0 {
				fmt.Println("Jumlah peserta udah ada.")
				break
			}
			
			fmt.Println("(Batas Maksimal : 50)")
			fmt.Print("Masukkan jumlah peserta: ")
			fmt.Scan(&totalPeserta)

			if totalPeserta > 0 {
				if totalPeserta > NMAX {
					var konfirmasi string
					fmt.Printf("Jumlah peserta melebihi batas maksimum (%d). Yakin lanjut (yes/no)? \n", NMAX)
					fmt.Printf("batas maksimum akan tetap sama.\n")
					fmt.Scan(&konfirmasi)

					if konfirmasi == "yes" || konfirmasi == "YES" {
						totalPeserta = NMAX
						menuTurnamen()
					}  else {
						fmt.Println("Silakan masukkan jumlah peserta yang valid.")
					}
				} else {
					menuTurnamen()
				}
			} else {
				fmt.Println("Jumlah tidak boleh kurang dari 1.")
			}

		case 2:
			if totalPeserta > 0 {
				menuTurnamen()
			} else {
				fmt.Println("error : Belum ada Peserta")
				break
			}
		case 3:
			fmt.Println("TERIMAKASIH TELAH MENGGUNAKAN PROGRAM INI")
			fmt.Println("=== Tubes Algoritma Pemograman ===")
			fmt.Println("Shalby Alghifari Firmansyah (103032400055)")
			fmt.Println("Rafif Hendra Saputra (103032400111)")
			return
		default:
			fmt.Println("Pilihan mu ga valid.")
		}
	}
}
