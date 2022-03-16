go mod vendor -v

SET CGO_ENABLED=0
SET GOARCH=amd64

SET GOOS=darwin
go build -o FastFileSafe-macos

SET GOOS=linux
go build -o FastFileSafe-linux


SET GOOS=windows
go build -ldflags "-H windows -w"

SET GOOS=linux
SET GOARCH=arm64
go build -o FastFileSafe-linux-arm64



