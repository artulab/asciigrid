package asciigrid

import (
	"bytes"
	"testing"
)

func TestEncode(t *testing.T) {
	buf := [][]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	grid := Grid{Ncols: 3, Nrows: 2, Xllcorner: 0.1, Yllcorner: 0.2, Cellsize: 1,
		Nodata: -9999, Buffer: buf}

	e := "ncols 3\n" +
		"nrows 2\n" +
		"xllcorner 0.100000\n" +
		"yllcorner 0.200000\n" +
		"cellsize 1.000000\n" +
		"NODATA_value -9999\n" +
		"0 1 2\n" +
		"3 4 5"

	var iobuf bytes.Buffer
	enc := NewEncoder(&iobuf)
	err := enc.Encode(&grid)

	if err == nil {
		t.Error(err)
	}

	if iobuf.String() != e {
		t.Error("Grid string returned isn't expected")
	}
}
