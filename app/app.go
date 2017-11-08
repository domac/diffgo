package app

import (
	"bufio"
	"errors"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var ErrDiff = errors.New("diff files fail")

func flagsInit() {
	AddFlagString(cli.StringFlag{
		Name:  "o",
		Value: "/tmp/diff.log",
		Usage: "output diff data into file",
	})
}

//应用启动Action
func appAction(c *cli.Context) (err error) {

	if len(c.Args()) < 2 {
		panic("please input source file and target file")
	}
	//以命令行的方式启动
	sourceFilePath := c.Args()[0]
	targetFilePath := c.Args()[1]
	output := c.String("o")
	return filesdiff(sourceFilePath, targetFilePath, output)
}

//文件内容比对
func filesdiff(sourceFilePath, targetFilePath, output string) error {

	//源文件内容
	sources, err := readFile(sourceFilePath)
	if err != nil {
		return err
	}

	//目标文件内容
	targets, err := readFile(targetFilePath)
	if err != nil {
		return err
	}

	diffs := DiffOnly(sources, targets)

	if len(diffs) == 0 {
		log.Println("there is no difference !")
		return nil
	}

	if output == "" {
		for _, di := range diffs {
			println(di.String())
		}
	} else {
		if !IsExist(output) {
			CreateFile(output)
		}
		//预分配空间
		result := make([]string, len(diffs))
		for i, di := range diffs {
			text := di.String()
			result[i] = text
		}
		//覆盖文件方式写入数据
		WriteIntoFile(output, result, WRITE_OVER)
	}
	return nil
}

//读取文件内容
func readFile(filepath string) ([]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	out := []string{}
	s := bufio.NewScanner(f)
	for s.Scan() {
		if t := s.Text(); t != "" {
			out = append(out, t)
		}
	}
	log.Printf("read  %s completed !\n", filepath)
	return out, s.Err()
}

//程序主入口
func Startup() {
	flagsInit()
	app := cli.NewApp()
	app.Name = "diffgo"
	app.Usage = "a tool for comparing the two files and show the differences"
	app.Version = "0.1.0"
	app.Flags = GetAppFlags()
	app.Action = ActionWrapper(appAction)
	app.Run(os.Args)
}
