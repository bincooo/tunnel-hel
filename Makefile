checklint:
ifeq (, $(shell which golangci-lint))
	@echo 'error: golangci-lint is not installed, please exec `brew install golangci-lint`'
	@exit 1
endif

patch:
ifeq (, $(shell which gopatch))
	@echo 'error: gopatch is not installed, please exec `go install github.com/uber-go/gopatch@latest`'
	@exit 1
endif

lint: checklint
	golangci-lint run --skip-dirs-use-default

fix: patch fix-quic

fix-quic: patch
	gopatch -p ./fixs/tongsq_gost+quic.go.patch $(GOPATH)/pkg/mod/github.com/tongsq/gost@v1.3.0/quic.go