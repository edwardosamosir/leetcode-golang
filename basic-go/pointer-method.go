package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}
func main() {
	edwardo := Man{"Edwardo"}
	edwardo.Married()

	fmt.Println(edwardo.Name)
}
