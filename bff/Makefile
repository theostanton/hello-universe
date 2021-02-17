run:
	go run .

build:
	go get .
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/bff
	mkdir -p ../deploy/dist
	cd dist && zip ../../deploy/dist/bff.zip bff