package tail

import (
	"fmt"
	"github.com/hpcloud/tail"
	"path/filepath"
)

var t *tail.Tail

func Init(fileName string) (err error) {
	config := tail.Config{
		Location: &tail.SeekInfo{
			Offset: 0,
			Whence: 2,
		},
		ReOpen:    true,
		MustExist: false,
		Poll:      true,
		Follow:    true,
	}
	fmt.Println(filepath.Abs(fileName))
	t, err = tail.TailFile(fileName, config)

	return err
}

func Read() chan *tail.Line {
	return t.Lines
}
