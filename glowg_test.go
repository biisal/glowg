package glowg

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"
)

func TestLogs(t *testing.T) {
	now = func() time.Time {
		return time.Date(2026, 1, 1, 12, 0, 0, 0, time.UTC)
	}

	tests := []struct {
		name     string
		input    string
		noColor  bool
		expected string
	}{
		{
			name:  "test info with color",
			input: "hello world",
			expected: fmt.Sprintf("%s%-9s%s %s[%s]%s%s %s\n",
				ColorCyan, "INFO:", NC,
				ColorBlack, "2026-01-01 12:00:00", NC,
				"", "hello world"),
		},
		{
			name:    "test info without color",
			input:   "hello world",
			noColor: true,
			expected: fmt.Sprintf("%-9s [%s] %s\n",
				"INFO:",
				"2026-01-01 12:00:00",
				"hello world"),
		},
		{
			name:  "unicode message",
			input: "नमस्ते 🌍",
			expected: fmt.Sprintf("%s%-9s%s %s[%s]%s%s %s\n",
				ColorCyan, "INFO:", NC,
				ColorBlack, "2026-01-01 12:00:00", NC,
				"", "नमस्ते 🌍"),
		},
		{
			name:  "message with newline",
			input: "hello\nworld",
			expected: fmt.Sprintf("%s%-9s%s %s[%s]%s%s %s\n",
				ColorCyan, "INFO:", NC,
				ColorBlack, "2026-01-01 12:00:00", NC,
				"", "hello\nworld"),
		},
		{
			name:    "empty message without color",
			input:   "",
			noColor: true,
			expected: fmt.Sprintf("%-9s [%s] %s\n",
				"INFO:",
				"2026-01-01 12:00:00",
				""),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			SetOutput(&buf)
			SetNoColor(tt.noColor)
			Info(tt.input)
			if buf.String() != tt.expected {
				t.Errorf("got = %q, \nwant = %q", buf.String(), tt.expected)
			}
		})
	}
}

func TestSetOutput(t *testing.T) {
	var buf bytes.Buffer

	SetOutput(&buf)
	Info("hello")

	if buf.Len() == 0 {
		t.Fatal("expected output to be written")
	}
}

func TestSetFile(t *testing.T) {
	makeFileName := func() string {
		bytes := make([]byte, 8)
		_, _ = rand.Read(bytes)
		randomString := hex.EncodeToString(bytes)
		return randomString + ".log"
	}
	fileName := makeFileName()
	defer func() { _ = os.Remove(fileName) }()

	err := SetLogFile(fileName)
	defer CloseFile()

	if err != nil {
		t.Fatal(err)
	}

	fileName = ""
	err = SetLogFile(fileName)
	if !errors.Is(err, ErrorEmptyFilename) {
		t.Fatalf("got error %v \nexpected %v", err, ErrorEmptyFilename)
	}
}
