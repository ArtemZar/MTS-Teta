default: lint

build:
	go build -v ./cmd/

lint:
	golangci-lint run -v ./...

test:
	go test -v -rase ./...

clean:
	@echo $(CLEANUP)
	$(foreach f,$(CLEANUP),rm -rf $(f);)

.PHONY: build lint test clean