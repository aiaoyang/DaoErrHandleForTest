#!/bin/bash
export GOBIN="$(pwd)/bin"
go install ./...
bin/query