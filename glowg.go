// Package glowg implements a simple logging library.
package glowg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type LogLevel int32

const (
	LevelDebug LogLevel = iota + 1
	LevelInfo
	LevelSuccess
	LevelWarning
	LevelError

	ColorBlack   = "\x1b[90m"
	ColorRed     = "\x1b[91m"
	ColorGreen   = "\x1b[92m"
	ColorYellow  = "\x1b[93m"
	ColorMagenta = "\x1b[95m"
	ColorCyan    = "\x1b[96m"
	NC           = "\x1b[0m"
)

var (
	mu           sync.Mutex
	noColor      atomic.Bool
	currentLevel atomic.Int32
	output       io.Writer = os.Stdout
	logFile      atomic.Pointer[os.File]
	now          = time.Now

	ErrorEmptyFilename = fmt.Errorf("glowg: filename can't be empty")
)

func logWrite(levelColor, levelLabel string, withCaller bool, format string, args []any) {
	msg := fmt.Sprintf(format, args...)
	ts := now().Format("2006-01-02 15:04:05")

	caller := ""
	if withCaller {
		_, file, line, ok := runtime.Caller(2)
		if ok {
			caller = fmt.Sprintf(" [%s:%d]", filepath.Base(file), line)
		}
	}

	mu.Lock()
	defer mu.Unlock()

	plain := fmt.Sprintf("%-9s [%s]%s %s\n", levelLabel+":", ts, caller, msg)
	if noColor.Load() {
		_, _ = fmt.Fprint(output, plain)
	} else {
		colored := fmt.Sprintf("%s%-9s%s %s[%s]%s%s %s\n",
			levelColor, levelLabel+":", NC,
			ColorBlack, ts, NC,
			caller, msg)
		_, _ = fmt.Fprint(output, colored)
	}

	if f := logFile.Load(); f != nil {
		_, _ = fmt.Fprint(f, plain)
	}
}

func checkLevel(l LogLevel) bool { return l >= LogLevel(currentLevel.Load()) }

func Debug(format string, args ...any) {
	if checkLevel(LevelDebug) {
		logWrite(ColorMagenta, "DEBUG", false, format, args)
	}
}

func Info(format string, args ...any) {
	if checkLevel(LevelInfo) {
		logWrite(ColorCyan, "INFO", false, format, args)
	}
}

func Success(format string, args ...any) {
	if checkLevel(LevelSuccess) {
		logWrite(ColorGreen, "SUCCESS", false, format, args)
	}
}

func Warning(format string, args ...any) {
	if checkLevel(LevelWarning) {
		logWrite(ColorYellow, "WARNING", true, format, args)
	}
}

func Error(format string, args ...any) {
	if checkLevel(LevelError) {
		logWrite(ColorRed, "ERROR", true, format, args)
	}
}

func Errorln(args ...any) {
	if checkLevel(LevelError) {
		msg := fmt.Sprintln(args...)
		logWrite(ColorRed, "ERROR", true, "%s", []any{msg[:len(msg)-1]})
	}
}
