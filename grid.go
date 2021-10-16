package asciigrid

import (
	"fmt"
	"strconv"
	"strings"
)

// Grid struct representing the AsciiGrid data.
type Grid struct {
	Ncols     int
	Nrows     int
	Xllcorner float32
	Yllcorner float32
	Cellsize  float32
	Nodata    int
	Buffer    [][]int
}

// String returns the grid as string.
func (grid *Grid) String() string {
	str := fmt.Sprintf("ncols\t\t\t%d\n", grid.Ncols)
	str += fmt.Sprintf("nrows\t\t\t%d\n", grid.Nrows)
	str += fmt.Sprintf("xllcorner\t\t%f\n", grid.Xllcorner)
	str += fmt.Sprintf("yllcorner\t\t%f\n", grid.Yllcorner)
	str += fmt.Sprintf("cellsize\t\t%f\n", grid.Cellsize)
	str += fmt.Sprintf("NODATA_value\t%d\n", grid.Nodata)

	var sb strings.Builder
	for i, row := range grid.Buffer {
		for j, col := range row {
			sb.WriteString(strconv.Itoa(col))
			if j != grid.Ncols-1 {
				sb.WriteString(" ")
			}
		}
		if i != int(grid.Nrows-1) {
			sb.WriteString("\n")
		}
	}

	str += sb.String()

	return str
}
