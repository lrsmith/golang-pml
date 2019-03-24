package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func matPrint(X mat.Matrix) {
	fa := mat.Formatted(X, mat.Prefix(""), mat.Squeeze())
	fmt.Printf("%v\n", fa)
}

func import_data(file string) (*mat.VecDense, *mat.VecDense) {

	fh, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fh), "\n")

	X := mat.NewVecDense(len(lines)-1, nil)
	Y := mat.NewVecDense(len(lines)-1, nil)

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

		x, _ := strconv.ParseFloat(fields[0], 64)
		y, _ := strconv.ParseFloat(fields[1], 64)

		X.SetVec(index-1, x)
		Y.SetVec(index-1, y)

	}

	return X, Y
}

func train(X []float64, Y []float64, iterations int, lr int) int {

	var w int
	return w
}

func predict(X []float64, w int) {

	return // X* w
}

func loss(X []float64, Y []float64, iterations int, lr int) {
	return
}

func main() {

	X, Y := import_data("data/pizza.txt")
	matPrint(X)
	matPrint(Y)
}
