run:
	go run .

build:
	go get .
	mkdir -p dist
	GOOS=linux GOARCH=amd64 go build -o dist/speedtester
	mkdir -p ../deploy/dist
	cd dist && zip ../../deploy/dist/speedtester.zip speedtester