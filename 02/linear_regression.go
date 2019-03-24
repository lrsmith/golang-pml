package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func vecPrint(vector mat.Matrix) {
	fa := mat.Formatted(vector.T(), mat.Prefix(""), mat.Squeeze())
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

func predict(xVector *mat.VecDense, weight float64) *mat.VecDense {

	tmpVector := mat.NewVecDense(xVector.Len(), nil)
	tmpVector.ScaleVec(weight, xVector)

	// vecPrint(xVector)
	// vecPrint(tmpVector)

	return tmpVector
}

func loss(xVector *mat.VecDense, yVector *mat.VecDense, weight float64) float64 {

	tmpVector := predict(xVector, weight)
	tmpVector.SubVec(tmpVector, yVector)       // vecPrint(tmpVector)
	tmpVector.MulElemVec(tmpVector, tmpVector) // vecPrint(tmpVector)

	sum := 0.0

	for i := 0; i < tmpVector.Len(); i++ {
		sum += tmpVector.AtVec(i)
		//fmt.Println(sum)
	}

	return sum / float64(tmpVector.Len())
}

func train(xVector *mat.VecDense, yVector *mat.VecDense, iterations int, learningRate float64) float64 {

	weight := 0.0
	for i := 0; i < iterations; i++ {
		currentLoss := loss(xVector, yVector, weight)
		fmt.Printf("Iteration %4d => Loss: %.6f\n", i, currentLoss)

		if loss(xVector, yVector, weight+learningRate) < currentLoss {
			weight += learningRate
		} else if loss(xVector, yVector, weight-learningRate) < currentLoss {
			weight -= learningRate
		} else {
			return weight
		}
	}
	log.Fatal("Couldn't Converge")
	return -1
}

func main() {

	X, Y := import_data("data/pizza.txt")
	weight := train(X, Y, 10000, .01)

	fmt.Printf("\nw=%.3f\n", weight)
}
