go mod vendor -v

SET CGO_ENABLED=0
SET GOARCH=amd64

SET GOOS=darwin
go build -o build-exe\FastFileSafe-macos

SET GOOS=linux
go build -o build-exe\FastFileSafe-linux


SET GOOS=windows
go build -ldflags "-w" -o build-exe\FastFileSafe.exe

SET GOOS=linux
SET GOARCH=arm64
go build -o -o build-exe\FastFileSafe-linux-arm64



