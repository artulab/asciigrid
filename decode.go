package asciigrid

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

var (
	// ErrInvalidGrid is returned when the given raster file does not conform to the asciigrid format.
	ErrInvalidGrid = errors.New("asciigrid: input is not in a valid format")
)

// Decoder reads and decodes raster file in asciigrid format.
type Decoder struct {
	reader  io.Reader
	scanner *bufio.Scanner
	err     error
}

// NewDecoder returns a decoder reading from the reader.
func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader: reader, scanner: bufio.NewScanner(reader)}
}

// Decode decodes the grid in the reader.
func (decoder *Decoder) Decode() (*Grid, error) {
	if decoder.err != nil {
		return nil, decoder.err
	}

	var grid Grid

	i := -1
	for decoder.scanner.Scan() {
		line := strings.TrimSpace(decoder.scanner.Text())

		i++

		// first line: ncols
		if i == 0 {
			words := strings.Fields(line)
			if len(words) != 2 {
				return nil, ErrInvalidGrid
			}

			if words[0] != "ncols" {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			num, err := strconv.ParseInt(words[1], 10, 32)
			if err != nil {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			grid.Ncols = int(num)
			continue
		}

		// second line: nrows
		if i == 1 {
			words := strings.Fields(line)
			if len(words) != 2 {
				return nil, ErrInvalidGrid
			}

			if words[0] != "nrows" {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			num, err := strconv.ParseInt(words[1], 10, 32)
			if err != nil {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			grid.Nrows = int(num)
			continue
		}

		// third line: xllcorner
		if i == 2 {
			words := strings.Fields(line)
			if len(words) != 2 {
				return nil, ErrInvalidGrid
			}

			if words[0] != "xllcorner" {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			num, err := strconv.ParseFloat(words[1], 32)
			if err != nil {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			grid.Xllcorner = float32(num)
			continue
		}

		// fourth line: yllcorner
		if i == 3 {
			words := strings.Fields(line)
			if len(words) != 2 {
				return nil, ErrInvalidGrid
			}

			if words[0] != "yllcorner" {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			num, err := strconv.ParseFloat(words[1], 32)
			if err != nil {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			grid.Yllcorner = float32(num)
			continue
		}

		// fifth line: cellsize
		if i == 4 {
			words := strings.Fields(line)
			if len(words) != 2 {
				return nil, ErrInvalidGrid
			}

			if words[0] != "cellsize" {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			num, err := strconv.ParseFloat(words[1], 32)
			if err != nil {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			grid.Cellsize = float32(num)
			continue
		}

		// sixth line: NODATA_value (optional)
		if i == 5 {
			words := strings.Fields(line)
			if len(words) == 2 && words[0] == "NODATA_value" {
				num, err := strconv.ParseInt(words[1], 10, 64)
				if err != nil {
					decoder.err = ErrInvalidGrid
					return nil, decoder.err
				}
				grid.Nodata = int(num)

				continue
			} else {
				// the ESRI default is -9999
				grid.Nodata = -9999
			}
		}

		// read grid data into memory

		// create a contiguous memory for holding cell values
		buffer := make([][]int, grid.Nrows)
		pixels := make([]int, grid.Nrows*grid.Ncols)
		for j := range buffer {
			buffer[j], pixels = pixels[:grid.Ncols], pixels[grid.Ncols:]
		}
		grid.Buffer = buffer

		row_index := 0
		for {
			// parse the grid data.
			cells := strings.Fields(line)

			// make sure we have a correct number of cells in a row
			if len(cells) != int(grid.Ncols) {
				decoder.err = ErrInvalidGrid
				return nil, decoder.err
			}

			for col, cell := range cells {
				num, err := strconv.ParseInt(cell, 10, 64)
				if err != nil {
					decoder.err = ErrInvalidGrid
					return nil, decoder.err
				}
				buffer[row_index][col] = int(num)
			}

			row_index += 1

			if row_index == int(grid.Nrows) {
				// consume all newlines if exists for convenience
				for decoder.scanner.Scan() {
					line = strings.TrimSpace(decoder.scanner.Text())
					// if file has somethine other than new line, report error
					if line != "" {
						decoder.err = ErrInvalidGrid
						return nil, decoder.err
					}
				}
				break
			} else {
				if decoder.scanner.Scan() {
					line = strings.TrimSpace(decoder.scanner.Text())
				} else {
					decoder.err = ErrInvalidGrid
					return nil, decoder.err
				}
			}
		}
	}

	// given empty file?
	if i == -1 {
		decoder.err = io.EOF
		return nil, decoder.err
	}

	// encountered an error during scanning?
	err := decoder.scanner.Err()
	if err != nil {
		decoder.err = err
		return nil, err
	}

	return &grid, nil
}
