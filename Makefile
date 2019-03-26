
linear_regression:
	go build -o ./bin/linear_regression ./02/linear_regression.go

plot:
	go build -o ./bin/plot ./02/plot.go

plot_linear_regression:
	go build -o ./bin/plot_linear_regression ./02/plot_linear_regression.go

clean:
	rm ./bin/*
	rm ./graphs/*
