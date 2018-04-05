PROJ = github.com/StephaneBunel/alertmanager2sms
EXE = am2sms
SRC = cmd/**/*.go \
      pkg/**/*.go \

LDFLAGS = -s
GOBUILD = go build -v --ldflags="$(LDFLAGS)"

all: $(EXE)

$(EXE): $(SRC)
	$(GOBUILD) $(PROJ)/cmd/am2sms

govendor:
	@govendor sync -v

clean:
	@rm -fv $(EXE)
	@rm -rfv ${GOPATH}/pkg/*/$(PROJ)/pkg/*

tests:
	go test -cover -timeout 60s $(PROJ)/pkg/...
