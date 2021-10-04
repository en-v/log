package log

import (
	"io"
	"os"
	"time"
)

var (
	fnfo *os.File
	fdbg *os.File
	ferr *os.File

	err error
)

func init() {

	path := "./logs"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	fdbg = fopen(path + "/debug.log")
	fnfo = fopen(path + "/info.log")
	ferr = fopen(path + "/error.log")
}

func fopen(name string) *os.File {

	fl, err := os.OpenFile(name, os.O_WRONLY|os.O_RDONLY|os.O_CREATE, DATAFILE_PERMISSION)
	if err != nil {
		panic(err)
	}

	_, err = fl.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}

	_, err = fl.WriteString("\n\r##### APP LAUNCH ##### " + time.Now().Format(TIME_FORMAT) + " ##### APP LAUNCH #####\n\r")
	if err != nil {
		panic(err)
	}

	return fl
}
