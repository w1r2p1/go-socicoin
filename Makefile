

.PHONY: gscc clean

GOBIN = $(shell pwd)/build/bin
GO ?= latest


gscc:
	build/env.sh go build -o $(GOBIN)/gscc ./cmd/gscc 
	#build/env.sh go install ./cmd/gscc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gscc\" to launch gscc."

all: gscc
	@gofmt -w .
	#build/env.sh go install
	
clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/* 
	rm -fr build/_workspace

