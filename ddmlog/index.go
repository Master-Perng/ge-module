package ddmlog

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

// InitLog 初始化日志
func InitLog(logFileName string) (*os.File, error) {
	writer, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	return writer, err
}

type Logger struct {
	mu        sync.Mutex
	prefix    string
	flag      int
	out       io.Writer
	buf       []byte
	isDiscard int32
}

func New(out io.Writer, prefix string, flag int) *Logger {
	l := &Logger{out: out, prefix: prefix, flag: flag}
	if out == io.Discard {
		l.isDiscard = 1
	}
	return l
}

func itoa(buf *[]byte, i int, wid int) {
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	*buf = append(*buf, b[bp:]...)
}

func (l *Logger) formatHeader(buf *[]byte, t time.Time, file string, line int) {
	if l.flag&log.Lmsgprefix == 0 {
		*buf = append(*buf, l.prefix...)
	}
	if l.flag&(log.Ldate|log.Ltime|log.Lmicroseconds) != 0 {
		if l.flag&log.LUTC != 0 {
			t = t.UTC()
		}
		if l.flag&log.Ldate != 0 {
			year, month, day := t.Date()
			*buf = append(*buf, '[')
			itoa(buf, year, 4)
			*buf = append(*buf, '/')
			itoa(buf, int(month), 2)
			*buf = append(*buf, '/')
			itoa(buf, day, 2)
			*buf = append(*buf, ']')
			*buf = append(*buf, ' ')
		}
		if l.flag&(log.Ltime|log.Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			*buf = append(*buf, '[')
			itoa(buf, hour, 2)
			*buf = append(*buf, ':')
			itoa(buf, min, 2)
			*buf = append(*buf, ':')
			itoa(buf, sec, 2)
			if l.flag&log.Lmicroseconds != 0 {
				*buf = append(*buf, '.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			*buf = append(*buf, ']')
			*buf = append(*buf, ' ')

		}
	}
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		if l.flag&log.Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		*buf = append(*buf, "["...)
		*buf = append(*buf, file...)

		*buf = append(*buf, ':')
		itoa(buf, line, -1)
		*buf = append(*buf, "]"...)
		*buf = append(*buf, ": "...)

	}
	if l.flag&log.Lmsgprefix != 0 {
		*buf = append(*buf, l.prefix...)
	}
}

func (l *Logger) Output(calldepth int, s string) error {
	now := time.Now() // get this early.
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		// Release lock while getting caller info - it's expensive.
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, file, line)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	_, err := l.out.Write(l.buf)
	return err
}

func (l *Logger) OutConsole(calldepth int, s string, mode string) error {
	now := time.Now() // get this early.
	var file string
	var line int
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.flag&(log.Lshortfile|log.Llongfile) != 0 {
		l.mu.Unlock()
		var ok bool
		_, file, line, ok = runtime.Caller(calldepth)
		if !ok {
			file = "???"
			line = 0
		}
		l.mu.Lock()
	}
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, now, file, line)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	//_, err := l.out.Write(l.buf)
	if mode == "Error" {
		fmt.Println(fmt.Sprintf("\033[1;31m[E] %s\033[0m\n", string(l.buf)))
	} else if mode == "Warn" {
		fmt.Println(fmt.Sprintf("\033[1;33m[W] %s\033[0m\n", string(l.buf)))
	} else if mode == "Info" {
		fmt.Println(fmt.Sprintf("\033[1;34m[I] %s\033[0m\n", string(l.buf)))
	} else if mode == "Debug" {
		fmt.Println(fmt.Sprintf("\033[1;33m[D] %s\033[0m\n", string(l.buf)))
	} else if mode == "Trace" {
		fmt.Println(fmt.Sprintf("\033[1;33m[T] %s\033[0m\n", string(l.buf)))
	} else {
		fmt.Println(fmt.Sprintf("\033[1;34m[I] %s\033[0m\n", string(l.buf)))
	}
	return nil
}

func Write(writer *os.File, format ...any) {
	logger := New(writer, fmt.Sprintf("[%s] ", format[0]), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	logger.Output(2, fmt.Sprint(format[1:]))

}

func Info(format ...any) {
	logger := New(os.Stdout, fmt.Sprintf("[%s] ", format[0]), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	logger.OutConsole(2, fmt.Sprint(format[1:]), "Info")
}

func Error(format ...any) {
	logger := New(os.Stdout, fmt.Sprintf("[%s] ", format[0]), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	logger.OutConsole(2, fmt.Sprint(format[1:]), "Error")
}

func Debug(format ...any) {
	logger := New(os.Stdout, fmt.Sprintf("[%s] ", format[0]), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	logger.OutConsole(2, fmt.Sprint(format[1:]), "Debug")
}

func Trace(format ...any) {
	logger := New(os.Stdout, fmt.Sprintf("[%s] ", format[0]), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	logger.OutConsole(2, fmt.Sprint(format[1:]), "Trace")
}

func Warn(format ...any) {
	logger := New(os.Stdout, fmt.Sprintf("[%s] ", format[0]), log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile|log.LstdFlags)
	logger.OutConsole(2, fmt.Sprint(format[1:]), "Warn")
}
