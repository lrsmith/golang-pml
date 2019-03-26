package main

import (
	"image/color"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func readData(file string) plotter.XYs {

	fh, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(fh), "\n")
	points := make(plotter.XYs, len(lines)-1)

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
		points[index-1].X, _ = strconv.ParseFloat(fields[0], 64)
		points[index-1].Y, _ = strconv.ParseFloat(fields[1], 64)

	}

	return points
}

func main() {

	scatterData := readData("data/pizza.txt")
	plot, err := plot.New()
	if err != nil {
		panic(err)
	}

	plot.Title.Text = "Reservations vs Pizzas"
	plot.X.Label.Text = "Reservations"
	plot.Y.Label.Text = "Pizzas"

	predictiveLine := plotter.NewFunction(func(x float64) float64 { return (x*1.1 + 12.930) })
	predictiveLine.Color = color.RGBA{B: 255, A: 255}

	// Draw a grid behind the data
	plot.Add(plotter.NewGrid())

	// Make a scatter plotter and set its style.
	scatter, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	scatter.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	plot.Add(scatter, predictiveLine)
	plot.X.Min = 0
	plot.X.Max = 30
	plot.Y.Min = 0
	plot.Y.Max = 55

	// Save the plot to a PNG file.
	if err := plot.Save(4*vg.Inch, 4*vg.Inch, "graphs/plot_linear_regression_with_bias.png"); err != nil {
		panic(err)
	}

}
