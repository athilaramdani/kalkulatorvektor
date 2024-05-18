package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const NMAX int = 1024

type tabFloat [NMAX]float64
type tabInt [NMAX]int

type vektor struct {
	dimensi int
	titik   tabFloat
	norm    float64
}

type tabVektor [NMAX]vektor

/*
Deskripsi:
Aplikasi ini digunakan untuk melakukan berbagai operasi matematika pada vektor. Input vector diantaranya titik dan dimensi. Pengguna dapat dengan mudah melakukan perhitungan vektor, seperti penjumlahan, pengurangan, perkalian dengan skalar, produk dan titik, serta operasi hitung lainnya.
Spesifikasi:
a. Pengguna harus bisa melakukan penambahan, perubahan (edit), dan penghapusan data vektor.
b.Pengguna harus bisa menambahkan, mengubah (edit), dan menghapus dimensi serta tiap dimensi pada vektor tersebut.
c.Setelah pengguna menginput data vector (nilai titik dan dimensinya), maka otomatis menambahkan norm/panjang pada vector tersebut.
d.Pengguna bisa menampilkan semua vektor yang sudah diinput, dan menampilkan sesuai dimensi atau Panjang yang dicari.
e.Pengguna bisa menampilkan semua vektor secara terurut sesuai dimensi atau panjang vektor.
f.Pengguna bisa bisa melakukan operasi hitung pada vektor, seperti penjumlahan, hasil kali silang, hasil kali titik, jarak, atau operasi yang lain, dan menampilkan operasi tersebut serta dapat ditambahkan pada data vektor.
g.Pengguna bisa melakukan operasi hitung jika dimensi pada vektor tersebut sama, jika vektor a berisi (1,2) dan b berisi (3,2,1) vektor tersebut tidak dapat dijumlahkan dikarenakan dimensi vektor a adalah 2 sendangkan dimensi vektor b adalah 3.
*/

func randomVektorData(arrV *tabVektor, n *int, jv, mind, maxd, mint, maxt int) {
	/*
	I.S. terdefinisi array Vektor,n sebagai Jumlah Vektor, jv sebagai jumlah vektor yang akan ditambahkan, minimal dimensi, maximal dimensi, minimal titik, dan  maximal titik
	F.S. perubahan array Vektor setelah diisi random dari range min max dan jumlah vektor yang di generate serta panjangnya
	*/
	var i, j, max int
	max = *n
	rand.Seed(time.Now().UnixNano())
	for i = *n; i < jv+max; i++ {
		arrV[i].dimensi = rand.Intn(maxd-mind+1) + mind
		for j = 0; j < arrV[i].dimensi; j++ {
			arrV[i].titik[j] = float64(rand.Intn(maxt-mint+1) + mint)
		}
		*n++
	}
}

func main() {
	clearline()
	var sampah string
	var arrVektor tabVektor
	var v vektor
	var stop, orthoValid bool
	var inputMenu, nVektor, changeInput int
	nVektor = 0
	if sampah == "tes" {
		sampah = ""
	}
	randomVektorData(&arrVektor, &nVektor, 5, 1, 10, 0, 10)
	generateNorm(&arrVektor, nVektor)
	for !stop {
		start()
		generateNorm(&arrVektor, nVektor)
		tampilkanDataVektor(arrVektor, nVektor)
		menu()
		fmt.Print("Pilih opsi berapa: ")
		fmt.Scan(&inputMenu)
		clearline()

		// menu input
		if inputMenu == 1 {
			for inputMenu != 4 {
				generateNorm(&arrVektor, nVektor)
				inputMenu = 0
				tampilkanDataVektor(arrVektor, nVektor)
				menu2()
				if inputMenu == 2 && nVektor <= 0 {
					colorText("tidak ada vektor yang dapat diganti, input vektor terlebih dahulu", 0, true)
				}
				fmt.Print("Pilih opsi berapa: ")
				fmt.Scan(&inputMenu)

				// pengisian data
				if inputMenu == 1 {
					isivektor(&arrVektor, &nVektor)

					// perubah data
				} else if inputMenu == 2 {
					if nVektor > 0 {
						fmt.Print("Vektor Mana yang mau diganti? ")
						fmt.Scan(&changeInput)
						changeInput = changeInput - 1
						changeVektor(&arrVektor, changeInput)
					}
					//delete
				} else if inputMenu == 3 {
					clearline()
					tampilkanDataVektor(arrVektor, nVektor)
					fmt.Print("Vektor Mana yang mau didelete? ")
					fmt.Scan(&changeInput)
					changeInput = changeInput - 1
					if nVektor > 0  && changeInput < nVektor && changeInput >= 0{
						deleteVektor(&arrVektor, &nVektor, changeInput)
					} else if nVektor == 0{
						colorText("DATA VEKTOR KOSONG", 0, true)
						time.Sleep(1 * time.Second)
					} else {
						colorText("INPUT TIDAK VALID", 0, true)
						time.Sleep(1 * time.Second)
					}
	
					// sort
				} else if inputMenu == 4 {
					clearline()
					tampilkanDataVektor(arrVektor, nVektor)
					menuSorting()
					fmt.Print("Pilih Opsi yang ingin di sort: ")
					fmt.Scan(&inputMenu)
					sort(&arrVektor, nVektor, inputMenu)

					// search
				} else if inputMenu == 5 {
					clearline()
					tampilkanDataVektor(arrVektor, nVektor)
					menuSearch()
					fmt.Print("Pilih Opsi yang ingin di search: ")
					fmt.Scan(&inputMenu)
					search(arrVektor, nVektor, inputMenu)
					// reset
				} else if inputMenu == 6 {
					resetVektor(&arrVektor, &nVektor)

					//generate random
				} else if inputMenu == 7 {
					generateRandomVektor(&arrVektor, &nVektor)
				} else if inputMenu == 8 {
					clearline()
					break
				}
				clearline()
			}
			// menu Penjumlahan Start
		} else if inputMenu == 2 {

		} else if inputMenu == 3 {
			if nVektor > 0 {
				v = penjumlahanVektor(arrVektor, nVektor)
				tambahVektorKeData(&arrVektor, v, &nVektor)
				clearline()
			} else {
				fmt.Println("input Vektor terlebih dahulu")
			}
			// menu Pengurangan Start
		} else if inputMenu == 4 {
			if nVektor > 0 {
				v = penguranganVektor(arrVektor, nVektor)
				tambahVektorKeData(&arrVektor, v, &nVektor)
				clearline()
			} else {
				fmt.Println("input Vektor terlebih dahulu")
			}

			// cari titik Awal
		} else if inputMenu == 5 {
			v = cariTitikAwal(arrVektor, nVektor)
			tambahVektorKeData(&arrVektor, v, &nVektor)
			clearline()

			//cari titik akhir
		} else if inputMenu == 6 {
			v = cariTitikAkhir(arrVektor, nVektor)
			tambahVektorKeData(&arrVektor, v, &nVektor)
			clearline()

			//nilai dari 2 titik
		} else if inputMenu == 7 {
			v = nilaiDariDuaTitik(arrVektor, nVektor)
			tambahVektorKeData(&arrVektor, v, &nVektor)
			clearline()
		
				//dot product
		} else if inputMenu == 8 {
			dotProduct(arrVektor, nVektor)
			fmt.Print("Ketik apapun untuk Lanjut: ")
			fmt.Scan(&sampah)
			clearline()

				//cross product
		} else if inputMenu == 9 {
			v = nilaiDariDuaTitik(arrVektor, nVektor)
			tambahVektorKeData(&arrVektor, v, &nVektor)
			clearline()

			//orthogonal
		} else if inputMenu == 10 {
			orthoValid = orthogonalValidation(arrVektor, nVektor)
			fmt.Println("Orthogonal =", orthoValid)
			fmt.Print("Ketik apapun untuk Lanjut: ")
			fmt.Scan(&sampah)
			clearline()

			//exit
		} else if inputMenu == 20 {
			clearline()
			break
			// end Else
		} else {
			fmt.Println("MASUKKAN INPUT YANG BENAR")
		}
	}
}

func search(arrV tabVektor, n, x int) {
	var value,i,j int
	var valid bool
	var sampah string
	if x == 1 {
		fmt.Print("Masukkan nilai titik yang dicari: ")
		fmt.Scan(&value)
		if n > 0 {
			fmt.Printf("\033[0;37mdata Vektor yang mengandung nilai titik \033[0;33m%d\033[0m: \n", value)
		}
		for i = 0; i < n; i++ {
			valid = false
			j = 0
			for !valid && j < arrV[i].dimensi {
				valid = arrV[i].titik[j] == float64(value)
				j++
			}
			if valid {
				fmt.Printf("%d. \033[0;37m( ", i+1)
		
				for j = 0; j < arrV[i].dimensi; j++ {
					if arrV[i].titik[j] == float64(value) {
						fmt.Printf("\033[0;33m%v \033[0;37m",arrV[i].titik[j])
					} else {
						fmt.Printf("%v ",arrV[i].titik[j])
					}
				}
				fmt.Printf(") dimension: %d , norm/panjang: %0.2f\n\033[0m", arrV[i].dimensi, arrV[i].norm)
			}
		}
	} else if x == 2 {
		fmt.Print("Masukkan nilai Dimensi yang dicari: ")
		fmt.Scan(&value)
		if n > 0 {
			fmt.Printf("\033[0;37mdata Vektor yang mengandung nilai Dimensi \033[0;33m%d\033[0m: \n", value)
		}
		for i = 0; i < n; i++ {
			valid = false
			j = 0
			for !valid && j < arrV[i].dimensi {
				valid = arrV[i].dimensi == value
				j++
			}
			if valid {
				fmt.Printf("%d. \033[0;37m( ", i+1)
		
				for j = 0; j < arrV[i].dimensi; j++ {
					fmt.Printf("%v ",arrV[i].titik[j])
				}
				fmt.Printf(") dimension: \033[0;33m%d033[0;37m, norm/panjang: %0.2f\n\033[0m", arrV[i].dimensi, arrV[i].norm)
			}
		}
	} else if x == 3 {
		fmt.Print("Masukkan nilai Norm yang dicari: ")
		fmt.Scan(&value)
		if n > 0 {
			fmt.Printf("\033[0;37mdata Vektor yang mengandung nilai Norm \033[0;33m%d\033[0m: \n", value)
		}
		for i = 0; i < n; i++ {
			valid = false
			j = 0
			for !valid && j < arrV[i].dimensi {
				valid = arrV[i].norm >= float64(value) && arrV[i].norm <= float64(value) + 0.9999
				j++
			}
			if valid {
				fmt.Printf("%d. \033[0;37m( ", i+1)
		
				for j = 0; j < arrV[i].dimensi; j++ {
					fmt.Printf("%v ",arrV[i].titik[j])
				}
				fmt.Printf(") dimension: %d, norm/panjang: \033[0;33m%0.2f\n\033[0m", arrV[i].dimensi, arrV[i].norm)
			}
		}
	}
	fmt.Print("Ketik apapun untuk Lanjut: ")
	fmt.Scan(&sampah)
	if sampah == "asidjasiodiasioasdjias" {
		fmt.Println()
	}
}

func deleteVektor(arrV *tabVektor, n *int, x int) {
	for i := x; i < *n; i++ {
		swap(&*arrV, i, i+1)
	}
	*n = *n -1
}

func sort(arrV *tabVektor, n, x int) {
    var opsi, i, idx int
    fmt.Println("1. descending (Terurut Terbesar ke Terkecil)")
    fmt.Println("2. ascending (Terurut Terkecil ke Terbesar)")
    fmt.Print("Pilih Opsi: ")
    fmt.Scan(&opsi)
    n = n - 1
    if x == 1 { // Sort by dimensi
        if opsi == 1 {
            // Sort descending by dimensi
            for i = 0; i <= n; i++ {
                idx = findMinDimensi(*arrV, n-i)
                swap(arrV, idx, n-i)
            }
        } else if opsi == 2 {
            // Sort ascending by dimensi
            for i = 0; i <= n; i++ {
                idx = findMaxDimensi(*arrV, n-i)
                swap(arrV, idx, n-i)
            }
        }
    } else if x == 2 { // Sort by norm
        if opsi == 1 {
            // Sort descending by norm
            for i = 0; i <= n; i++ {
                idx = findMinNorm(*arrV, n-i)
                swap(arrV, idx, n-i)
            }
        } else if opsi == 2 {
            // Sort ascending by norm
            for i = 0; i <= n; i++ {
                idx = findMaxNorm(*arrV, n-i)
                swap(arrV, idx, n-i)
            }
        }
    }
}

func findMaxDimensi(arrV tabVektor, end int) int {
    var imax int
    imax = 0
    for i := 1; i <= end; i++ { // <= end to include the end element
        if arrV[i].dimensi > arrV[imax].dimensi {
            imax = i
        }
    }
    return imax
}

func findMinDimensi(arrV tabVektor, end int) int {
    var imin int
    imin = 0
    for i := 1; i <= end; i++ { // <= end to include the end element
        if arrV[i].dimensi < arrV[imin].dimensi {
            imin = i
        }
    }
    return imin
}

func findMaxNorm(arrV tabVektor, end int) int {
    var imax int
    imax = 0
    for i := 1; i <= end; i++ { // <= end to include the end element
        if arrV[i].norm > arrV[imax].norm {
            imax = i
        }
    }
    return imax
}

func findMinNorm(arrV tabVektor, end int) int {
    var imin int
    imin = 0
    for i := 1; i <= end; i++ { // <= end to include the end element
        if arrV[i].norm < arrV[imin].norm {
            imin = i
        }
    }
    return imin
}

func swap(arrV *tabVektor, idx1, idx2 int) {
    temp := arrV[idx1]
    arrV[idx1] = arrV[idx2]
    arrV[idx2] = temp
}

func generateRandomVektor(arrV *tabVektor, n *int) {
	var jumlahVektor, minDimensi, maxDimensi, minTitik, maxTitik int
	fmt.Print("masukkan jumlah vektor yang ingin digenerate: ")
	fmt.Scan(&jumlahVektor)
	for jumlahVektor < 0 {
		colorText("masukkan tidak valid, negatif tidak diperuntukkan", 0, true)
		fmt.Print("masukkan jumlah vektor yang ingin digenerate: ")
		fmt.Scan(&jumlahVektor)
	}
	for jumlahVektor > NMAX-*n {
		colorText("masukkan tidak valid, data tidak boleh lebih dari 1024", 0, true)
		fmt.Print("masukkan jumlah vektor yang ingin digenerate: ")
		fmt.Scan(&jumlahVektor)
	}
	for jumlahVektor == 0 {
		colorText("masukkan tidak valid, tidak ada vektor yang ditambahkan", 0, true)
		fmt.Print("masukkan jumlah vektor yang ingin digenerate: ")
		fmt.Scan(&jumlahVektor)
	}
	fmt.Print("masukkan min Dimensi: ")
	fmt.Scan(&minDimensi)
	for minDimensi <= 0{
		colorText("masukkan tidak valid, vektor tidak mempunyai dimensi", 0, true)
		fmt.Print("masukkan min Dimensi: ")
		fmt.Scan(&minDimensi)
	}
	fmt.Print("masukkan max Dimensi: ")
	fmt.Scan(&maxDimensi)
	for maxDimensi < minDimensi {
		colorText("masukkan tidak valid, MaxDimensi kurang dari MinDimensi", 0, true)
		fmt.Print("masukkan max Dimensi: ")
		fmt.Scan(&maxDimensi)
	}
	for maxDimensi > 1024 {
		colorText("masukkan tidak valid, MaxDimensi tidak boleh lebih dari 1024", 0, true)
		fmt.Print("masukkan max Dimensi: ")
		fmt.Scan(&maxDimensi)
	}
	fmt.Print("masukkan min Titik: ")
	fmt.Scan(&minTitik)
	fmt.Print("masukkan max Titik: ")
	fmt.Scan(&maxTitik)
	for maxTitik < minTitik {
		colorText("masukkan tidak valid, MaxTitik kurang dari MinTitik", 0, true)
		fmt.Print("masukkan max Titik: ")
		fmt.Scan(&maxTitik)
	}
	randomVektorData(&*arrV, &*n, jumlahVektor, minDimensi, maxDimensi, minTitik, maxTitik)
}

func tambahVektorKeData(arrV *tabVektor, v vektor, n *int) {
	var sentinel string
	fmt.Print(" = ")
	tampilkanVektor(v)
	fmt.Println()
	fmt.Println("Apakah vektor tersebut ingin dimasukkan ke data? (Y/N)")
	fmt.Scan(&sentinel)
	if sentinel == "Y" || sentinel == "y" {
		menambahkan(arrV, v, *n)
		*n++
	}
}

func generateNorm(arrV *tabVektor, n int) {
	var i int
	var norm float64
	for i = 0; i < n; i++ {
		norm = math.Sqrt(jumlahKuadratVektor(arrV[i]))
		arrV[i].norm = norm
	}
}

func jumlahKuadratVektor(v vektor) float64 {
	jkv := 0.0
	for i := 0; i < v.dimensi; i++ {
		jkv += math.Pow(float64(v.titik[i]), 2)
	}
	return jkv
}

func isivektor(arrV *tabVektor, n *int) {
	var v vektor
	var sentinel string
	sentinel = "y"
	for sentinel == "y" || sentinel == "Y" {
		v = createdvektor(*arrV)
		menambahkan(&*arrV, v, *n)
		fmt.Print("lanjutkan pengisian (Y/N): ")
		fmt.Scan(&sentinel)
		*n++
	}
}

func resetVektor(arrV *tabVektor, n *int) {
	var i, j int
	for i = 0; i < *n; i++ {
		for j = 0; j < arrV[i].dimensi; j++ {
			arrV[i].titik[j] = 0
		}
	}
	*n = 0
}

func changeVektor(arrV *tabVektor, c int) {
	v := createdvektor(*arrV)
	menambahkan(&*arrV, v, c)
}

func createdvektor(arrV tabVektor) vektor {
	var v vektor
	var i int
	fmt.Print("Masukkan dimensi vektor: ")
	fmt.Scan(&v.dimensi)
	fmt.Print("Masukkan nilai vektor sesuai dimensinya yaitu ", v.dimensi, " : ")
	for i = 0; i < v.dimensi; i++ {
		fmt.Scan(&v.titik[i])
	}
	return v
}

func penjumlahanVektor(arrV tabVektor, n int) vektor {
	var v vektor
	var x, i, j int
	var y tabInt
	var dimensiValid bool
	dimensiValid = false
	for !dimensiValid {
		for i = 0; i < 1024; i++ {
			v.titik[i] = 0
		}
		tampilkanDataVektor(arrV, n)
		fmt.Print("ada berapa vektor yang ingin dijumlahkan? ")
		fmt.Scan(&x)
		fmt.Print("masukkan vektor apa saja yang ingin dijumlah (mengurut pada data vektornya): ")
		fmt.Scan(&y[0])
		y[0] = y[0] - 1
		for i = 1; i < x; i++ {
			fmt.Scan(&y[i])
			y[i] = y[i] - 1
		}
		dimensiValid = cekKesamaanDimensi(arrV, x, y)
		if !dimensiValid {
			colorText("dimensi nya tidak sama", 0, true)
		}
	}
	v.dimensi = arrV[y[0]].dimensi
	v.titik = arrV[y[0]].titik
	tampilkanVektor(arrV[y[0]])
	for i = 1; i < x; i++ {
		fmt.Print(" + ")
		tampilkanVektor(arrV[y[i]])
		for j = 0; j < arrV[y[0]].dimensi; j++ {
			v.titik[j] = v.titik[j] + arrV[y[i]].titik[j]
		}
	}
	return v
}

func penguranganVektor(arrV tabVektor, n int) vektor {
	var v vektor
	var x, i, j int
	var y tabInt
	var dimensiValid bool
	dimensiValid = false
	for !dimensiValid {
		for i = 0; i < 1024; i++ {
			v.titik[i] = 0
		}
		tampilkanDataVektor(arrV, n)
		fmt.Print("ada berapa vektor yang ingin dikurangi? ")
		fmt.Scan(&x)
		fmt.Print("masukkan vektor apa saja yang ingin dikurang (mengurut pada data vektornya): ")
		fmt.Scan(&y[0])
		y[0] = y[0] - 1
		for i = 1; i < x; i++ {
			fmt.Scan(&y[i])
			y[i] = y[i] - 1
		}
		dimensiValid = cekKesamaanDimensi(arrV, x, y)
		if !dimensiValid {
			colorText("dimensi nya tidak sama", 0, true)
		}
	}
	v.dimensi = arrV[y[0]].dimensi
	v.titik = arrV[y[0]].titik
	tampilkanVektor(arrV[y[0]])
	for i = 1; i < x; i++ {
		fmt.Print(" - ")
		tampilkanVektor(arrV[y[i]])
		for j = 0; j < arrV[y[0]].dimensi; j++ {
			v.titik[j] = v.titik[j] - arrV[y[i]].titik[j]
		}
	}
	return v
}

func cariTitikAwal(arrV tabVektor, n int) vektor {
	var v vektor
	var i, j int
	var y tabInt
	var dimensiValid bool
	dimensiValid = false
	for !dimensiValid {
		for i = 0; i < 1024; i++ {
			v.titik[i] = 0
		}
		tampilkanDataVektor(arrV, n)
		fmt.Print("masukkan Titik Akhir: ")
		fmt.Scan(&y[0])
		y[0] = y[0] - 1
		fmt.Print("masukkan Vektor: ")
		fmt.Scan(&y[1])
		y[1] = y[1] - 1
		dimensiValid = cekKesamaanDimensi(arrV, 2, y)
		if !dimensiValid {
			colorText("dimensi nya tidak sama", 0, true)
		}
	}
	v.dimensi = arrV[y[0]].dimensi
	v.titik = arrV[y[0]].titik
	tampilkanVektor(arrV[y[0]])
	for i = 1; i < 2; i++ {
		fmt.Print(" - ")
		tampilkanVektor(arrV[y[i]])
		for j = 0; j < arrV[y[0]].dimensi; j++ {
			v.titik[j] = v.titik[j] - arrV[y[i]].titik[j]
		}
	}
	return v
}

func cariTitikAkhir(arrV tabVektor, n int) vektor {
	var v vektor
	var i, j int
	var y tabInt
	var dimensiValid bool
	dimensiValid = false
	for !dimensiValid {
		for i = 0; i < 1024; i++ {
			v.titik[i] = 0
		}
		tampilkanDataVektor(arrV, n)
		fmt.Print("masukkan Titik Awal: ")
		fmt.Scan(&y[0])
		y[0] = y[0] - 1
		fmt.Print("masukkan Vektor: ")
		fmt.Scan(&y[1])
		y[1] = y[1] - 1
		dimensiValid = cekKesamaanDimensi(arrV, 2, y)
		if !dimensiValid {
			colorText("dimensi nya tidak sama", 0, true)
		}
	}
	v.dimensi = arrV[y[0]].dimensi
	v.titik = arrV[y[0]].titik
	tampilkanVektor(arrV[y[0]])
	for i = 1; i < 2; i++ {
		fmt.Print(" - ")
		tampilkanVektor(arrV[y[i]])
		for j = 0; j < arrV[y[0]].dimensi; j++ {
			v.titik[j] = v.titik[j] + arrV[y[i]].titik[j]
		}
	}
	return v
}

func nilaiDariDuaTitik(arrV tabVektor, n int) vektor {
	var v vektor
	var i int
	var y tabInt
	var dimensiValid bool
	dimensiValid = false
	for !dimensiValid {
		for i = 0; i < 1024; i++ {
			v.titik[i] = 0
		}
		tampilkanDataVektor(arrV, n)
		fmt.Print("masukkan Titik Awal: ")
		fmt.Scan(&y[0])
		y[0] = y[0] - 1
		fmt.Print("masukkan Titik Akhir: ")
		fmt.Scan(&y[1])
		y[1] = y[1] - 1
		dimensiValid = cekKesamaanDimensi(arrV, 2, y)
		if !dimensiValid {
			colorText("dimensi nya tidak sama", 0, true)
		}
	}
	v.dimensi = arrV[y[1]].dimensi
	v.titik = arrV[y[1]].titik
	tampilkanVektor(arrV[y[1]])
	fmt.Print(" - ")
	tampilkanVektor(arrV[y[0]])
	for i = 0; i < arrV[y[0]].dimensi; i++ {
		v.titik[i] = v.titik[i] - arrV[y[0]].titik[i]
	}
	return v
}

func dotProduct(arrV tabVektor, n int) {
	var y tabInt
	var dimensiValid bool
	dimensiValid = false
	for !dimensiValid {
		tampilkanDataVektor(arrV, n)
		fmt.Print("masukkan Vektor 1: ")
		fmt.Scan(&y[0])
		y[0] = y[0] - 1
		fmt.Print("masukkan Vektor 2: ")
		fmt.Scan(&y[1])
		y[1] = y[1] - 1
		dimensiValid = cekKesamaanDimensi(arrV, 2, y)
		if !dimensiValid {
			colorText("dimensi nya tidak sama", 0, true)
		}
	}
	test := dotProductValue(arrV[y[0]],arrV[y[1]])
	if test == 101010 {
		fmt.Print("")
	}
}

func dotProductValue(v1,v2 vektor) float64 {
	var jum float64
	jum = 0
	fmt.Print("( ")
	for i:=0; i< v1.dimensi; i++ {
		jum += v1.titik[i] * v2.titik[i]
		fmt.Print(v1.titik[i]," * ",v2.titik[i] ," ")
		if i < v1.dimensi - 1 {
			fmt.Print("+ ")
		}
	}
	fmt.Print(") = ", jum, "\n")
	return jum
}

func orthogonalValidation(arrV tabVektor, n int) bool {
	var jum float64
	var i,j,x int
	var y tabInt
	var dimensiValid, orthoValid bool
	dimensiValid = false
	for !dimensiValid {
		tampilkanDataVektor(arrV, n)
		fmt.Print("ada berapa vektor yang ingin dicek orthogonalitas nya? ")
		fmt.Scan(&x)
		fmt.Print("masukkan vektor apa saja yang ingin dicek (mengurut pada data vektornya): ")
		fmt.Scan(&y[0])
		y[0] = y[0] - 1
		for i = 1; i < x; i++ {
			fmt.Scan(&y[i])
			y[i] = y[i] - 1
		}
		dimensiValid = cekKesamaanDimensi(arrV, x, y)
		if !dimensiValid {
			colorText("dimensi nya tidak sama", 0, true)
		}
	}
	i = 0
	orthoValid = true
	for i < x && orthoValid == true {
		j = i + 1
		for j < x && orthoValid == true {
			jum = dotProductValue(arrV[y[i]],arrV[y[j]])
			orthoValid = jum == 0
			fmt.Println(orthoValid)
			j++
		}
		i++
	}
	return orthoValid
}

func cekKesamaanDimensi(A tabVektor, n int, cari tabInt) bool {
	sama := true
	for i := 1; i < n && sama; i++ {
		sama = A[cari[0]].dimensi == A[cari[i]].dimensi
	}
	return sama
}

func menambahkan(arrV *tabVektor, v vektor, n int) {
	arrV[n] = v
}

func start() {
	gsd3kali(17)
	fmt.Println()
	colorText("   Welcome To Vector Calculator AthDanz Zero Six", 2, true)
	gsd3kali(17)
	fmt.Println()
}

func gsd3kali(n int) {
	var i int
	for i = 1; i <= n; i++ {
		colorText("===", i%4, false)
	}
}

func menu() {
	colorText("========== MENU =========", 3, true)
	colorText("1. Edit Data", 0, true)
	colorText("2. Edit Visual Data", 1, true)
	colorText("3. Jumlahkan", 4, true)
	colorText("4. Kurangi", 4, true)
	colorText("5. Cari Titik Awal", 4, true)
	colorText("6. Cari Titik Akhir", 4, true)
	colorText("7. Cari vektor dari 2 Titik", 4, true)
	colorText("8. Dot Product", 4, true)
	colorText("9. Cross Product", 4, true)
	colorText("10. cek orthogonal", 4, true)
	colorText("20. Exit", 0, true)
}

func menu2() {
	colorText("========== MENU INPUT =========", 3, true)
	colorText("1. input", 2, true)
	colorText("2. ganti", 2, true)
	colorText("3. delete", 0, true)
	colorText("4. sort", 5, true)
	colorText("5. search", 5, true)
	colorText("6. reset", 0, true)
	colorText("7. Generate Random Vektor", 3, true)
	colorText("8. exit", 4, true)
}

func menuSorting() {
	colorText("========== MENU SORTING =========", 3, true)
	colorText("1. Sort Dimensi", 2, true)
	colorText("2. Sort Norm", 2, true)
}

func menuSearch() {
	colorText("========== MENU SEARCH =========", 3, true)
	colorText("1. Search Nilai Titik", 2, true)
	colorText("2. Search Dimensi", 2, true)
	colorText("3. Search Norm", 2, true)
}

func colorText(text string, color int, line bool) {
	/*I.S. terdefinisi text sebagai string dan color sebagai warna
	F.S. Output dari string dengan warnanya*/
	const reset string = "\033[0m"
	var colorCode string
	if color == 0 {
		//red
		colorCode = "\033[0;31m"
	} else if color == 1 {
		//green
		colorCode = "\033[0;32m"
	} else if color == 2 {
		//cyan
		colorCode = "\033[0;36m"
	} else if color == 3 {
		//pink
		colorCode = "\033[0;35m"
	} else if color == 4 {
		//white
		colorCode = "\033[0;37m"
	} else if color == 5 {
		//yellow
		colorCode = "\033[0;33m"
	} else {
		colorCode = ""
	}
	if line {
		fmt.Printf("%s%s%s\n", colorCode, text, reset)
	} else {
		fmt.Printf("%s%s%s", colorCode, text, reset)
	}
}

func tampilkanVektor(v vektor) {
	fmt.Print("( ")
	for i := 0; i < v.dimensi; i++ {
		fmt.Print(v.titik[i], " ")
	}
	fmt.Print(")")
}

func tampilkanDataVektor(arrV tabVektor, n int) {
	var i, j int
	if n > 0 {
		colorText("data Vektor: ", 5, true)
	}
	for i = 0; i < n; i++ {
		fmt.Printf("%d. \033[0;37m( ", i+1)
		for j = 0; j < arrV[i].dimensi; j++ {
			fmt.Print(arrV[i].titik[j], " ")
		}
		fmt.Printf(") dimension: %d , norm/panjang: %0.2f\n\033[0m", arrV[i].dimensi, arrV[i].norm)
	}
}

func clearline() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
