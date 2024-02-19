package framework

import (
	"fmt"
	"math/rand"
	"time"

	"signaling/src/glog"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GetLogId32() uint32 {
	return rand.Uint32()
}

func NewSource(i int64) {
	panic("unimplemented")
}

type logItem struct {
	filed string
	value string
}

type timeItem struct {
	filed     string
	beginTime int64
	endTime   int64
}

type ComLog struct {
	mainLog []logItem
	timeLog []timeItem
}

func (l *ComLog) AddLogItem(filed, value string) {
	l.mainLog = append(l.mainLog, logItem{filed, value})
}

func (l *ComLog) TimeBegin(filed string) {
	l.timeLog = append(l.timeLog, timeItem{filed, time.Now().UnixNano() / 1000, 0})
}

func (l *ComLog) TimeEnd(filed string) {
	for i, item := range l.timeLog {
		if item.filed == filed {
			l.timeLog[i].endTime = time.Now().UnixNano() / 1000
			return
		}
	}
}

func (l *ComLog) getPrefixLog() string {
	var prefix string = ""
	for _, item := range l.mainLog {
		prefix += fmt.Sprintf("	%s [%s]\n", item.filed, item.value)
	}

	for _, item := range l.timeLog {
		prefix += fmt.Sprintf("	%s costs [%.3fms]\n", item.filed, float64(item.endTime-item.beginTime)/1000.0)
	}
	return prefix
}

func (l *ComLog) Debugf(format string, args ...interface{}) {
	totalLog := l.getPrefixLog() + format
	glog.Debugf(totalLog, args...)
}

func (l *ComLog) Infof(format string, args ...interface{}) {
	totalLog := l.getPrefixLog() + format
	glog.Infof(totalLog, args...)
}

func (l *ComLog) Warningf(format string, args ...interface{}) {
	totalLog := l.getPrefixLog() + format
	glog.Warningf(totalLog, args...)
}

func (l *ComLog) Fatalf(format string, args ...interface{}) {
	totalLog := l.getPrefixLog() + format
	glog.Fatalf(totalLog, args...)
}

func (l *ComLog) Errorf(format string, args ...interface{}) {
	totalLog := l.getPrefixLog() + format
	glog.Errorf(totalLog, args...)
}
