package asciigrid

import (
	"io"
)

// Encoder writes raster file in asciigrid format.
type Encoder struct {
	writer io.Writer
}

// NewEncoder creates an encoder object to write to a writer.
func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{writer: writer}
}

// Encode encodes the given grid writing to the writer.
func (encoder *Encoder) Encode(grid *Grid) error {
	_, err := encoder.writer.Write([]byte(grid.String()))
	if err != nil {
		return err
	}

	return nil
}
