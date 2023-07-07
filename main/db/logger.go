package db

import (
	log "github.com/sirupsen/logrus"
	l "xorm.io/xorm/log"
)

type Logger struct {
	*log.Entry
	level l.LogLevel
}

func (l *Logger) SetLevel(level l.LogLevel) {
	l.level = level
}

func (l *Logger) Level() l.LogLevel {
	return l.level
}

func (l *Logger) ShowSQL(show ...bool) {}
func (l *Logger) IsShowSQL() bool      { return false }
