build:
	go build -o bin/app

run dev: build
	./bin/app DEV

test:
	go test -v ./... -count=1