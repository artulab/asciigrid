package main

import (
	"log"
	"os"

	"github.com/artulab/asciigrid"
)

func main() {
	buf := [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	grid := asciigrid.Grid{Ncols: 3, Nrows: 2, Xllcorner: 0.1, Yllcorner: 0.2, Cellsize: 1.1,
		Nodata: -9999, Buffer: buf}

	file, err := os.Create("grid2.asc")
	if err != nil {
		log.Fatal(err)
	}

	enc := asciigrid.NewEncoder(file)
	err = enc.Encode(&grid)

	if err != nil {
		log.Fatal(err)
	}
}
