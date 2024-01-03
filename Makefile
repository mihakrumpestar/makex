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

test:
	cd tests/single_target && go run ../../main.go deploy
	# cd tests/single_target && go run ../../main.go destroy

	# cd tests/multi_target && go run ../../main.go deploy
	# cd tests/multi_target && go run ../../main.go destroy

fmt:
	go fmt ./...

install:
	go install
