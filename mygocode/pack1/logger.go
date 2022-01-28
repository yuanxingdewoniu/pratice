package main

import "log"

//声明日志写入器接口
type LogWriter interface {
	Write(data interface{}) error
}


// 日志器
type Logger struct {
	writerList []LogWriter
}

// 注册一个日志写入器
func （l *Logger） RegisterWriter(writer LogWriter) {
	log.WriterList = append( l.writerList, writer)

}

// 将一个data 类型的数据写入日志
func (l *Logger) Log(data interface{}) {
	//for _, writer := range l.writerList {
	// 将日志输出到每一个写入器中
	writer.Write(data)
}

func NewLogger() *Logger {
	return &Logger{}
}

