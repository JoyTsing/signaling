package framework

import (
	"flag"

	"signaling/src/glog"
)

var gconf *FrameworkConf

func Init(confFile string) error {
	var err error

	gconf, err = loadConf(confFile)
	if err != nil {
		return err
	}

	//fmt.Printf("%+v\n", gconf)

	flag.Parse() //glog需要

	glog.SetLogDir(gconf.logDir)
	glog.SetLogFileName(gconf.logFile)
	glog.SetLogLevel(gconf.logLevel)
	glog.SetLogToStderr(gconf.logToStderr)
	return nil
}
