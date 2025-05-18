up:
	cd src && reflex -s -r '\.go$$' go run cmd/app/main.go

format:
	cd src && go fmt ./...

setup:
	go install github.com/cespare/reflex@latest

.PHONY: up setup format
