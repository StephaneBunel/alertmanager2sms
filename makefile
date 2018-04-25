SRC           = $(shell find . -name '*.go' -not -path "./vendor/*" )
VERSION       = $(shell git describe --always --tags)
COMMIT        = $(shell git rev-parse --short=8 HEAD)
GO           := go
LDFLAGS      := -s -X cmd.am2sms.VersionBuild=$(shell date --iso=s)
GOLINT       := golint
BINARY       := am2sms

define formatme
echo "-- formating…"
$(GO) fmt ./...
echo "-- Ok"
endef

define buildme
echo "-- building…"
$(GO) build -tags release -ldflags='$(LDFLAGS)' -o $(BINARY) main.go
echo "-- Ok"
endef

define lintme
echo "-- linting…"
$(GOLINT) ./pkg/... ./cmd/...
echo "-- Ok"
endef

define testme
echo "-- testing…"
$(GO) test -race -cover -timeout 60s ./pkg/... ./cmd/...
echo "-- Ok"
endef

$(BINARY): $(SRC)
	@$(formatme)
	@$(buildme)
	@$(lintme)
	@$(testme)

.PHONY: format
format:
	@$(formatme)

.PHONY: build
build:
	@$(buildme)

.PHONY: lint
lint:
	@$(lintme)

.PHONY: test
test:
	@$(testme)

.PHONY: vendor
vendor:
	@echo "-- syncing vendor dependencies…"
	@govendor sync -v

.PHONY: clean
clean:
	@echo "-- cleaning…"
	@rm -fv $(BINARY)
