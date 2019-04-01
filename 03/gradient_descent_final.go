package main

// Switch from gonum/mat to using []float64 so can leverage gonum/floats and gonum/stat

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
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

func predict(X []float64, weight, bias float64) []float64 {

	// Passing slice so passing reference, need to copy to avoid
	//  modifying the data structure backing the slice.
	predicted := make([]float64, len(X))
	copy(predicted, X)

	floats.Scale(weight, predicted)
	floats.AddConst(bias, predicted)

	return predicted
}

func loss(X, Y []float64, weight, bias float64) float64 {

	loss := predict(X, weight, bias)
	floats.Sub(loss, Y)
	floats.Mul(loss, loss)

	return stat.Mean(loss, nil)
}

func gradient(X, Y []float64, weight, bias float64) (float64, float64) {

	var weightGradient, biasGradient float64

	// predict(X,w,b)
	gradientArray := predict(X, weight, bias)
	// predict(X,w,b) - Y
	floats.Sub(gradientArray, Y)
	// 2 * predict(X,w,b) - Y
	floats.Scale(2.0, gradientArray)

	biasGradient = stat.Mean(gradientArray, nil)

	// X * 2 * predict(X,w,b) - Y
	floats.Mul(gradientArray, X)

	weightGradient = stat.Mean(gradientArray, nil)

	return weightGradient, biasGradient
}

func train(X, Y []float64, iterations int, learningRate float64) (float64, float64) {

	var weight, bias float64

	for i := 0; i < iterations; i++ {
		fmt.Printf("Iteration %d => Loss %.10f\n", i, loss(X, Y, weight, bias))
		weightGradient, biasGradient := gradient(X, Y, weight, bias)
		weight -= weightGradient * learningRate
		bias -= biasGradient * learningRate
	}

	return weight, bias
}

func main() {

	X, Y := importData("data/pizza.txt")
	weight, bias := train(X, Y, 20000, 0.001)
	fmt.Printf("\nweight=%.10f, bias=%.10f", weight, bias)
	fmt.Printf("Prediction: x=%d => y=%.2f\n", 20, 20.0*weight+bias)
}
