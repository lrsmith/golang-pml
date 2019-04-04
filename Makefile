
# Chapter 02

linear_regression:								# Refactored to []float64
	go build -o ./bin/linear_regression \
		./02/linear_regression/linear_regression.go

linear_regression_with_bias:
	go build -o ./bin/linear_regression_with_bias \
		./02/linear_regression_with_bias/linear_regression_with_bias.go

# Refactor and add plot capabilties to above
#plot:
#	go build -o ./bin/plot ./02/plot.go
#plot_linear_regression:
#	go build -o ./bin/plot_linear_regression ./02/plot_linear_regression.go
#plot_linear_regression_with_bias:
#	go build -o ./bin/plot_linear_regression_with_bias ./02/plot_linear_regression_with_bias.go

# Chapter 03

gradient_descent_without_bias:
	go build -o ./bin/gradient_descent_without_bias ./03/gradient_descent_without_bias.go

gradient_descent_final:
	go build -o ./bin/gradient_descent_final ./03/gradient_descent_final.go

clean:
	rm ./bin/*
	rm ./graphs/*
