package main

import (
	"fmt"
	"math"
	"sort"
)

// Penjumlahan menggunakan variadic function
func Penjumlahan(numbers ...float64) float64 {
	total := 0.0
	for _, number := range numbers {
		total += number
	}
	return total
}

// Rata-rata menggunakan variadic function
func RataRata(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	total := Penjumlahan(numbers...)
	return total / float64(len(numbers))
}

// Median dari daftar angka
func Median(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	sort.Float64s(numbers)
	middle := len(numbers) / 2

	// Jika jumlah elemen ganjil, ambil elemen tengah
	if len(numbers)%2 != 0 {
		return numbers[middle]
	}

	// Jika genap, ambil rata-rata dari dua elemen tengah
	return (numbers[middle-1] + numbers[middle]) / 2.0
}

// Luas Lingkaran dengan jari-jari r
func LuasLingkaran(r float64) float64 {
	return math.Pi * r * r
}

// Luas Trapesium dengan sisi sejajar a, b dan tinggi t
func LuasTrapesium(a, b, t float64) float64 {
	return 0.5 * (a + b) * t
}

func main() {
	// Contoh penggunaan fungsi
	fmt.Println("Penjumlahan:", Penjumlahan(2,5,6,8,10,20)) // Hasil: 15
	fmt.Println("Rata-rata:", RataRata(2,3,4,5,6,7))     // Hasil: 3
	fmt.Println("Median:", Median(2,2,3,4,4,5,6,7,8,8,9))          // Hasil: 3
	fmt.Println("Luas Lingkaran:", LuasLingkaran(7))        // Hasil: 153.93804002589985
	fmt.Println("Luas Trapesium:", LuasTrapesium(4, 8, 12))  // Hasil: 17.5
}