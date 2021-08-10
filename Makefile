tag=product-api
version=0.0.1

build:
	go get -d ./...
	go build -v ./cmd/...

buildImage:
	docker build -t $(tag):$(version) .
	docker tag $(tag):$(version) $(tag):latest