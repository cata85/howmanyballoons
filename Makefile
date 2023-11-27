build:
	go build -o bin/app

run dev: build
	./bin/app DEV

run prod: build
	./bin/app PROD

test:
	go test -v ./... -count=1