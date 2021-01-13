package log

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	//红色
	errLog = log.New(os.Stdout, "\033[31m[error]\033[0m", log.LstdFlags|log.Lshortfile)
	//蓝色
	infoLog = log.New(os.Stdout, "\033[34m[info]\033[0m", log.LstdFlags|log.Lshortfile)
	loggers = []*log.Logger{errLog, infoLog}
	mu      sync.Mutex

	Error  = errLog.Println
	ErrorF = errLog.Printf
	Info   = infoLog.Println
	Infof  = infoLog.Printf
)

//log level
const (
	InfoLevel = iota
	ErrorLevel
	Disabled
)

// set log level
func SetLevel(level int) {
	mu.Lock()
	defer mu.Unlock()

	for _, logger := range loggers {
		logger.SetOutput(os.Stdout)
	}
	//如果设置为ErrorLevel、InfoLevel的输出会被定向到ioutil.Discard，即不打印该日志
	if ErrorLevel < level {
		errLog.SetOutput(ioutil.Discard)
	}
	if InfoLevel < level {
		infoLog.SetOutput(ioutil.Discard)
	}
}
