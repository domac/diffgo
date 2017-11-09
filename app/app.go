package app

import (
	"bufio"
	"errors"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"sync"
)

var ErrDiff = errors.New("diff files fail")

func flagsInit() {
	AddFlagString(cli.StringFlag{
		Name:  "o",
		Value: "/tmp/diff.log",
		Usage: "output diff data into file",
	})

	AddFlagString(cli.StringFlag{
		Name:  "mode",
		Value: "",
		Usage: "diff mode: add or delete",
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
	mode := c.String("mode")
	return filesdiff(sourceFilePath, targetFilePath, output, mode)
}

type MyDiff struct {
	sourcesMap map[string]string
	targetsMap map[string]string
}

func NewMyDiff() *MyDiff {
	return &MyDiff{
		sourcesMap: make(map[string]string),
		targetsMap: make(map[string]string),
	}
}

//文件内容比对
func filesdiff(sourceFilePath, targetFilePath, output, mode string) error {

	wg := sync.WaitGroup{}
	wg.Add(1)
	wg.Add(1)
	//源文件内容

	sources := []string{}
	targets := []string{}
	mydiff := NewMyDiff()

	if mode == "delete" || mode == "add" {
		go mydiff.readFileToMap(sourceFilePath, &wg, 0)
		go mydiff.readFileToMap(targetFilePath, &wg, 1)
	} else {
		log.Println("进入高速模式")
		go readFile(sourceFilePath, &sources, &wg)
		go readFile(targetFilePath, &targets, &wg)
	}

	wg.Wait()

	diffs := []DiffRecord{}

	if mode == "delete" {
		diffs = DiffSimpleDictDelete(mydiff.sourcesMap, mydiff.targetsMap)
	} else if mode == "add" {
		diffs = DiffSimpleDictInsert(mydiff.sourcesMap, mydiff.targetsMap)
	} else {
		diffs = DiffOnly(sources, targets)
	}

	mydiff.sourcesMap = nil
	mydiff.targetsMap = nil

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
func readFile(filepath string, out *[]string, wg *sync.WaitGroup) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		wg.Done()
	}()

	s := bufio.NewScanner(f)
	for s.Scan() {
		if t := s.Text(); t != "" {
			*out = append(*out, t)
		}
	}
	log.Printf("read  %s completed !\n", filepath)
	return s.Err()
}

//读取文件内容
func (self *MyDiff) readFileToMap(filepath string, wg *sync.WaitGroup, rtype int) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer func() {
		f.Close()
		wg.Done()
	}()

	s := bufio.NewScanner(f)

	if rtype == 0 {
		for s.Scan() {
			if t := s.Text(); t != "" {
				self.sourcesMap[t] = t
			}
		}
	} else if rtype == 1 {
		for s.Scan() {
			if t := s.Text(); t != "" {
				self.targetsMap[t] = t
			}
		}
	}

	log.Printf("read  %s completed !\n", filepath)
	return s.Err()
}

//程序主入口
func Startup() {
	flagsInit()
	app := cli.NewApp()
	app.Name = "diffgo"
	app.Usage = "a tool for comparing the two files and show the differences"
	app.Version = "0.1.1"
	app.Flags = GetAppFlags()
	app.Action = ActionWrapper(appAction)
	app.Run(os.Args)
}
