

.PHONY: gscc clean

GOBIN = build/bin
GO ?= latest



gscc:
	build/env.sh go build  ./cmd/gscc
	build/env.sh go install ./cmd/gscc
	@echo "Done building."
	@echo "Run \"$(GOBIN)/gscc\" to launch gscc."

all: 
	build/env.sh go install
	
clean:
	rm -fr build/_workspace/pkg/ $(GOBIN)/*
	rm -fr build/_workspace