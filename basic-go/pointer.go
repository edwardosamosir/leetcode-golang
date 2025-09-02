package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value) :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	address1 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address2 := address1
	//Karena di atas pass by value maka data di address1 tidak berubah ketika data di address2 diubah
	address2.City = "Bandung"
	fmt.Println(address1)
	fmt.Println(address2)

	//Pointer concept
	address3 := Address{"Bekasi", "Jawa Barat", "Indonesia"}
	address4 := &address3 //Pointer using & sign

	//Karena di atas pass by reference maka data di address3 berubah ketika data di address2 diubah
	address3.City = "Bandung"
	fmt.Println(address3)
	fmt.Println(*address4)

	// or another syntax of pointer
	var address5 Address = Address{"Bekasi", "Jawa Barat", "Indonesia"}
	var address6 *Address = &address3
	fmt.Println(address5)
	fmt.Println(*address6)
}
