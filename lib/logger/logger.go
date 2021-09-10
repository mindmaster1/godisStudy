package logger

import(
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"
)

type Settings struct{
	Path string 'yaml: "path"'
	Name string 'yaml: "name" '
	Ext  string 'yaml: "ext'
	TiemFormat string 'yaml: "time-format'
}

var(
	logFile   *os.file
	defalutPrefix  = "" 
	defaultCallerDepth = 2
	logger    *log.logger
	mu sync.Mutex
	logPrefix    = ""
	levelFlags   = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"} 
)

type logLevel int

// log levels
const (
	DEBUG logLevel = iota
	INFO
	WARNING
	ERROR
	FATAL
)

//LstFlags是标准Logger的初始值
const flags = log.LstdFlags

//New创建一个Logger。参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）
func init(){
	logger = log.New(os.StdOut,defalutPrefix,flags)
}

//logger设置初始化函数
func Setup(settings *Settings){
	var err error
	dir := settings.Path
	//Sprintf根据format参数生成格式化的字符串并返回该字符串。
	fileName := fmt.Sprintf("%s-%s.%s",
		settings.Name
		//Now获取当前时间，Format根据layout指定的格式返回t代表的时间点的格式化文本表示
		time.Now.Format(settinngs.TimeFormat),
		settings.Ext)
	//创建一个初始logger配置文件
	logFile,err = mustOpen(fileName,dir)
	if err!= nil{
		log.Fatalf("logging.Setup err:%s",err)
	}

	mw: = io.MultiWriter(os.StdOut,logFile)
	logger = log.New(mw,defaultPrefix,flags)
}

func setPrefix(level logLevel){
	//Caller报告当前go程调用栈所执行的函数的文件和行号信息
	_,file,line,ok := runtime.Caller(defaultCallerDepth)
	if ok{
		logPrefix = fmt.Sprintf("[%s][%s:%d]",levelFlags[level],filepath.Base(file),line)
	}else{
		logPrefix = fmt.Sprintf("[%s]",levelFlags[level])
	}
	logger.SetPrefix(logPrefix)
}

//Debug可以接受多个不定参数
func Debug(v ...interface()){
	//加锁
	mu.Lock()
	defer mu.Unlock()
	setPrefiX(DEBUG)
	//关于...，是go的一种语法糖。 1.主要用于函数有多个不定参数的情况 2.slice可以被打散进行处理
	//将v打散一个一个输出
	logger.Println(v....)
}

// Info prints normal log
func Info(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(INFO)
	logger.Println(v...)
}

// Warn prints warning log
func Warn(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(WARNING)
	logger.Println(v...)
}

// Error prints error log
func Error(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(ERROR)
	logger.Println(v...)
}

// Fatal prints error log then stop the program
func Fatal(v ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	setPrefix(FATAL)
	logger.Fatalln(v...)
}

















