package main

import (
	"io/ioutil"
	"log"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func importData(file string) (*mat.Dense, *mat.Dense) {

	var X, Y []float64

	// Assumes we know the size of the data set
	xMatrix := mat.NewDense(30, 3, nil)

	fh, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fh), "\n")

	for index, line := range lines {
		// Skip header
		if index == 0 {
			continue
		}
		// Skip empty line
		if len(line) == 0 {
			continue
		}
		fields := strings.Fields(line)

	}

	return xMatrix, yMatrix
}

func main() {

	importData("data/more_pizza.txt")
}
