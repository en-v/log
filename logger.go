package log

import (
	"errors"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const CALLER_STACK = true
const SHOW_DBG = true
const SHOW_TIME = false
const TS_FORMAT = "06.01.02 15:04:05.000"
const HDR_DLMTR = " "

func head(sb *strings.Builder) {

	if SHOW_TIME {
		sb.WriteString(time.Now().Format(TS_FORMAT))
	}

	if CALLER_STACK {
		_, file, line, ok := runtime.Caller(3)
		//sb.WriteString(" ")

		if ok {

			sb.Write(colorCyan)
			fileparts := strings.Split(file, "/")
			fpl := len(fileparts)

			if fpl >= 2 {
				fileparts[fpl-1] = strings.Replace(fileparts[fpl-1], ".go", ":", 1)

				sb.WriteString(fileparts[fpl-2])
				sb.WriteString("/")
				sb.WriteString(fileparts[fpl-1])

			} else {
				sb.WriteString(file)
			}

			sb.WriteString(strconv.Itoa(line))
			sb.Write(colorOff)

		} else {
			sb.WriteString("LOGGER ERROR: caller info is not OK!!! ")
		}
	}
	sb.WriteString(HDR_DLMTR)
}

func buildArrayString(sb *strings.Builder, args []interface{}, pairs bool) {

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

func printToDestination(mt int, sb *strings.Builder) {

	sb.WriteString(EOL)
	str := sb.String()

	var err error
	switch mt {
	case T_DBG:
		print(str)
		_, err = dbgfile.WriteString(str)

	case T_NFO:
		print(str)
		_, err = nfofile.WriteString(str)

	case T_ERR:
		print(str)
		_, err = errfile.WriteString(str)

	default:
		err = errors.New("Unknown log file type " + strconv.Itoa(mt))
	}

	if err != nil {
		panic(err)
	}
}
