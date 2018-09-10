# This is a very simple Makefile used to build the http-facter binary.
init:
	env GOOS=linux GOARCH=amd64 go build -o ./cmd/linux/http-facter .
	env GOOS=windows GOARCH=amd64 go build -o ./cmd/windows/http-facter .
	env GOOS=darwin GOARCH=amd64 go build -o ./cmd/mac/http-facter .

mac:
	env GOOS=darwin GOARCH=amd64 go build -o ./cmd/mac/http-facter .

linux:
	env GOOS=linux GOARCH=amd64 go build -o ./cmd/linux/http-facter .

windows:
	env GOOS=windows GOARCH=amd64 go build -o ./cmd/windows/http-facter.exe .