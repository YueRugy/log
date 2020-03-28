package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main() {
	fileNmae := "./myLog"
	config := tail.Config{
		Location:    &tail.SeekInfo{Offset: 0, Whence: 2},
		ReOpen:      true,
		MustExist:   false,
		Poll:        true,
		Pipe:        false,
		RateLimiter: nil,
		Follow:      true,
		MaxLineSize: 0,
		Logger:      nil,
	}

	tails, err := tail.TailFile(fileNmae, config)
	if err != nil {
		fmt.Println("tail file failed ", err)
	}

	var (
		line *tail.Line
		ok   bool
	)
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen ,fileNmae =%s", fileNmae)
		}
		fmt.Println(line.Text)
	}
}
