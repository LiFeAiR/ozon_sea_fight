LOCAL_BIN := $(CURDIR)/bin

.deps:
	$(info Installing dependencies...)
	go mod download

deps: .deps

.install-gomock:
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.43.2

GENMOCKS   = $(shell find . -name 'genmock.go'| sort -u)
GOMOCKERY  = $(LOCAL_BIN)/mockery
$(LOCAL_BIN)/mockery: REPOSITORY=github.com/vektra/mockery/v2@v2.43.2

.PHONY: generate-mocks
generate-mocks: .install-gomock deps
	$(GOMOCKERY) --version; \
    for f in $(GENMOCKS); do \
        echo $$f;\
        go generate $$f;\
    done;
