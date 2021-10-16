# asciigrid
![Version](https://img.shields.io/badge/version-v1.0.0-blue.svg?cacheSeconds=2592000)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

> asciigrid is a Go package that implements decoder and encoder for the Esri ASCII grid format, also known as ARC/INFO ASCII GRID.

## Install

```sh
go get -v github.com/artulab/asciigrid
```

## Usage

Import the package:
```go
import "github.com/artulab/asciigrid"
```

### Decoder

Construct a decoder object out of any I/O stream implementing Go's Reader interface:
```go
file, err := os.Open("grid.asc")
if err != nil {
	log.Fatal(err)
}

enc := asciigrid.NewDecoder(file)
```

Read geographic header data and grid values into memory:

```go
grid, err := enc.Decode()

fmt.Printf("ncols: %v\n", grid.Ncols)
fmt.Printf("nrows: %v\n", grid.Nrows)
fmt.Printf("nodata: %v\n", grid.Nodata)

for j := 0; j < grid.Nrows; j++ {
	for i := 0; i < grid.Ncols; i++ {
		fmt.Printf("%d ", grid.Buffer[j][i])
	}
	fmt.Println()
}
```

### Encoder

Construct an encoder object out of an I/O stream object representing files, memory buffers etc.:

```go
file, err := os.Create("grid.asc")
if err != nil {
	log.Fatal(err)
}

enc := asciigrid.NewEncoder(file)
```

Write the grid data in memory into an I/O stream implementing Go's Writer interface:

```go
buf := [][]int{
	{0, 1, 2},
	{3, 4, 5},
}

grid := asciigrid.Grid{Ncols: 3, Nrows: 2, Xllcorner: 0.1, Yllcorner: 0.2, Cellsize: 1, Nodata: -9999, Buffer: buf}

enc.Encode(&grid)
```

See [examples](https://github.com/artulab/asciigrid/examples) for more information.

## Run tests

```sh
go test
```

## Author

👤 **Ahmet Artu Yildirim**

* Website: https://www.artulab.com
* E-Mail: ahmet@artulab.com

## 🤝 Contributing

Contributions, issues and feature requests are welcome!

Feel free to check [issues page](https://github.com/artulab/asciigrid/issues). 

## Show your support

Give a ⭐️ if this project helped you!


## 📝 License

Copyright © 2021 [Ahmet Artu Yildirim](https://www.artulab.com).

This project is [MIT](https://opensource.org/licenses/MIT) licensed.
