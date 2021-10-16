package main

import (
	"fmt"
	"log"
	"os"

	"github.com/artulab/asciigrid"
)

func main() {
	file, err := os.Open("grid.asc")
	if err != nil {
		log.Fatal(err)
	}

	enc := asciigrid.NewDecoder(file)
	grid, err := enc.Decode()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ncols: %v\n", grid.Ncols)
	fmt.Printf("nrows: %v\n", grid.Nrows)
	fmt.Printf("nodata: %v\n", grid.Nodata)

	for j := 0; j < grid.Nrows; j++ {
		for i := 0; i < grid.Ncols; i++ {
			fmt.Printf("%d ", grid.Buffer[j][i])
		}
		fmt.Println()
	}
}
