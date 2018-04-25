SRC           = $(shell find . -name '*.go' -not -path "./vendor/*" )
VERSION       = $(shell git describe --always --tags)
COMMIT        = $(shell git rev-parse --short=8 HEAD)
PACKAGE      := github.com/StephaneBunel/alertmanager2sms
GO           := go
LDFLAGS      := -s -X "$(PACKAGE)/cmd/am2sms.VersionBuild=$(shell date --iso=s)"
GOLINT       := golint
BINARY       := am2sms
GOBUILD      := $(GO) build -tags release -ldflags='$(LDFLAGS)'

define formatme
echo -n "-- formating… "
$(GO) fmt ./...
echo "Ok"
endef

define buildme
echo -n "-- building… "
$(GOBUILD) -o $(BINARY) main.go
echo "Ok"
endef

define lintme
echo -n "-- linting… "
$(GOLINT) ./pkg/... ./cmd/...
echo "Ok"
endef

define testme
echo "-- testing…"
$(GO) test -race -cover -timeout 60s ./pkg/... ./cmd/...
echo "Ok"
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

.PHONY: release
release:
	@echo "-- building release… "
	@rm -rf release
	@mkdir release
	@for os in linux darwin freebsd windows; do \
		echo -n "        $${os}/amd64… "; \
		GOOS=$${os} GOARCH=amd64 $(GOBUILD) -o release/$(BINARY)-$${os}-amd64 main.go; \
		echo "Ok"; \
	done
	@echo "Ok"

.PHONY: vendor
vendor:
	@echo "-- syncing vendor dependencies…"
	@govendor sync -v

.PHONY: clean
clean:
	@echo -n "-- cleaning… "
	@rm -f $(BINARY)
	@rm -rf release
	@echo "Ok"
