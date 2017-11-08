package app

import (
	"bufio"
	"errors"
	"github.com/codegangsta/cli"
	"log"
	"os"
)

var ErrDiff = errors.New("diff files fail")

//应用启动Action
func appAction(c *cli.Context) (err error) {

	if len(os.Args) < 3 {
		panic("please input source file and target file")
	}

	//以命令行的方式启动
	sourceFilePath := os.Args[1]
	targetFilePath := os.Args[2]
	return filesdiff(sourceFilePath, targetFilePath)
}

//文件内容比对
func filesdiff(sourceFilePath, targetFilePath string) error {
	sources, err := readFile(sourceFilePath)
	if err != nil {
		return err
	}

	targets, err := readFile(targetFilePath)
	if err != nil {
		return err
	}

	diffs := DiffOnly(sources, targets)

	if len(diffs) == 0 {
		log.Println("no difference !")
		return nil
	}

	for _, di := range diffs {
		println(di.String())
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
	app := cli.NewApp()
	app.Name = "diffgo"
	app.Usage = "a tool for comparing the two files and show the differences"
	app.Version = "0.1.0"
	app.Flags = GetAppFlags()
	app.Action = ActionWrapper(appAction)
	app.Run(os.Args)
}
