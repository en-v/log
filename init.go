package log

var id string

var (
	nocolor = []byte("\033[0m")
	red     = []byte("\033[0;31m")
	green   = []byte("\033[0;32m")
	orange  = []byte("\033[0;33m")
	//colorBlue   = []byte("\033[0;34m")
	purple = []byte("\033[0;35m")
	cyan   = []byte("\033[0;36m")
	gray   = []byte("\033[0;37m")
)

func Init(color ColorName, idstr string) {

	switch color {
	case ORANGE:
		id = string(orange) + idstr + string(nocolor)

	case GREEN:
		id = string(green) + idstr + string(nocolor)

	case PURPLE:
		id = string(purple) + idstr + string(nocolor)

	default:
		id = string(gray) + " NO ID " + string(nocolor)
	}

}
