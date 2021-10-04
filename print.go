package log

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func head(sb *strings.Builder) {

	if SHOW_EVENT_TIME {
		sb.WriteString(time.Now().Format(TIME_FORMAT))
	}

	if SHOW_CODE_STACK {
		_, file, line, ok := runtime.Caller(3)

		if ok {

			sb.Write(cyan)
			fnparts := strings.Split(file, "/")
			pcount := len(fnparts)

			if pcount >= 2 {
				fnparts[pcount-1] = strings.Replace(fnparts[pcount-1], ".go", ":", 1)

				sb.WriteString(fnparts[pcount-2])
				sb.WriteString("/")
				sb.WriteString(fnparts[pcount-1])

			} else {
				sb.WriteString(file)
			}

			sb.WriteString(strconv.Itoa(line))
			sb.Write(nocolor)

		} else {
			sb.WriteString("LOG-ERROR: call stack info!")
		}
	}
	sb.WriteString(HDR_DELIMITER)
}

func args(sb *strings.Builder, args []interface{}, pairs bool) {

	end := len(args) - 1

	for p := range args {

		switch args[p].(type) {
		case string:
			sb.WriteString(args[p].(string))

		case int:
			sb.WriteString(strconv.Itoa(args[p].(int)))

		case []byte:
			barr := args[p].([]byte)
			if len(barr) > 512 {
				sb.WriteString(string(barr[:127]) + "... (+" + strconv.Itoa(len(barr)-128) + " bytes )")
			} else {
				sb.WriteString(string(barr))
			}

		default:
			sb.WriteString(fmt.Sprint(args[p]))
		}

		if pairs {
			if p%2 == 0 {
				sb.WriteString(" ")
			} else {
				if p != end {
					sb.WriteString(", ")
				}
			}

		} else {
			if p != end {
				sb.WriteString(", ")
			}
		}
	}
}

func tofile(mt int, sb *strings.Builder) {

	sb.WriteString(EOL)
	str := sb.String()

	var err error

	switch mt {
	case EVENT_TYPE_DEBUG:
		print(str)
		_, err = fdbg.WriteString(str)

	case EVENT_TYPE_INFO:
		print(str)
		_, err = fnfo.WriteString(str)

	case EVENT_TYPE_ERROR:
		print(str)
		_, err = ferr.WriteString(str)

	default:
		err = errors.New("Unknown log file type " + strconv.Itoa(mt))
	}

	if err != nil {
		panic(err)
	}
}
