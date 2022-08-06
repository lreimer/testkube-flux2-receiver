NAME=testkube-flux2-receiver
VERSION=v1.0.0

.PHONY: run test clean

build:
	@goreleaser build --single-target --snapshot --rm-dist

run:
	@go run main.go

test:
	@go test -v

clean:
	@go clean
	@rm -f $(NAME)
	@rm -rf dist/

docker:
	@docker build -t lreimer/$(NAME) --build-arg version=$(VERSION) .

release:
	@goreleaser --rm-dist
