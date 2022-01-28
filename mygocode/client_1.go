package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main()  {
	conn, err := net.Dial("tcp","localhost:500000")

	if err != nil {
		// 由于目标计算机积极拒绝而无法创建连接
		fmt.Println("Error dialing ",err.Error())
		return //终止程序
	}
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("First, what is your name")

	clientName, _ := inputReader.ReadString('\n')

	trimmedClient := string.Trim(clientName, '\r\n')

	for {
		fmt.Println("What to send to the server ? Type Q to quit.")
		input, _
	}

}