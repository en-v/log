package log

import (
	"fmt"
	"strings"
)

var TYPE_STRINGS = [3]string{
	"DBG ",
	"INF ",
	string(red) + "ERR " + string(nocolor)}

const EVENT_TYPE_DEBUG, EVENT_TYPE_INFO, EVENT_TYPE_ERROR = 0, 1, 2

func keyval(T int, msg string, keysvals ...interface{}) {
	if T == EVENT_TYPE_DEBUG && !SHOW_DEBUG_ITEMS {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(id)
	sb.WriteString(TYPE_STRINGS[T])
	head(&sb)
	sb.WriteString(msg)
	if len(keysvals) > 0 {
		sb.WriteString(", ")
	}
	args(&sb, keysvals, true)
	tofile(T, &sb)
}

func trc(T int, values ...interface{}) {
	if T == EVENT_TYPE_DEBUG && !SHOW_DEBUG_ITEMS {
		return
	}

	sb := strings.Builder{}
	sb.WriteString(id)
	sb.WriteString(TYPE_STRINGS[T])

	head(&sb)

	sb.WriteString(string(red))
	args(&sb, values, false)
	sb.WriteString(string(nocolor))

	tofile(T, &sb)
}

func asarr(T int, values ...interface{}) {
	if T == EVENT_TYPE_DEBUG && !SHOW_DEBUG_ITEMS {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(id)
	sb.WriteString(TYPE_STRINGS[T])

	head(&sb)

	args(&sb, values, false)
	tofile(T, &sb)
}

func asfmt(T int, msg string, vals ...interface{}) {
	if T == EVENT_TYPE_DEBUG && !SHOW_DEBUG_ITEMS {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(id)
	sb.WriteString(TYPE_STRINGS[T])

	head(&sb)

	sb.WriteString(fmt.Sprintf(msg, vals...))
	sb.WriteString(";")

	args(&sb, vals, false)
	tofile(T, &sb)
}
