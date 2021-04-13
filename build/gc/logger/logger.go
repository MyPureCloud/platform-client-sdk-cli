package logger

import (
	"fmt"
	"github.com/mypurecloud/platform-client-sdk-cli/build/gc/config"
	"log"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	traceLogger    *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	fatalLogger    *log.Logger
	loggingEnabled bool
)

func InitLogger(cmd *cobra.Command) {
	profileName, _ := cmd.Flags().GetString("profile")
	c, err := config.GetConfig(profileName)
	if err != nil {
		return
	}
	loggingEnabled = c.LoggingEnabled()

	fileName := c.LogFilePath()
	if fileName == "" {
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
	}

	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}

	flags := log.Ldate | log.Ltime
	traceLogger = log.New(file, "TRACE: ", flags)
	infoLogger = log.New(file, "INFO: ", flags)
	warningLogger = log.New(file, "WARNING: ", flags)
	fatalLogger = log.New(file, "FATAL: ", flags)

	traceLogger.SetOutput(file)
	infoLogger.SetOutput(file)
	warningLogger.SetOutput(file)
	traceLogger.SetOutput(file)

	log.SetOutput(file)
}

func Trace(v ...interface{}) {
	fmt.Fprint(os.Stderr, v...)
	if traceLogger != nil && loggingEnabled {
		traceLogger.Println(v...)
	}
}

func Tracef(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	if traceLogger != nil && loggingEnabled {
		traceLogger.Printf(format, v...)
	}
}

func Info(v ...interface{}) {
	if infoLogger != nil && loggingEnabled {
		infoLogger.Println(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if infoLogger != nil && loggingEnabled {
		infoLogger.Printf(format, v...)
	}
}

func Warn(v ...interface{}) {
	if warningLogger != nil && loggingEnabled {
		warningLogger.Println(v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if warningLogger != nil && loggingEnabled {
		warningLogger.Printf(format, v...)
	}
}

func Fatal(v ...interface{}) {
	fmt.Fprint(os.Stderr, v...)
	if fatalLogger != nil && loggingEnabled {
		fatalLogger.Fatal(v...)
	} else {
		os.Exit(1)
	}
}

func Fatalf(format string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, format, v...)
	if fatalLogger != nil && loggingEnabled {
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
