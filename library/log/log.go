package log

import (
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/tokopedia/logistic/svc/mp-logistic/infra/config"
	c "github.com/tokopedia/logistic/svc/mp-logistic/infra/constant"
	dd "github.com/tokopedia/logistic/svc/mp-logistic/infra/datadog"
	"github.com/tokopedia/logistic/svc/mp-logistic/infra/errors"
)

var (
	stdOut = flag.String("mp_log_i", "", "log file for stdout")
	stdErr = flag.String("mp_log_e", "", "log file for stderr")

	logOut *logrus.Logger
	logErr *logrus.Logger
	once   sync.Once
)

const rootPath = "github.com/tokopedia/logistic/svc/mp-logistic/"

func InitLogger(cfg config.LoggerCfg) {
	once.Do(func() {
		out := reopen(1, *stdOut)
		if out == nil {
			out = os.Stdout
		}
		logOut = createLogger(cfg, out)

		err := reopen(2, *stdErr)
		if err == nil {
			err = os.Stderr
		}
		logErr = createLogger(cfg, err)
	})
}

func InitMockLogger() {
	cfg := config.LoggerCfg{
		JsonFormat: false,
		Level:      5,
	}
	logOut = createLogger(cfg, os.Stdout)
	logErr = createLogger(cfg, os.Stderr)
}

func createLogger(cfg config.LoggerCfg, io io.Writer) *logrus.Logger {
	log := logrus.New()
	if cfg.JsonFormat {
		log.Formatter = &logrus.JSONFormatter{
			TimestampFormat: c.TimeFormatStd,
		}
	} else {
		log.Formatter = &logrus.TextFormatter{
			ForceColors:     true,
			TimestampFormat: c.TimeFormatStd,
			FullTimestamp:   true,
		}
	}

	log.Out = io
	log.Level = logrus.Level(cfg.Level)
	return log
}

func Info(string string, options ...Option) {
	fields := logrus.Fields{}
	for _, opt := range options {
		k, v := opt()
		fields[k] = v
	}
	logOut.WithFields(fields).Info(string)
}

func Infof(string string, args ...interface{}) {
	logOut.Infof(string, args...)
}

func Debug(string string, options ...Option) {
	fields := logrus.Fields{}
	for _, opt := range options {
		k, v := opt()
		fields[k] = v
	}
	logOut.WithFields(fields).Debug(string)
}

func Debugf(string string, args ...interface{}) {
	logOut.Debugf(string, args...)
}

func Errors(string string, options ...Option) {
	fields := logrus.Fields{}
	for _, opt := range options {
		k, v := opt()
		fields[k] = v
	}

	fun, file, line := getLineInfo()
	s := fmt.Sprintf("[%s]%s:%d", fun, file, line)
	fields["line"] = trimRootPath(s)

	logErr.WithFields(fields).Error(string)
	logDD(fields)
}

func Errorw(error error, options ...Option) {
	if error == nil {
		return
	}
	errWrapper, ok := error.(errors.ErrorWrapper)
	if !ok || errWrapper == nil {
		logErr.Error(error.Error())
		return
	}

	fields := logrus.Fields{}
	for _, opt := range options {
		k, v := opt()
		fields[k] = v
	}

	for k, v := range errWrapper.GetFields() {
		fields[k] = v
	}

	fun, file, line := getFuncFileLine(errWrapper)
	s := fmt.Sprintf("[%s]%s:%d", fun, file, line)
	fields["line"] = trimRootPath(s)

	logErr.WithFields(fields).
		Error(fmt.Sprintf("%s: %s", error.Error(), errWrapper.GetMessage()))
	logDD(fields)
}

func logDD(f logrus.Fields) {
	if ddTags, ok := f[errors.DDTagsErr].(string); ok {
		tags := strings.Split(ddTags, "|")
		dd.Count(fmt.Sprintf("%s.%s.%s", c.AppName, "errors", "hit_count"), 1, tags, 1)
	}
}

type Option func() (string, interface{})

func WithErrDDTags(d errors.DDTags) Option {
	return func() (string, interface{}) {
		return errors.DDTagsErr, d.GetString()
	}
}

func WithOrderID(orderID int64) Option {
	return func() (string, interface{}) {
		return "order_id", orderID
	}
}

func WithDeliveryID(dID int64) Option {
	return func() (string, interface{}) {
		return "delivery_id", dID
	}
}

func WithHistoryID(hID int64) Option {
	return func() (string, interface{}) {
		return "history_id", hID
	}
}

func WithMessage(s string) Option {
	return func() (string, interface{}) {
		return "message", s
	}
}

func WithMessagef(format string, args ...interface{}) Option {
	return func() (string, interface{}) {
		return "message", fmt.Sprintf(format, args...)
	}
}

func WithField(key string, value interface{}) Option {
	return func() (string, interface{}) {
		return key, value
	}
}

func WithFieldf(key string, format string, args ...interface{}) Option {
	return func() (string, interface{}) {
		return key, fmt.Sprintf(format, args...)
	}
}

func getFuncFileLine(wrapper errors.ErrorWrapper) (string, string, int) {
	st := wrapper.StackTrace()
	f := st[0]
	pc := uintptr(f) - 1
	fn := runtime.FuncForPC(pc)
	file, line := fn.FileLine(pc)
	return funcname(fn.Name()), file, line
}

func funcname(name string) string {
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}

func getLineInfo() (string, string, int) {
	pc := make([]uintptr, 100)
	c := runtime.Callers(3, pc)
	if c == 0 {
		return "???", "???", 1
	}
	frames := runtime.CallersFrames(pc)
	f, _ := frames.Next()
	return funcname(f.Function), f.File, f.Line
}

//reopen opens the existing files as a log destination
//arg filename:	file to be reopened
//arg fd: new file descriptor which is used by dup2() to create a copy.
//returns file:	result file after opened, if success, otherwise nil
func reopen(fd int, filename string) *os.File {
	if filename == "" {
		return nil
	}

	logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		logrus.Error("fail opening file log")
		return nil
	}

	if err = syscall.Dup2(int(logFile.Fd()), fd); err != nil {
		logrus.Error("fail to dup")
	}

	return logFile
}

func trimRootPath(file string) string {
	lastBin := strings.LastIndex(file, rootPath)
	if (lastBin+len(rootPath)) > len(file) || lastBin == -1 {
		return file
	}
	return file[lastBin+len(rootPath):]
}
