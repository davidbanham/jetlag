#!/bin/bash
GOOS=darwin GOARCH=amd64 go build -o jetlag-osx-arm64
GOOS=linux GOARCH=arm64 go build -o jetlag-linux-arm64
GOOS=linux GOARCH=amd64 go build -o jetlag-linux-amd64
GOOS=linux GOARCH=386 go build -o jetlag-linux-386
GOOS=windows GOARCH=386 go build -o jetlag-windows-386.exe
GOOS=windows GOARCH=amd64 go build -o jetlag-windows-amd64.exe
