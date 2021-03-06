package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/gonum/floats"
	"github.com/gonum/stat"
)

func importData(file string) ([]float64, []float64) {

	var X, Y []float64

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

		x, _ := strconv.ParseFloat(fields[0], 64)
		y, _ := strconv.ParseFloat(fields[1], 64)

		X = append(X, x)
		Y = append(Y, y)

	}

	return X, Y
}

func predict(X []float64, weight float64) []float64 {

	// Passing slice so passing reference, need to copy to avoid
	//  modifying the data structure backing the slice.
	predicted := make([]float64, len(X))
	copy(predicted, X)

	floats.Scale(weight, predicted)

	return predicted
}

func loss(X, Y []float64, weight float64) float64 {

	loss := predict(X, weight)
	floats.Sub(loss, Y)
	floats.Mul(loss, loss)

	return stat.Mean(loss, nil)
}

func train(X, Y []float64, iterations int, learningRate float64) float64 {

	weight := 0.0
	for i := 0; i < iterations; i++ {
		currentLoss := loss(X, Y, weight)
		//fmt.Printf("Iteration %4d => Loss: %.6f\n", i, currentLoss)

		if loss(X, Y, weight+learningRate) < currentLoss {
			weight += learningRate
		} else if loss(X, Y, weight-learningRate) < currentLoss {
			weight -= learningRate
		} else {
			return weight
		}
	}
	log.Fatal("Couldn't Converge")
	return -1
}

func main() {

	X, Y := importData("data/pizza.txt")
	weight := train(X, Y, 10000, .01)
	fmt.Printf("\nw=%f\n", weight)
	fmt.Printf("Prediction: x=%d => y=%.2f\n", 20, 20.0*weight)
}
