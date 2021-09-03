package log

import (
	"fmt"
	"strings"
)

var TYPES = [3]string{
	"DBG ",
	"INF ",
	string(colorRed) + "ERR " + string(colorOff)}

const T_DBG, T_NFO, T_ERR = 0, 1, 2

func jsn(T int, msg string, keysAndValues ...interface{}) {
	if T == T_DBG && !SHOW_DBG {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(pref_id)
	sb.WriteString(TYPES[T])
	head(&sb)
	sb.WriteString(msg)
	if len(keysAndValues) > 0 {
		sb.WriteString(", ")
	}
	buildArrayString(&sb, keysAndValues, true)
	printToDestination(T, &sb)
}

func trc(T int, values ...interface{}) {
	if T == T_DBG && !SHOW_DBG {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(pref_id)
	sb.WriteString(TYPES[T])
	head(&sb)
	sb.WriteString(string(colorRed))
	buildArrayString(&sb, values, false)
	printToDestination(T, &sb)
	sb.WriteString(string(colorOff))
}

func arr(T int, values ...interface{}) {
	if T == T_DBG && !SHOW_DBG {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(pref_id)
	sb.WriteString(TYPES[T])
	head(&sb)

	buildArrayString(&sb, values, false)
	printToDestination(T, &sb)
}

func fm(T int, msg string, keysAndValues ...interface{}) {
	if T == T_DBG && !SHOW_DBG {
		return
	}
	sb := strings.Builder{}
	sb.WriteString(pref_id)
	sb.WriteString(TYPES[T])
	head(&sb)

	sb.WriteString(fmt.Sprintf(msg, keysAndValues...))
	sb.WriteString(";")
	buildArrayString(&sb, keysAndValues, false)
	printToDestination(T, &sb)
}

func Debugw(msg string, keysAndValues ...interface{}) { jsn(T_DBG, msg, keysAndValues...) }
func Infow(msg string, keysAndValues ...interface{})  { jsn(T_NFO, msg, keysAndValues...) }
func Errorw(msg string, keysAndValues ...interface{}) { jsn(T_ERR, msg, keysAndValues...) }

func Debug(args ...interface{})       { arr(T_DBG, args...) }
func Info(args ...interface{})        { arr(T_NFO, args...) }
func Error(args ...interface{})       { arr(T_ERR, args...) }
func ErrorWrap(err error, msg string) { arr(T_ERR, msg+" > "+err.Error()) }

func Debugf(template string, args ...interface{}) { fm(T_DBG, template, args...) }
func Infof(template string, args ...interface{})  { fm(T_NFO, template, args...) }
func Errorf(template string, args ...interface{}) { fm(T_ERR, template, args...) }

func Trace(args ...interface{}) { trc(T_DBG, args...) }
