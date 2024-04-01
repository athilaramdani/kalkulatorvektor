package main

import (
	"fmt"
	"math"
)

type vektor struct {
	x, y,z float64
}
func main() {
	version()
	var (
		menuInput int
		vektor1, vektor2,vektor3 vektor
		dimensir2 bool = false
	)
	fmt.Println("")
	loop1:
		for {
			menu(vektor1,vektor2,vektor3, dimensir2 )
			fmt.Print("Pilih apa? ")
			fmt.Scan(&menuInput)
			switch menuInput {
			case 1:
				vektor1 = buatvektorr2()
				vektor2 = buatvektorr2()
				vektor3 = buatvektorr2()
				fmt.Print("apakah 2 dimensi? (y,n): " )
				var pastikan string
				fmt.Scan(&pastikan)
				if pastikan == "y" {
					dimensir2 = true
				}
			case 2:
				fmt.Println("Hasil Penjumlahan:")
				jumlahVektor(vektor1, vektor2,dimensir2)
			case 3:
				var scalar float64
				fmt.Print("Masukkan scalar: ")
				fmt.Scan(&scalar)
				kaliSkalarvektor(vektor1, scalar,dimensir2)
				kaliSkalarvektor(vektor2, scalar,dimensir2)
			case 4:
				fmt.Print("Hasil Kali Titik: ")
				fmt.Println(kaliTitikvektor(vektor1, vektor2,dimensir2))
			case 5:
				fmt.Print("Hasil Kali Silang: ")
				fmt.Println(kaliSilangvektor(vektor1, vektor2,dimensir2))
			case 6:
				fmt.Print("Hasil norm vektor v1 dan v2: ")
				fmt.Println(normvektor(vektor1,dimensir2), normvektor(vektor2,dimensir2))
			case 7:
				fmt.Print("jarak antar 2 vektor adalah: ")
				fmt.Println(javektor(vektor1,vektor2,dimensir2))
			case 8:
				fmt.Println("ortogonality :", ortho(vektor1,vektor2,dimensir2))
			case 9:
				Parallelogram(vektor1,vektor2,vektor3)
			default:
				break loop1
			}
		}
		fmt.Println("keluar program")
}

func version() {
	// Fungsi-fungsi untuk mencetak teks dengan warna-warni
	red := color("\033[31m%s\033[0m")
	green := color("\033[32m%s\033[0m")
	yellow := color("\033[33m%s\033[0m")
	blue := color("\033[34m%s\033[0m")
	magenta := color("\033[35m%s\033[0m")
	cyan := color("\033[36m%s\033[0m")

	fmt.Printf(red("A ") + green("T ") + yellow("H ") + blue("I ") + magenta("L ") + cyan("A ") + "KALKULATOR VEKTOR 1.0")
}

func color(colorStr string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(colorStr, fmt.Sprint(args...))
	}
}

func menu(vektor1,vektor2, vektor3 vektor, dimensir2 bool) {
	fmt.Println("--------------------------------")
	if dimensir2 {
		fmt.Println("vektor a:", vektor1.x,vektor1.y,"vektor b:", vektor2.x,vektor2.y,"vektor c:", vektor3.x,vektor3.y)
	} else {
		fmt.Println("vektor a:", vektor1.x,vektor1.y,vektor1.z,"vektor b:", vektor2.x,vektor2.y,vektor2.z,"vektor c:", vektor3.x,vektor3.y,vektor3.z)
	}
	fmt.Println("1. Isivektor (WAJIB BANGET ISI)")
	fmt.Println("2. Penjumlahan")
	fmt.Println("3. skalar")
	fmt.Println("4. hasilKali titik")
	fmt.Println("5. hasilKali silang (dot product)")
	fmt.Println("6. norm vektor ||v||")
	fmt.Println("7. jarak antar 2 vektor")
	fmt.Println("8. cek ortogonal")
	fmt.Println("9. area Parallelogram")
	fmt.Println("--------------------------------")
}

func menuisivektor() {
	fmt.Println("--------------------------------")
	fmt.Println("1. R2 (2 dimensi)")
	fmt.Println("2. R3 (3 dimensi)")
	fmt.Println("3 or default : keluar")
	fmt.Println("--------------------------------")
}

func buatvektorr2() vektor {
	var vektordimensi vektor
	fmt.Print("Masukkan nilai x dan y, dan z. untuk z opsional agar 3 dimensi: ")
	fmt.Scan(&vektordimensi.x, &vektordimensi.y,&vektordimensi.z)
	return vektordimensi
}

func jumlahVektor(v1, v2 vektor, dimensir2 bool) {
	if dimensir2 {
		fmt.Print("(",v1.x + v2.x, v1.y + v2.y,")\n" )
	} else {
		fmt.Print("(",v1.x + v2.x, v1.y + v2.y,v2.y,v1.z + v2.z,")\n" )
	}
}

func kaliSkalarvektor(v vektor, scalar float64, dimensir2 bool) {
	if dimensir2 {
		fmt.Print("(",v.x * scalar,v.y * scalar,")\n" )
	} else {
		fmt.Print("(",v.x * scalar,v.y * scalar,v.z * scalar,")\n" )
	}
}

func kaliTitikvektor(v1, v2 vektor, dimensir2 bool) float64 {
	if dimensir2 {
		return v1.x*v2.x + v1.y*v2.y
	} else {
		return v1.x*v2.x + v1.y*v2.y + v1.z*v2.z
	}
}

func kaliSilangvektor(v1, v2 vektor, dimensir2 bool) float64 {
	if dimensir2 {
		return v1.x*v2.y - v1.y*v2.x
	} else {
		return v1.x*v2.y - v1.y*v2.x - v1.z*v2.z
	}
}

func normvektor(v vektor, dimensir2 bool) float64 {
	if dimensir2 {
		return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2))
	} else {
		return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2)+ math.Pow(v.z, 2))
	}
}

func javektor(v1, v2 vektor, dimensir2 bool) float64 {
	if dimensir2 {
		return math.Sqrt(math.Pow(v2.x -v1.x, 2) + math.Pow(v2.y - v1.y, 2))
	} else {
		return math.Sqrt(math.Pow(v2.x -v1.x, 2) + math.Pow(v2.y - v1.y, 2) + math.Pow(v2.z - v1.z, 2))
	}
}

func ortho(v1, v2 vektor, dimensir2 bool) bool {
	if dimensir2 {
		return 0 == v1.x*v2.x + v1.y*v2.y
	} else {
		return 0 == v1.x*v2.x + v1.y*v2.y  + v1.z*v2.z

	}
}

func titikakhir(v1,v2 vektor) vektor {
	var v vektor
	v.x = v2.x - v1.x
	v.y = v2.y - v1.y
	v.z = v2.z - v1.z
	return v
}

func crossproduct(ab,ac vektor) vektor {
	var c vektor
	c.x = (ab.y * ac.z - ac.y * ab.z)
	c.y = (ab.x * ac.z - ac.x * ab.z)
	c.z = (ab.x * ac.y - ac.x * ab.y)
	return c
}

func Parallelogram(v1,v2,v3 vektor,) {
	var ab, ac,c vektor
	var luas,sluas float64
	
	ab = titikakhir(v1,v2)
	fmt.Print("ab = (", ab.x,ab.y,ab.z,")\n")
	ac = titikakhir(v1,v3)
	fmt.Print("ac = (", ac.x,ac.y,ac.z,")\n")
	c = crossproduct(ab,ac)
	fmt.Print("c = ", c.x,"i - ", c.y,"j + ", c.z,"k \n")
	sluas = math.Pow(c.x,2) + math.Pow(c.y,2) + math.Pow(c.z,2)
	luas = 0.5 * math.Sqrt(sluas)
	fmt.Print("Luas = 1/2 √", sluas, " = ", luas,"\n")
}
