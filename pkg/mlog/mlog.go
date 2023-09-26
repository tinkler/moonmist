package mlog

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

var mlogSourceDir, mstSourceDir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	mlogSourceDir = regexp.MustCompile(`mlog\.go`).ReplaceAllString(file, "")
	mstSourceDir = filepath.ToSlash(filepath.Clean(mlogSourceDir + "/../mst"))
}

func writeRuntimeMessage() string {
	for i := 2; i < 15; i++ {
		_, file, line, ok := runtime.Caller(i)
		if ok && (!strings.HasPrefix(file, mlogSourceDir) && !strings.HasPrefix(file, mstSourceDir)) || strings.HasSuffix(file, "_test.go") {
			return fmt.Sprintf("\033[%d;1m%s:%d\033[0m", CT_BLACK, file, line)
		}
	}
	return ""
}

func wrapMessage(message string, level messageLevel) string {
	ts := time.Now().Format("2006-01-02 15:04:05")
	switch level {
	case ML_TITLE:
		message = fmt.Sprintf("%s \033[%d;1m[标题]\033[0m %s", ts, CT_CYAN, message)
	case ML_DEBUG:
		message = fmt.Sprintf("%s \033[%d;1m[调试]\033[0m %s", ts, CT_GREEN, message)
	case ML_INFO:
		message = fmt.Sprintf("%s \033[%d;1m[日志]\033[0m %s", ts, CT_BLUE, message)
	case ML_WARN:
		message = fmt.Sprintf("%s \033[%d;1m[警告]\033[0m %s", ts, CT_YELLOW, message)
	case ML_ERR:
		message = fmt.Sprintf("%s \033[%d;1m[错误]\033[0m %s", ts, CT_RED, message)
	}
	return message
}

func Format(f interface{}, vs ...any) string {
	var message string
	switch fv := f.(type) {
	case string:
		message = fv
		if len(vs) == 0 {
			return message
		}
		if !strings.Contains(message, "%") {
			message += strings.Repeat(" %v", len(vs))
		}
	default:
		message = fmt.Sprint(f)
		if len(vs) == 0 {
			return message
		}
		message += strings.Repeat(" %v", len(vs))
	}
	if len(vs) > 0 {
		return fmt.Sprintf(message, vs...)
	}
	return fmt.Sprint(message)
}

func Error(f interface{}, vs ...any) {
	b := bytes.NewBufferString(writeRuntimeMessage())
	b.WriteByte('\n')
	b.WriteString(wrapMessage(Format(f, vs...), ML_ERR))
	b.WriteString(ConsoleMessageSeparate)
	os.Stdout.Write(b.Bytes())
	writeToFile(b.Bytes(), ML_ERR)
}

func Warn(f interface{}, vs ...any) {
	b := bytes.NewBufferString(writeRuntimeMessage())
	b.WriteByte('\n')
	b.WriteString(wrapMessage(Format(f, vs...), ML_WARN))
	b.WriteString(ConsoleMessageSeparate)
	if ConsoleLevel > L_SILENT {
		os.Stdout.Write(b.Bytes())
	}
	if LoggerFileLevel > L_SILENT {
		writeToFile(b.Bytes(), ML_WARN)
	}
}

func Info(f interface{}, vs ...any) {
	b := bytes.NewBufferString(writeRuntimeMessage())
	b.WriteByte('\n')
	b.WriteString(wrapMessage(Format(f, vs...), ML_INFO))
	b.WriteString(ConsoleMessageSeparate)
	if ConsoleLevel > L_NORMAL {
		os.Stdout.Write(b.Bytes())
	}
	if LoggerFileLevel > L_NORMAL {
		writeToFile(b.Bytes(), ML_INFO)
	}
}

func Debug(f interface{}, vs ...any) {
	b := bytes.NewBufferString(writeRuntimeMessage())
	b.WriteByte('\n')
	b.WriteString(wrapMessage(Format(f, vs...), ML_DEBUG))
	b.WriteString(ConsoleMessageSeparate)
	if ConsoleLevel > L_LOG {
		os.Stdout.Write(b.Bytes())
	}
	if LoggerFileLevel > L_LOG {
		writeToFile(b.Bytes(), ML_DEBUG)
	}
}

func Title(f interface{}, vs ...any) {
	b := bytes.NewBufferString(wrapMessage(Format(f, vs...), ML_TITLE))
	b.WriteByte('\n')
	os.Stdout.Write(b.Bytes())
	writeToFile(b.Bytes(), ML_TITLE)
}
