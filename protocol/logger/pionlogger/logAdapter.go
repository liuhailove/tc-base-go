package pionlogger

import (
	"fmt"
	"strings"

	"github.com/liuhailove/tc-base-go/protocol/logger"
)

// implements webrtc.LeveledLogger
type logAdapter struct {
	logger          logger.Logger
	ignoredPrefixes []string
}

func (l *logAdapter) Trace(msg string) {
	// ignore trace
}

func (l *logAdapter) Tracef(format string, args ...interface{}) {
	// ignore trace
}

func (l *logAdapter) Debug(msg string) {
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Debugw(msg)
}

func (l *logAdapter) Debugf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Debugw(msg)
}

func (l *logAdapter) Info(msg string) {
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Infow(msg)
}

func (l *logAdapter) Infof(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Infow(msg)
}

func (l *logAdapter) Warn(msg string) {
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Warnw(msg, nil)
}

func (l *logAdapter) Warnf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Warnw(msg, nil)
}

func (l *logAdapter) Error(msg string) {
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Errorw(msg, nil)
}

func (l *logAdapter) Errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	if l.shouldIgnore(msg) {
		return
	}
	l.logger.Errorw(msg, nil)
}

func (l *logAdapter) shouldIgnore(msg string) bool {
	for _, prefix := range l.ignoredPrefixes {
		if strings.HasPrefix(msg, prefix) {
			return true
		}
	}
	return false
}
