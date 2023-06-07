# 编译 main.go 到目录 dist 文件名为 aichat_linux 可执行文件

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64
go build -ldflags "-s -w" -o dist/aichat_linux main.go

#export CGO_ENABLED=0
#export GOOS=windows
#export GOARCH=amd64
#go build -ldflags "-s -w" -o dist/aichat_window.exe main.go

#
#export CGO_ENABLED=0
#export GOOS=darwin
#export GOARCH=amd64
#go build -ldflags "-s -w" -o dist/aichat_darwin main.go
