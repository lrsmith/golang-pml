package main

import (
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestLoadData(t *testing.T) {

	X, Y := importData("../data/more_pizza.txt")

	xR, xC := X.Dims()
	yR, yC := Y.Dims()

	if xR != 30 || xC != 3 {
		t.Errorf("\n Unexpected size for X: %d x %d\n", xR, xC)
	}

	if yR != 30 || yC != 1 {
		t.Errorf("\n Unexpected size for Y: %d x %d\n", yR, yC)
	}

}

func TestPredict(t *testing.T) {

	/*
	  predictedMatrix : 30 x 1 of zero
	*/
	var predictedMatrix *mat.Dense
	var expectedMatrix *mat.Dense

	X, _ := importData("../data/more_pizza.txt")

	xR, xC := X.Dims()
	weight := mat.NewDense(xC, 1, nil)

	predictedMatrix = predict(X, weight)

	expectedMatrix = mat.NewDense(xR, 1, nil)

	if !mat.Equal(mat.Matrix(expectedMatrix), mat.Matrix(predictedMatrix)) {
		t.Errorf("\nUnexpected predictedmatrix\n%v\n%v\n", expectedMatrix, predictedMatrix)
	}

}

func addOne(i, j int, v float64) float64 {
	return v + 1
}
func TestPredictOne(t *testing.T) {

	var expectedMatrix, predictedMatrix *mat.Dense

	expectedMatrix = mat.NewDense(30, 1, []float64{
		48, 22, 37, 57, 45, 15, 50, 38, 53, 18, 18, 53, 27, 44, 18,
		55, 25, 38, 39, 49, 35, 60, 39, 44, 32, 43, 24, 29, 36, 36,
	})

	X, _ := importData("../data/more_pizza.txt")

	weight := mat.NewDense(3, 1, nil)
	weight.Apply(addOne, weight)

	predictedMatrix = predict(X, weight)

	if !mat.Equal(mat.Matrix(expectedMatrix), mat.Matrix(predictedMatrix)) {
		t.Errorf("\nUnexpected predictedmatrix\n%v\n%v\n", expectedMatrix, predictedMatrix)
	}

}

func TestLoss(t *testing.T) {
	/*
	  predictMatix : 30 x 1 of zeros
	  MatrixSum : 4007
	  lossvalue : 4007 / 30 * 1
	*/
	var lossValue float64

	X, Y := importData("../data/more_pizza.txt")
	weight := mat.NewDense(3, 1, nil)

	lossValue = loss(X, Y, weight)
	if lossValue != 1333.56666666666660603369 {
		t.Errorf("\nUnexpected lossValue : %.20f\n", lossValue)
	}
}

func TestGradient(t *testing.T) {

	var tMatrix mat.Matrix

	X, Y := importData("../data/more_pizza.txt")
	weight := mat.NewDense(3, 1, nil)

	tMatrix = gradient(X, Y, weight)
	//fa := mat.Formatted(tMatrix, mat.Prefix(""), mat.Squeeze())

	t.Errorf("\n%v\n%v\n%v\n%v\n", X, Y, weight, tMatrix)
}
