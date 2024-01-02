MAKEFLAGS += --always-run --always-make --warn-undefined-variables # --silent --no-print-directory

# SOURCE_DATE_EPOCH ?= $(shell date +%s)

# export GO111MODULE=on
# unexport GOPATH

# export LDFLAGS := -extldflags '$(LDFLAGS)'
# export GCFLAGS := all=-trimpath '$(PWD)'
# export ASMFLAGS := all=-trimpath '$(PWD)'

run:
	go run main.go ${ARGS}

build:
	$(GO_RUN_BUILD)

fmt:
	go fmt ./...

install:
	go install
