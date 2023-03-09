package main

import (
	"fmt"
	"project-go/modul"
)

func Add(angka1 float32, angka2 float32) float32 {
	return angka1 / angka2
}

func perbandingan() string {
	var value = 2 + 2
	var isEqual = value == 4
	a := fmt.Sprintf("hasil: %d %t", value, isEqual)
	return a
}

func logika() {

}

func main() {
	/*var firstName string = "Nandhika"

	var lastName string
	lastName = "Pratama"

	fmt.Printf("halo %s %s/n", firstName, lastName)*/
	hitung := Add(1, 3)
	println(hitung)

	hasilBanding := perbandingan()
	println(hasilBanding)
	modul.Hello()
}
