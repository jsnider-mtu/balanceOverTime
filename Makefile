build: linux windows

linux:
	GOOS=linux GOARCH=amd64 go build -o bot_linux_amd64 .

windows:
	GOOS=windows GOARCH=amd64 go build -o bot_windows_amd64.exe .
