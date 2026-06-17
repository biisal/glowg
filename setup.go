package glowg

import (
	"io"
	"os"
	"path/filepath"
)

func SetLogFile(path string) error {
	if path == "" {
		return ErrorEmptyFilename
	}
	if dir := filepath.Dir(path); dir != "." {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}

	if old := logFile.Swap(f); old != nil {
		_ = old.Close()
	}
	return nil
}

func CloseFile() {
	if f := logFile.Swap(nil); f != nil {
		_ = f.Close()
	}
}
func SetOutput(w io.Writer) { mu.Lock(); output = w; mu.Unlock() }

func SetLogLevel(l LogLevel) { currentLevel.Store(int32(l)) }

func SetNoColor(v bool) { noColor.Store(v) }
