GO := go
NAME := thead
VERSION := 1.0.0
DIST := $(NAME)-$(VERSION)

all: test build

setup: update_version

define _createDist
	mkdir -p dist/$(1)_$(2)/$(DIST)
	GOOS=$1 GOARCH=$2 go build -o dist/$(1)_$(2)/$(DIST)/$(NAME)$(3) cmd/$(NAME)/*.go
	cp -r README.md LICENSE completions dist/$(1)_$(2)/$(DIST)
	cp -r docs/public dist/$(1)_$(2)/$(DIST)/docs
	tar cfz dist/$(DIST)_$(1)_$(2).tar.gz -C dist/$(1)_$(2) $(DIST)
endef

test: setup
	$(GO) test -covermode=count -coverprofile=coverage.out $$(go list ./...)

build: setup
	$(GO) build -o $(NAME) cmd/thead/*.go

docs: docs/public

docs/public:
	(cd docs; make)

dist: build docs
	@$(call _createDist,darwin,amd64)
	@$(call _createDist,darwin,arm64)
	@$(call _createDist,windows,amd64,.exe)
	@$(call _createDist,windows,386,.exe)
	@$(call _createDist,linux,amd64)
	@$(call _createDist,linux,386)

clean:
	$(GO) clean
	rm -rf $(NAME) dist

define _update_docker
	(sed -e '$$d' Dockerfile ; echo $(1)) > a
	mv a Dockerfile
endef
