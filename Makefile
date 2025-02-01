EXISTING_VERSION := $(shell git describe --abbrev=0 --tags)
NEW_VERSION := $(shell echo $(EXISTING_VERSION) | awk -F. '{print ""$$1"."$$2"."$$3 + 1}')

init:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2
	go install golang.org/x/tools/go/analysis/passes/fieldalignment/cmd/fieldalignment@v0.29.0
	go install github.com/gotesttools/gotestfmt/v2/cmd/gotestfmt@latest
	go install github.com/GokselKUCUKSAHIN/go-run-bench@v1.0.1

tag_and_push:
	git tag $(NEW_VERSION)
	git push origin $(NEW_VERSION)

run-test:
	go test ./es/... -v -race -coverprofile=coverage.txt -covermode=atomic

linter:
	golangci-lint run -c .golangci.yml --timeout=5m -v --fix

fixfieldalignment:
	fieldalignment --fix ./...

unit-test-pretty:
	go test ./... -count=1 -v -json | gotestfmt

run-benchmark:
	go-run-bench -cooldown=15 -benchmem=true -save=csv