BINARY_NAME=go-cmd-example
#FORMATTER=gofmt
FORMATTER=goimports

.PHONY: tools
tools:
	go install golang.org/x/tools/cmd/goimports@latest
	# TODO: use another linter
	#go get -u golang.org/x/lint/golint
	go mod tidy

.PHONY: deps
deps:
	go mod download

.PHONY: format
format:
	find ./ -name '*.go' | xargs $(FORMATTER) -l -w

.PHONY: lint
lint:
	# TODO: use another linter
	#golint -set_exit_status ./...
	go vet ./...

.PHONY: test
test: format lint
	mkdir -p reports
	go test -v -cover -coverprofile=reports/coverage.out -json ./... > reports/test.json

.PHONY: build
build:
	go build -o $(BINARY_NAME) -v

.PHONY: run
run: build
	$(BINARY_NAME)

.PHONY: release-dryrun
release-dryrun:
	goreleaser --snapshot --skip-publish --rm-dist

.PHONY: release
release:
	goreleaser --rm-dist

.PHONY: clean
clean:
	go clean
	rm -rf ./dist ./reports
	rm -f $(BINARY_NAME)
