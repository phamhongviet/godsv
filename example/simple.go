package main

import (
	"fmt"
	"strings"

	"github.com/phamhongviet/godsv"
)

func main() {
	var data string
	data = `1:Venice:Italy
2:Paris:South Africa
3:Prague:Czech Republic
4:Lisbon:Portugal
5:Rio De Janeiro:Brazil`

	cities := make([]godsv.Row, 5)
	dsv := godsv.New()
	for i, row := range strings.Split(data, "\n") {
		cities[i] = dsv.Unmarshal(row)
	}

	fmt.Println("Top 5 most beautiful cities:")
	for _, city := range cities {
		fmt.Printf("Number %s: %s of %s\n", city[0], city[1], city[2])
	}

	fmt.Println("\nBut Paris is a France's city...\n")
	cities[1][2] = "France"

	fmt.Println("Correct data:")

	for _, city := range cities {
		fmt.Println(dsv.Marshal(city))
	}
}
