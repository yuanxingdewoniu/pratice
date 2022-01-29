package main

import "io"

// 声明一个设备结构体
type device struct {

}

// 实现io.Writer的Write()方法
func (d *device) Write(p []byte) (n int, err error){
	return 0, nil
}

// io.Closer的Close()方法
func (d *device) close() error  {
 return nil
}

func main()  {
	// 声明写入关闭器，并赋于device实例
	var wc io.WriteCloser = new(device)

    // 写入数据
    wc.Write(nil)

	// 关闭设备
	wc.Close()

	// 声明写入器， 并赋于device 实例
	var writeOnly io.ByteWriter = new(device)

	// 写入数据
	writeOnly.Write(nil)
}

