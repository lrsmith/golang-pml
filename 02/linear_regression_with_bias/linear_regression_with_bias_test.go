package main

import (
	"testing"

	"github.com/gonum/floats"
)

func TestPredict(t *testing.T) {
	X, _ := importData("../../data/pizza.txt")
	predictedX := predict(X, 1, 0)
	if !(floats.Equal(X, predictedX)) {
		t.Errorf("Predicted X does not equal X for weight = 1 bias =0\n\t%v\n\t%v\n", predictedX, X)
	}
}

func TestLost(t *testing.T) {
	X, Y := importData("../../data/pizza.txt")
	loss := loss(X, Y, 1, 0)
	if loss != 224.86666666666667 {
		t.Errorf("Unexpected Loss : %v\n", loss)
	}
}

func TestMain(t *testing.T) {

	var weight, bias float64

	X, Y := importData("../../data/pizza.txt")
	weight, bias = train(X, Y, 10000, .01)
	//	t.Errorf("\n%f : %f", weight, bias)
	if weight != 1.100000 ||
		bias != 12.929999 {
		t.Errorf("\nUnexpected weight. >%v< or\nbias. >%.20f<\n", weight, bias)
	}
}
