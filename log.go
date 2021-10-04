package log

// key-values view

func Debugw(msg string, keysAndValues ...interface{}) {
	keyval(EVENT_TYPE_DEBUG, msg, keysAndValues...)
}
func Infow(msg string, keysAndValues ...interface{}) { keyval(EVENT_TYPE_INFO, msg, keysAndValues...) }
func Errorw(msg string, keysAndValues ...interface{}) {
	keyval(EVENT_TYPE_ERROR, msg, keysAndValues...)
}

// all-in-row view

func Debug(args ...interface{})       { asarr(EVENT_TYPE_DEBUG, args...) }
func Info(args ...interface{})        { asarr(EVENT_TYPE_INFO, args...) }
func Error(args ...interface{})       { asarr(EVENT_TYPE_ERROR, args...) }
func ErrorWrap(err error, msg string) { asarr(EVENT_TYPE_ERROR, msg+" > "+err.Error()) }

// formatted view

func Debugf(template string, args ...interface{}) { asfmt(EVENT_TYPE_DEBUG, template, args...) }
func Infof(template string, args ...interface{})  { asfmt(EVENT_TYPE_INFO, template, args...) }
func Errorf(template string, args ...interface{}) { asfmt(EVENT_TYPE_ERROR, template, args...) }

// all-in-red-row view 

func Trace(args ...interface{}) { trc(EVENT_TYPE_DEBUG, args...) }
