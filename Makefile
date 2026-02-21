APP=errorframework

.PHONY: test lint build clean release

test:
	go test ./...

lint:
	golangci-lint run

build:
	go build ./...

clean:
	rm -rf bin/

release:
	./scripts/release.sh