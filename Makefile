all:
	make build-decrypt-linux
	make build-encrypt-linux
	mv decrypt-linux-amd64 encrypt-linux-amd64 executed-files/linux
	make build-decrypt-mac-m1
	make build-encrypt-mac-m1
	mv decrypt-darwin-arm64 encrypt-darwin-arm64 executed-files/mac
	make build-decrypt-windows
	make build-encrypt-windows
	mv decrypt-windows-amd64 encrypt-windows-amd64 executed-files/windows

build-decrypt-linux:
	GOOS=linux GOARCH=amd64 go build -o decrypt-linux-amd64 decrypt.go
	chmod +x decrypt-linux-amd64

build-decrypt-mac-m1:
	GOOS=darwin GOARCH=arm64 go build -o decrypt-darwin-arm64 decrypt.go
	chmod +x decrypt-darwin-arm64

build-decrypt-windows:
	GOOS=windows GOARCH=amd64 go build -o decrypt-windows-amd64 decrypt.go
	chmod +x decrypt-windows-amd64

build-encrypt-linux:
	GOOS=linux GOARCH=amd64 go build -o encrypt-linux-amd64 encrypt.go
	chmod +x encrypt-linux-amd64

build-encrypt-mac-m1:
	GOOS=darwin GOARCH=arm64 go build -o encrypt-darwin-arm64 encrypt.go
	chmod +x encrypt-darwin-arm64

build-encrypt-windows:
	GOOS=windows GOARCH=amd64 go build -o encrypt-windows-amd64 encrypt.go
	chmod +x encrypt-windows-amd64