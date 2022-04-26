package main

import "deepin-security-enhance/pkg/serve"

// 主函数
func main() {
	srv := serve.GetService()
	err := srv.Init()
	if err != nil {
		panic(err)
	}
	srv.Loop()
}
