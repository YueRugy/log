package main

import (
	"fmt"
	"github.com/yueRugy/log/logAgent/kfaka"
	"github.com/yueRugy/log/logAgent/tail"
	"gopkg.in/ini.v1"
	"time"
)

var conf = new(Conf)

type Conf struct {
	Kafak *Kafak `ini:"kafak"`
	Tail  *Tail  `ini:"tail"`
}

type Kafak struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}
type Tail struct {
	Path string `ini:"path"`
}

func run() {
	//tail读
	for {
		select {
		case line := <-tail.Read():
			kfaka.SendToKfaka(conf.Kafak.Topic, line.Text)
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func main() {
	ctx, err4 := ini.Load("/home/yue/code/github.com/yueRugy/log/logAgent/conf/config.ini")

	if err4 != nil {
		fmt.Printf("load config failed %v", err4)
		return
	}
	err5 := ctx.MapTo(conf)
	if err5 != nil {
		fmt.Printf("map failed %v", err5)
	}
	err := kfaka.Init([]string{conf.Kafak.Address})
	if err != nil {
		fmt.Println("init kfaka failed ", err)
		return
	}

	err1 := tail.Init(conf.Tail.Path)
	if err1 != nil {
		fmt.Println("init tail failed", err1)
		return
	}

	run()
}
