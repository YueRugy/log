package tail

import "github.com/hpcloud/tail"

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
	t, err = tail.TailFile(fileName, config)

	return err
}

func Read() chan *tail.Line {
	return t.Lines
}
