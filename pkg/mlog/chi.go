package mlog

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

type logEntry struct {
	*LogFormatter
	request *http.Request
	buf     *bytes.Buffer
}

func (l *logEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	ts := time.Now().Format("2006-01-02 15:04:05")
	switch {
	case status < 200:
		l.buf.WriteString(fmt.Sprintf("%s \033[%d;1m[%s]\033[0m", ts, CT_BLUE, strconv.Itoa(status)))
	case status < 300:
		l.buf.WriteString(fmt.Sprintf("%s \033[%d;1m[%s]\033[0m", ts, CT_GREEN, strconv.Itoa(status)))
	case status < 400:
		l.buf.WriteString(fmt.Sprintf("%s \033[%d;1m[%s]\033[0m", ts, CT_CYAN, strconv.Itoa(status)))
	case status < 500:
		l.buf.WriteString(fmt.Sprintf("%s \033[%d;1m[%s]\033[0m", ts, CT_YELLOW, strconv.Itoa(status)))
	default:
		l.buf.WriteString(fmt.Sprintf("%s \033[%d;1m[%s]\033[0m", ts, CT_RED, strconv.Itoa(status)))
	}

	l.buf.WriteString(fmt.Sprintf("\033[%d;1m %dB \033[0m", CT_BLUE, bytes))

	l.buf.WriteString(" in ")
	if elapsed < 500*time.Millisecond {
		l.buf.WriteString(fmt.Sprintf("\033[%d;1m %s \033[0m", CT_GREEN, elapsed.String()))
	} else if elapsed < 5*time.Second {
		l.buf.WriteString(fmt.Sprintf("\033[%d;1m %s \033[0m", CT_YELLOW, elapsed.String()))
	} else {
		l.buf.WriteString(fmt.Sprintf("\033[%d;1m %s \033[0m", CT_RED, elapsed.String()))
	}
	l.buf.WriteRune('\n')

	if ConsoleLevel > L_LOG {
		os.Stdout.Write(l.buf.Bytes())
	}
}

type LogFormatter struct {
	routePath map[string]string
}

func NewLogFormatter(noColor bool) *LogFormatter {
	return &LogFormatter{routePath: make(map[string]string)}
}

func (l *LogFormatter) AddRouteInfo(routePath map[string]string) {
	if ConsoleLevel == L_LOG {
		for k, v := range routePath {
			l.routePath[k] = v
		}
	}
}

func (l *LogFormatter) NewLogEntry(r *http.Request) middleware.LogEntry {
	entry := &logEntry{
		LogFormatter: l,
		request:      r,
		buf:          &bytes.Buffer{},
	}

	if fileLine, ok := l.routePath[r.RequestURI]; ok {
		entry.buf.WriteString(
			fmt.Sprintf("\033[%d;1m %s \033[0m\n", CT_BLACK, fileLine))
	}

	reqID := middleware.GetReqID(r.Context())
	if reqID != "" {
		entry.buf.WriteString(fmt.Sprintf("\033[%d;1m["+reqID+"]\033[0m", CT_YELLOW))
	}
	entry.buf.WriteString(fmt.Sprintf("\033[%d;1m\"\033[0m", CT_CYAN))
	entry.buf.WriteString(fmt.Sprintf("\033[%d;1m %s \033[0m", CT_PURPLE, r.Method))

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	entry.buf.WriteString(fmt.Sprintf("\033[%d;1m%s://%s%s %s\" \033[0m", CT_CYAN, scheme, r.Host, r.RequestURI, r.Proto))

	entry.buf.WriteString("from ")
	entry.buf.WriteString(r.RemoteAddr)
	entry.buf.WriteString(" - ")

	return entry
}

func (l *logEntry) Panic(v interface{}, stack []byte) {
	middleware.PrintPrettyStack(v)
}

// NewLogEntry creates a new LogEntry for the request.
// reference from go-chi/chi/v5/middleware/logger.go
func RequestLogger(f middleware.LogFormatter) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			entry := f.NewLogEntry(r)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				entry.Write(ww.Status(), ww.BytesWritten(), ww.Header(), time.Since(t1), nil)
			}()

			next.ServeHTTP(ww, middleware.WithLogEntry(r, entry))
		}
		return http.HandlerFunc(fn)
	}
}

func ChiLogger(f func(formatter *LogFormatter)) func(next http.Handler) http.Handler {
	color := true
	if runtime.GOOS == "windows" {
		color = false
	}
	formatter := NewLogFormatter(!color)
	if f != nil {
		f(formatter)
	}
	return RequestLogger(formatter)
}
