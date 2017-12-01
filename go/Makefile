all: deps lint test build

deps:
	./tools/deps.sh

lint:
	./tools/lint.sh

test:
	go test $(glide novendor)

build:
	$(GO) build ./cmd/...