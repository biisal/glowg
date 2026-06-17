# glowg

A minimal, colorful logging library for Go. Zero dependencies.

<img width="1427" height="346" alt="glowg" src="https://github.com/user-attachments/assets/5dc742b4-3ad4-4bb7-8642-7016d8b62fd8" />


## Install
```sh
go get github.com/biisal/glowg
```

## Usage

```go
package main

import "github.com/biisal/glowg"

func main() {
    glowg.Debug("starting application")
    glowg.Info("Hello, %s!", "world")
    glowg.Success("connected to database")
    glowg.Warning("memory usage is high")
    glowg.Error("failed to connect: %v", err)
    glowg.Errorln("something went wrong", "details here")
}
```

## Log Levels

| Level | Color | Caller Info |
|---------|---------|-------------|
| `LevelDebug` | Magenta | No |
| `LevelInfo` | Cyan | No |
| `LevelSuccess` | Green | No |
| `LevelWarning` | Yellow | Yes |
| `LevelError` | Red | Yes |

Set the minimum level to filter output:

```go
glowg.SetLogLevel(glowg.LevelWarning) // only Warning and Error
```

## Options

```go
// Write logs to a file (in addition to stdout)
glowg.SetLogFile("logs/app.log")
defer glowg.CloseFile()

// Redirect console output
glowg.SetOutput(os.Stderr)

// Disable ANSI colors
glowg.SetNoColor(true)
```

## License

[MIT](LICENSE) 
