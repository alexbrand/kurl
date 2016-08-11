ifeq ($(origin VERSION), undefined)
  VERSION=$(git rev-parse --short HEAD)
endif
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)
REPOPATH = alexbrand/kurl

build: vendor
	go build -o bin/kurl -ldflags "-X $(REPOPATH).Version=$(VERSION)"

test: bin/glide
	go test $(shell ./bin/glide novendor)

vet: bin/glide
	go vet $(shell ./bin/glide novendor)

fmt: bin/glide
	go fmt $(shell ./bin/glide novendor)

vendor: bin/glide
	./bin/glide install

bin/glide:
	@echo "Downloading glide"
	mkdir -p bin
	curl -L https://github.com/Masterminds/glide/releases/download/v0.11.1/glide-v0.11.1-$(GOOS)-$(GOARCH).tar.gz | tar -xz -C bin
	mv bin/$(GOOS)-$(GOARCH)/glide bin/glide
	rm -r bin/$(GOOS)-$(GOARCH)