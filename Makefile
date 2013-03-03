benchmike:
	mkdir -p bin
	go build -o bin/benchmike benchmike.go

clean:
	rm -rf bin
