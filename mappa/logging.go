package mappa

import (
	"io"
	"log"
	"log/syslog"
)

type InternalLogger struct {
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	debugLogger   *log.Logger
}

type SysLogger struct {
	InternalLogger
	sysLogTag string
	syslogger *syslog.Writer
}

func CreateSysLogger(sysLogTag string) *SysLogger {
	syslogger, err := syslog.New(syslog.LOG_INFO, sysLogTag)
	if err != nil {
		log.Fatalf("Failed to create syslogger '%s' %v", sysLogTag, err)
	}
	logger := SysLogger{sysLogTag: sysLogTag, syslogger: syslogger}
	logger.createLoggers(syslogger)
	return &logger
}

func (logger *InternalLogger) LogInfof(format string, a ...interface{}) {
	logger.infoLogger.Printf(format, a...)
}

func (logger *InternalLogger) LogWarningf(format string, a ...interface{}) {
	logger.warningLogger.Printf(format, a...)
}

func (logger *InternalLogger) LogDebugf(format string, a ...interface{}) {
	logger.debugLogger.Printf(format, a...)
}

func (logger *InternalLogger) LogErrorf(format string, a ...interface{}) {
	logger.errorLogger.Printf(format, a...)
}

func (internalLogger *InternalLogger) createLoggers(writer io.Writer) {
	internalLogger.infoLogger = log.New(writer, "INFO :", log.Ldate|log.Ltime|log.Lshortfile)
	internalLogger.warningLogger = log.New(writer, "WARN :", log.Ldate|log.Ltime|log.Lshortfile)
	internalLogger.errorLogger = log.New(writer, "ERROR:", log.Ldate|log.Ltime|log.Lshortfile)
	internalLogger.debugLogger = log.New(writer, "DEBUG:", log.Ldate|log.Ltime|log.Lshortfile)
}
