package main

import (
	"fmt"
	"github.com/yueRugy/log/logAgent/kfaka"
	"github.com/yueRugy/log/logAgent/tail"
	"time"
)

func run() {
	//tail读
	for {
		select {
		case line := <-tail.Read():
			fmt.Println(line.Text)
			kfaka.SendToKfaka("web_log", line.Text)
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	err := kfaka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Println("init kfaka failed ", err)
		return
	}

	err1 := tail.Init("/home/yue/code/github.com/yueRugy/log/logAgent/myLog")
	if err1 != nil {
		fmt.Println("init tail failed", err1)
		return
	}

	run()
}
