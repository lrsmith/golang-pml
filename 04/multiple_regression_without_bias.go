package main

/*

ŷ = x1 * w1 + x2 * w2 + x3 * w3 + ...

Data Set
  X : 30 x 3
  Y : 30 x 1
*/

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func importData(file string) (*mat.Dense, *mat.Dense) {

	var rowCount, colCount int
	var yArray, rowMajorArray []float64

	fh, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(fh), "\n")

	for index, line := range lines {

		fields := strings.Fields(line)

		// Skip the header, Count Columns
		if index == 0 {
			colCount = len(fields)
			continue
		}
		// Skip blank lines
		if len(line) == 0 {
			continue
		}

		// Append to Row Major Array for creating X Matrix
		for _, field := range fields[:colCount-1] {
			tmpField, _ := strconv.ParseFloat(field, 64)
			rowMajorArray = append(rowMajorArray, tmpField)
		}
		rowCount++

		// Append to array for creating Y
		y, _ := strconv.ParseFloat(fields[len(fields)-1], 64)
		yArray = append(yArray, y)
	}

	// Create Matrices
	X := mat.NewDense(rowCount, colCount-1, rowMajorArray)
	Y := mat.NewDense(rowCount, 1, yArray)

	return X, Y
}

func predict(X, weight *mat.Dense) *mat.Dense {

	/*
		X         : 30 x 3
		weight    :  3 x 1
		mulMatrix : 30 x 1
	*/

	var mulMatrix mat.Dense

	mulMatrix.Mul(X, weight)

	return &mulMatrix
}

func loss(X, Y, weight *mat.Dense) float64 {

	/*
		X       : 30 x 3
		Y       : 30 x 1
		weight  :  3 x 1
		lMatrix : 30 x 1
	*/
	//return stat.Mean(col, nil)

	var lMatrix *mat.Dense

	lMatrix = predict(X, weight)
	lMatrix.Sub(lMatrix, Y)
	lMatrix.Apply(sqrElem, lMatrix)

	r, c := lMatrix.Dims()
	return mat.Sum(lMatrix) / float64(r*c)
}

func sqrElem(i, j int, v float64) float64 {
	return v * v
}

func gradient(X, Y, weight *mat.Dense) mat.Matrix {

	/*
	   Must return a 3 x 1 matrix.
	   2 * np.matmul(X.T, (predict(X, w) - Y)) / X.shape[0]

	   X     : 30 x 3
	   X.T() : 3 x 30
	*/
	/*

		tmpR, _ := X.Dims()

		mulMatrix.Scale(float64(1/tmpR), &mulMatrix)

		return mulMatrix
	*/

	var tMatrix *mat.Dense
	var iMatrix mat.Dense
	tMatrix = predict(X, weight)
	tMatrix.Sub(tMatrix, Y) // 30 x 1 matrix

	gMatrix := X.T()
	iMatrix.Mul(gMatrix, tMatrix)
	iMatrix.Scale(2, &iMatrix)

	tmpR, _ := X.Dims()
	iMatrix.Scale(float64(1/tmpR), &iMatrix)

	return mat.Matrix(&iMatrix)
}

/*
func train(X, Y *mat.Dense, iterations int, learningRate float64) *mat.Dense {

	_, xC := X.Dims()
	weight := mat.NewDense(xC, 1, nil)

	for i := 0; i < iterations; i++ {
		var tmpMatrix mat.Dense
		loss(X, Y, weight)
		fmt.Printf("Iteration %d => Loss %.10f\n", i, loss(X, Y, weight))
		tmpMatrix = gradient(X, Y, weight)
		tmpMatrix.Scale(learningRate, &tmpMatrix)
		weight.Sub(weight, &tmpMatrix)
	}

	return weight
}
*/

func main() {

	importData("../data/more_pizza.txt")

	//	train(X, Y, 500000, 0.001)
}
