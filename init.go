package log

var pref string
var id string
var pref_id string

const STREAMER string = "S"
const CONTROLLER string = "C"
const NODE string = "N"

var (
	colorOff    = []byte("\033[0m")
	colorRed    = []byte("\033[0;31m")
	colorGreen  = []byte("\033[0;32m")
	colorOrange = []byte("\033[0;33m")
	//colorBlue   = []byte("\033[0;34m")
	colorPurple = []byte("\033[0;35m")
	colorCyan   = []byte("\033[0;36m")
	colorGray   = []byte("\033[0;37m")
)

func Init(prefixString string, identityString string) {
	pref = string(prefixString) + "="
	id = identityString + " "

	switch prefixString {
	case STREAMER:
		pref_id = string(colorOrange) + pref + id + string(colorOff)

	case NODE:
		pref_id = string(colorGreen) + pref + id + string(colorOff)

	case CONTROLLER:
		pref_id = string(colorPurple) + pref + id + string(colorOff)

	default:
		pref_id = string(colorGray) + "__________" + string(colorOff)
	}

}
