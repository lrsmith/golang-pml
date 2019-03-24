
linear_regression:
	go build -o ./bin/linear_regression ./02/linear_regression.go

plot:
	go build -o ./bin/plot ./02/plot.go

clean:
	rm ./bin/*
	rm ./graphs/*
