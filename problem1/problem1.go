package main

import "fmt"

func KonversiNilai(nilai int) string {
	// your code here
	var konversi string

	if nilai >= 80 && nilai <= 100 {
		konversi = "A"
	} else if nilai >= 65 && nilai <= 79 {
		konversi = "B+"
	} else if nilai >= 50 && nilai <= 64 {
		konversi = "B"
	} else if nilai >= 35 && nilai <= 49 {
		konversi = "C"
	} else if nilai >= 0 && nilai <= 34 {
		konversi = "D"
	} else if nilai > 100 && nilai < 0 {
		konversi = "Kesalahan"
	}

	return konversi
}

func main() {
	var nilai int = 80

	fmt.Println(KonversiNilai(nilai))
}
