package main

import (
	"fmt"
	"net"
)

func  main()  {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err :=net.Listen("tcp", "localhost:50000")
	if err !=nil {
		fmt.Println("Error listening", err.Error())
		return // 终止程序
	}
// 监听并接受客户端的链接
	for {
		conn, err :=  listener.Accept()
		if err != nil {
			fmt.Println("Error accecpting ", err.Error())
		}
		go doServerStuff(conn)
	}
}

func  doServerStuff(conn net.Conn)  {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)

		if err != nil {
			fmt.Println("Error reading ", err.Error())
			return //终止程序
		}
		fmt.Println("Received data : %v", string(buf[:len]))
	}

}


