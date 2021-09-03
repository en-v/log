package log

import (
	"io"
	"os"
	"time"
)

var (
	nfofile *os.File
	dbgfile *os.File
	errfile *os.File

	err error
)

func init() {

	path := "./logs"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	dbgfile = openFile(path + "/debug.log")
	nfofile = openFile(path + "/info.log")
	errfile = openFile(path + "/error.log")
}

func openFile(name string) *os.File {

	fl, err := os.OpenFile(name, os.O_WRONLY|os.O_RDONLY|os.O_CREATE, DATAFILE_PERMISSION)
	if err != nil {
		panic(err)
	}

	_, err = fl.Seek(0, io.SeekEnd)
	if err != nil {
		panic(err)
	}

	_, err = fl.WriteString("\n\r##### APP LAUNCH ##### " + time.Now().Format(TS_FORMAT) + " ##### APP LAUNCH #####\n\r")
	if err != nil {
		panic(err)
	}

	return fl

}
