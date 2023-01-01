#!/bin/bash

go fmt ./...
go vet ./...

go test -v
go test -bench=.
godoc -http=:8080
