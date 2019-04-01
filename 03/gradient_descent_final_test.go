package main

import (
	"testing"

	"github.com/gonum/floats"
)

func TestPredict(t *testing.T) {

	X, _ := importData("../data/pizza.txt")
	predictedX := predict(X, 1, 0)
	if !(floats.Equal(X, predictedX)) {
		t.Errorf("Predicted X does not equal X for weight = 1 bias =0\n\t%v\n\t%v\n", predictedX, X)
	}
}

func TestLoss(t *testing.T) {
	X, Y := importData("../data/pizza.txt")
	loss := loss(X, Y, 1, 0)
	if loss != 224.86666666666667 {
		t.Errorf("Unexpected Loss : %v\n", loss)
	}
}

func TestGradient(t *testing.T) {
	X, Y := importData("../data/pizza.txt")
	weight, bias := gradient(X, Y, 1, 0)
	if weight != -369.200000 &&
		bias != -28.400000 {
		t.Errorf("Unexpected weight and bias. %f %f\n", weight, bias)
	}
}
