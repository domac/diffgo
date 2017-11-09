# diffgo

```
NAME:
   diffgo - a tool for comparing the two files and show the differences

USAGE:
   diffgo [global options] command [command options] [arguments...]
   
VERSION:
   0.1.1
   
COMMANDS:
   help, h      Shows a list of commands or help for one command
   
GLOBAL OPTIONS:
   -o "/tmp/diff.log"   output diff data into file
   --mode               diff mode: add or delete
   --help, -h           show help
   --version, -v        print the version
```

### 获取工具

```
go get github.com/domac/diffgo
```

### 使用例子

默认执行

```
go run main.go  /tmp/t1.txt /tmp/t2.txt
```

带输出路径

```
go run main.go -o=/tmp/demo.txt  /tmp/t1.txt /tmp/t2.txt
```

带比对模式

```
.go run main.go -mode=add -o=/tmp/add.log  /tmp/t1.txt /tmp/t2.txt
```