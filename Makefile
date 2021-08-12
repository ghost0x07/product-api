tag=product-api
version=1.0.0

build:
	go get -d ./...
	go build -v ./cmd/...

buildImage:
	docker build -t $(tag):$(version) .
	docker tag $(tag):$(version) $(tag):latest