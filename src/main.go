package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	demo01()
}

// 你可以创建很多instance
var lgr = log.New()

func demo02() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		lgr.Out = file
	} else {
		lgr.Info("Failed to log to file, using default stderr")
	}
	lgr.WithFields(log.Fields{
		"filename": "123.txt",
	}).Info("打开文件失败")
}

// 设置log的参数
//func init() {
//	//设置输出样式，自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
//	log.SetFormatter(&log.TextFormatter{})
//	//设置output,默认为stderr,可以为任何io.Writer，比如文件*os.File
//	log.SetOutput(os.Stdout)
//	//设置最低loglevel
//	log.SetLevel(log.InfoLevel)
//
//}

// 这是一个使用了fields的例子，可以添加多对field
func demo01() {
	log.WithFields(log.Fields{
		"animal": "walrus", "a": 11,
	}).Info("A walrus appears")
}

// logrus 日志等级
func logLevel() {

	log.SetLevel(log.DebugLevel)

	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	log.Fatal("Bye.")         //log之后会调用os.Exit(1)
	log.Panic("I'm bailing.") //log之后会panic()
}
