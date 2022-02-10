package main

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writeList []LogWriter
}

func (l *Logger) RegisterWriter(writer LogWriter) {
	l.writeList = append(l.writeList, writer)
}

func (l *Logger) Log(data interface{}) {

	//遍历所有注册的写入器
	for _, writer := range l.writeList {

		//将日志输出到每一个写入器中
		writer.Write(data)
	}

}
N