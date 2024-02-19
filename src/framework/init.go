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
	//fmt.Printf("gconf: %+v\n", gconf)

	//log
	glogSet()

	//rpc communication
	if err = loadXrpcClient(); err != nil {
		return err
	}

	return nil
}

func glogSet() {
	flag.Parse() //glog使用前需要
	glog.SetLogDir(gconf.logDir)
	glog.SetLogFileName(gconf.logFile)
	glog.SetLogLevel(gconf.logLevel)
	glog.SetLogToStderr(gconf.logToStderr)
}
