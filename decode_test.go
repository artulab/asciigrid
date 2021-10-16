package asciigrid

import (
	"bytes"
	"testing"
)

func TestDecode(t *testing.T) {
	ebuf := [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	egrid := Grid{Ncols: 3, Nrows: 2, Xllcorner: 0.1, Yllcorner: 0.2, Cellsize: 1.1,
		Nodata: -9999, Buffer: ebuf}

	e := "ncols 3\n" +
		"nrows 2\n" +
		"xllcorner 0.100000\n" +
		"yllcorner 0.200000\n" +
		"cellsize 1.100000\n" +
		"NODATA_value -9999\n" +
		"0 1 2\n" +
		"3 4 5"

	var iobuf bytes.Buffer
	iobuf.WriteString(e)

	enc := NewDecoder(&iobuf)
	grid, err := enc.Decode()

	if err != nil {
		t.Error("Grid data decoded isn't expected: " + err.Error())
	}

	if grid.Ncols != egrid.Ncols {
		t.Error("Grid Ncols decoded isn't expected")
	}

	if grid.Nrows != egrid.Nrows {
		t.Error("Grid Nrows decoded isn't expected")
	}

	if grid.Xllcorner != egrid.Xllcorner {
		t.Error("Grid Xllcorner decoded isn't expected")
	}

	if grid.Yllcorner != egrid.Yllcorner {
		t.Error("Grid Yllcorner decoded isn't expected")
	}

	if grid.Cellsize != egrid.Cellsize {
		t.Error("Grid Cellsize decoded isn't expected")
	}

	if grid.Nodata != egrid.Nodata {
		t.Error("Grid Yllcorner decoded isn't expected")
	}

	for y, row := range egrid.Buffer {
		if len(row) != len(grid.Buffer[y]) {
			t.Error("Grid Buffer row length isn't expected")
		}

		for x, val := range row {
			if grid.Buffer[y][x] != val {
				t.Error("Grid Buffer cell value isn't expected")
			}
		}
	}
}
