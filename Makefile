EXISTING_VERSION := $(shell git describe --abbrev=0 --tags)
NEW_VERSION := $(shell echo $(EXISTING_VERSION) | awk -F. '{print ""$$1"."$$2"."$$3 + 1}')

tag_and_push:
	git tag $(NEW_VERSION)
	git push origin $(NEW_VERSION)

run-test:
	go test ./es/... -v -race -coverprofile=coverage.txt -covermode=atomic

linter:
	golangci-lint run -c .golangci.yml --timeout=5m -v --fix

unit-test-pretty:
	go test ./... -count=1 -v -json | gotestfmt

run-benchmark:
	go run ./benchmarks/run_bench.go -cooldown=15 -save=csv