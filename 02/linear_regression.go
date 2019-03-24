package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func import_data(file string) ([]float64, []float64) {

	fh, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fh), "\n")

	X := make([]float64, len(lines)-1)
	Y := make([]float64, len(lines)-1)

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
		X[index-1], _ = strconv.ParseFloat(fields[0], 64)
		Y[index-1], _ = strconv.ParseFloat(fields[1], 64)

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
	fmt.Printf("%v\n", X)
	fmt.Printf("%v\n", Y)
}
