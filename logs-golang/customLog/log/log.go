package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	RESET_COLOR = "\033[0m"
	BLUE        = "\033[34m"
	YELLOW      = "\033[33m"
	RED         = "\033[31m"
	CYAN        = "\u001b[36m"
)

var (
	logger *lumberjack.Logger
	projectPath, _ = os.Getwd()
	projectRoot    = projectPath + string(filepath.Separator)
)

func init() {

	logFilePath := "app.log"

	logger = &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    1,
		MaxBackups: 2,
		MaxAge:     30,
		Compress:   false,
	}

	consoleWriter := os.Stdout
	multiWriter := io.MultiWriter(consoleWriter, logger)

	log.SetOutput(multiWriter)
	log.SetFlags(log.LstdFlags)
}

func Info(args ...any) {
	msg := createLogMessage(args)
	logging("INFO", msg, BLUE)
}

func Warn(args ...any) {
	msg := createLogMessage(args)
	logging("WARN", msg, YELLOW)
}

func Error(args ...any) {
	msg := createLogMessage(args)
	logging("ERROR", msg, RED)
}

func getStacktrace(msg string) string {
	err := errors.WithStack(fmt.Errorf("%s", msg))
	stackTrace := fmt.Sprintf("%+v\n", err)

	stackTrace = sanitizeStacktrace(stackTrace)

	return stackTrace
}

func createLogMessage(args ...any) string {
	var msgParts []string

	if len(args) == 1 {
		if slice, ok := args[0].([]interface{}); ok {
			args = slice
		}
	}

	for _, arg := range args {
		msgParts = append(msgParts, fmt.Sprintf("%v", arg))
	}

	msg := strings.Join(msgParts, " ")
	return msg
}

func logging(level, str, color string) {
	if len(str) > 0 {
		pc, file, line, ok := runtime.Caller(2)
		if ok {
			level := fmt.Sprintf("%s[%s]%s", color, level, RESET_COLOR)
			funcName := getFunctionName(runtime.FuncForPC(pc).Name())
			relativeFilePath := sanitizeFilePath(file)
			file := fmt.Sprintf("%s[%s.%s:%d]%s", CYAN, relativeFilePath, funcName, line, RESET_COLOR)
			log.Printf("%s %s: %v", level, file, str)
		} else {
			log.Printf("%v", str)
		}

		if level == "ERROR" {
			stackTrace := getStacktrace(str)
			header := fmt.Sprintf("%sStack Trace%s:", color, RESET_COLOR)
			log.Printf("%s %s", header, stackTrace)
		}
	}
}

func getFunctionName(fullName string) string {
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}

func sanitizeFilePath(filePath string) string {
	if strings.HasPrefix(filePath, projectRoot) {
		return strings.Replace(filePath, projectRoot, "", 1)
	}
	return filePath
}

func sanitizeStacktrace(stackTrace string) string {
	if strings.Contains(stackTrace, projectRoot) {
		return strings.Replace(stackTrace, projectRoot, "", -1)
	}
	return stackTrace
}
