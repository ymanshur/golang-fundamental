package main

import "fmt"

// construct
type BangunDatar interface {
	HitungLuas() int
}

// Persegi implement BangunDatar?
type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	// bangunDatar := Persegi{Sisi: 4}
	persegiPanjang := PersegiPanjang{Panjang: 4, Lebar: 5}
	luas := SeberapaLuas(persegiPanjang)
	// luas := SeberapaLuas2(bangunDatar)
	fmt.Println("Luas Persegi Panjang:", luas)

	persegi := Persegi{Sisi: 4}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas Persegi:", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}

// func SeberapaLuas2(bangunDatar PersegiPanjang) int {
// 	return bangunDatar.HitungLuas()
// }
