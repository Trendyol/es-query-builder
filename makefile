EXISTING_VERSION := $(shell git describe --abbrev=0 --tags)
NEW_VERSION := $(shell echo $(EXISTING_VERSION) | awk -F. '{print "v"$$1"."$$2"."$$3 + 1}')

tag_and_push:
	git tag $(NEW_VERSION)
	git push origin $(NEW_VERSION)

test:
	go test -v ./...