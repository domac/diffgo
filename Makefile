.PHONY: all clean

all: format test build

test:
	go test -v . 

format:
	gofmt -w ./app  ./main.go

build:
	# 设置交叉编译参数:
	# GOOS为目标编译系统, mac os则为 "darwin", window系列则为 "windows"
	# 生成二进制执行文件 diffgo , 如在windows下则为 diffgo.exe
	GOOS="linux" GOARCH="amd64" go build -v -o diffgo ./main.go #&& cp -rf config_sample.json builds/config.json

clean:
	go clean -i