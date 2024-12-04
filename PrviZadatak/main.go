package main

import "fmt"

func proizvod(products map[string]int){
	var najskuplji string
	najvecaCijena:= 0

	for naziv, cijena := range products{
		if cijena>= najvecaCijena{
			najvecaCijena=cijena
			najskuplji=naziv
		}

		if cijena >= 150{
			fmt.Printf("Proizvod: %s, cjena: %d\n", naziv, cijena)
		}
	}

	fmt.Printf("Najskuplji proizvod je %s sa cijenom od %d\n",najskuplji,najvecaCijena)
}

func main() {
	products := map[string]int{
		"TV": 100,
		"Laptop": 300,
		"MObiutel": 150,
		"Ruter": 90,
		"Tipkovnica": 15,
		"Mis": 5,
	}
	proizvod(products)
}
