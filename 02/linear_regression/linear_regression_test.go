package main

import (
	"testing"

	"github.com/gonum/floats"
)

func TestPredict(t *testing.T) {
	X, _ := importData("../../data/pizza.txt")
	predictedX := predict(X, 1)
	if !(floats.Equal(X, predictedX)) {
		t.Errorf("Predicted X does not equal X for weight = 1 bias =0\n\t%v\n\t%v\n", predictedX, X)
	}
}

func TestLost(t *testing.T) {
	X, Y := importData("../../data/pizza.txt")
	loss := loss(X, Y, 1)
	if loss != 224.86666666666667 {
		t.Errorf("Unexpected Loss : %v\n", loss)
	}
}

func TestMain(t *testing.T) {

	var weight float64

	X, Y := importData("../../data/pizza.txt")
	weight = train(X, Y, 10000, .01)
	if weight != 1.840000 {
		t.Errorf("\nUnexpected weight. >%f<\n", weight)
	}
}
