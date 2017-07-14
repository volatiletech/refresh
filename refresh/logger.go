package refresh

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/fatih/color"
	homedir "github.com/mitchellh/go-homedir"
)

const lformat = "=== %s ==="

type Logger struct {
	log *log.Logger
}

func NewLogger(c *Configuration) *Logger {
	color.NoColor = !c.EnableColors
	if runtime.GOOS == "windows" {
		color.NoColor = true
	}
	lname := String(c.LogName, "refresh")
	return &Logger{
		log: log.New(os.Stdout, fmt.Sprintf("%s: ", lname), log.LstdFlags),
	}
}

func String(s1, s2 string) string {
	if s1 == "" {
		return s2
	}
	return s1
}

func (l *Logger) Success(msg interface{}, args ...interface{}) {
	l.log.Print(color.GreenString(fmt.Sprintf(lformat, msg), args...))
}

func (l *Logger) Error(msg interface{}, args ...interface{}) {
	l.log.Print(color.RedString(fmt.Sprintf(lformat, msg), args...))
}

func (l *Logger) Print(msg interface{}, args ...interface{}) {
	l.log.Printf(fmt.Sprintf(lformat, msg), args...)
}

var LogLocation = func() string {
	dir, _ := homedir.Dir()
	dir, _ = homedir.Expand(dir)
	dir = path.Join(dir, ".refresh")
	os.MkdirAll(dir, 0755)
	return dir
}

var ErrorLogPath = func() string {
	return path.Join(LogLocation(), ID()+".err")
}
