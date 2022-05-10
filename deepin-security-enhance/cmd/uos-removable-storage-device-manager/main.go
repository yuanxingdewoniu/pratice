package main

import (
	"deepin-security-enhance/pkg/serve"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"deepin-security-enhance/pkg/netlink"

	"github.com/kr/pretty"
)

// 主函数
func main() {
	srv := serve.GetService()
	err := srv.Init(serve.RemovableStorageDeviceDaemon)
	if err != nil {
		panic(err)
	}

	uevent, err := getOptionnalMatcher()
	if err != nil {
		log.Fatalln(err)
	}

	monitor(uevent)

	srv.Loop()
}

var filePath *string

// getOptionnalMatcher Parse and load config file which contains rules for matching
func getOptionnalMatcher() (matcher netlink.Matcher, err error) {

	if filePath == nil || *filePath == "" {
		return nil, nil
	}

	stream, err := ioutil.ReadFile(*filePath)
	if err != nil {
		return nil, err
	}

	if stream == nil {
		return nil, fmt.Errorf("Empty, no rules provided in \"%s\", err: %w", *filePath, err)
	}

	var rules netlink.RuleDefinitions
	if err := json.Unmarshal(stream, &rules); err != nil {
		return nil, fmt.Errorf("Wrong rule syntax, err: %w", err)
	}

	return &rules, nil
}

// monitor run monitor mode
func monitor(matcher netlink.Matcher) {
	log.Println("Monitoring UEvent kernel message to user-space...")

	conn := new(netlink.UEventConn)
	if err := conn.Connect(netlink.UdevEvent); err != nil {
		log.Fatalln("Unable to connect to Netlink Kobject UEvent socket")
	}
	defer conn.Close()

	queue := make(chan netlink.UEvent)
	errors := make(chan error)

	log.Println("start Handlen  message \n")
	// Handling message from queue
	for {
		select {
		case uevent := <-queue:
			//解析uevent 事件， 同时发送状态信息
			log.Println("Handlen  new ", pretty.Sprint(uevent))
		case err := <-errors:
			log.Println("ERROR:", err)
		}
	}

}
