package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) { // Pointer Address agar data dapat berubah
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Bandung", "Jawa Barat", ""}
	ChangeCountryToIndonesia(&address)
	fmt.Println(address)
}
