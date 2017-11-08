# diffgo

```
NAME:
   diffgo - a tool for comparing the two files and show the differences

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.1.0

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -o "/tmp/diff.log"	output diff data into file
   --help, -h		show help
   --version, -v	print the version
```

### 获取工具

```
go get github.com/domac/diffgo
```

### 使用例子

```
go run main.go  /tmp/t1.txt /tmp/t2.txt
```

或 

```
go run main.go -o=/tmp/demo.txt  /tmp/t1.txt /tmp/t2.txt
```