package log

const DATAFILE_PERMISSION = 0664
const EOL = "\r\n"

const SHOW_CODE_STACK = true
const SHOW_DEBUG_ITEMS = true
const SHOW_EVENT_TIME = false

const TIME_FORMAT = "06.01.02 15:04:05.000"
const HDR_DELIMITER = " "

type ColorName string

const ORANGE ColorName = "orange"
const PURPLE ColorName = "purple"
const GREEN ColorName = "green"
