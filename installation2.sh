#!/bin/bash

#!/bin/bash
export CGO_ENABLED=0   # remove once headers are fixed, if you want cgo back for dlv

go install golang.org/x/tools/gopls@latest
go install github.com/go-delve/delve/cmd/dlv@latest
go install github.com/air-verse/air@latest
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
