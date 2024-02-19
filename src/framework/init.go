package framework

import (
	"flag"

	"signaling/third/glog"
)

var gconf *FrameworkConf

func Init(confFile string) error {
	var err error

	if gconf, err = loadConf(confFile); err != nil {
		return err
	}

	flag.Parse() //glog使用前需要
	//ini解析
	//fmt.Printf("gconf: %+v\n", gconf)

	glog.SetLogDir(gconf.logDir)
	glog.SetLogFileName(gconf.logFile)
	glog.SetLogLevel(gconf.logLevel)
	glog.SetLogToStderr(gconf.logToStderr)
	return nil
}
