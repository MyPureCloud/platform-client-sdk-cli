package logger

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var (
	traceLogger   *log.Logger
	infoLogger    *log.Logger
	warningLogger *log.Logger
	fatalLogger   *log.Logger
)

func InitLogger() {
	var fileName string
	switch runtime.GOOS {
	case "windows":
		logsDir := fmt.Sprintf("%s\\%s", os.Getenv("TEMP"), "GenesysCloud")
		if err := mkdirIfNotExist(logsDir); err != nil {
			return
		}
		fileName = fmt.Sprintf("%s\\%s", logsDir, "gc.txt")
	case "darwin":
		homeDir, _ := os.UserHomeDir()
		logsDir := fmt.Sprintf("%s/Library/Logs/GenesysCloud", homeDir)
		if err := mkdirIfNotExist(logsDir); err != nil {
			return
		}
		fileName = fmt.Sprintf("%s/%s", logsDir, "gc.log")
	default:
		logsDir := "/tmp/GenesysCloud"
		if err := mkdirIfNotExist(logsDir); err != nil {
			return
		}
		fileName = fmt.Sprintf("%s/%s", logsDir, "gc.log")
	}

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return
	}

	flags := log.Ldate | log.Ltime
	traceLogger = log.New(file, "TRACE: ", flags)
	infoLogger = log.New(file, "INFO: ", flags)
	warningLogger = log.New(file, "WARNING: ", flags)
	fatalLogger = log.New(file, "FATAL: ", flags)

	log.SetOutput(file)
}

func Trace(v ...interface{}) {
	fmt.Fprint(os.Stderr, v...)
	if traceLogger != nil {
		traceLogger.Println(v...)
	}
}

func Tracef(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	if traceLogger != nil {
		traceLogger.Printf(format, v...)
	}
}

func Info(v ...interface{}) {
	if infoLogger != nil {
		infoLogger.Println(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if infoLogger != nil {
		infoLogger.Printf(format, v...)
	}
}

func Warn(v ...interface{}) {
	if warningLogger != nil {
		warningLogger.Println(v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if warningLogger != nil {
		warningLogger.Printf(format, v...)
	}
}

func Fatal(v ...interface{}) {
	fmt.Fprint(os.Stderr, v...)
	if fatalLogger != nil {
		fatalLogger.Fatal(v...)
	} else {
		os.Exit(1)
	}
}

func Fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	if fatalLogger != nil {
		fatalLogger.Fatalf(format, v...)
	} else {
		os.Exit(1)
	}
}

func mkdirIfNotExist(directory string) error {
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		err = os.Mkdir(directory, 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
