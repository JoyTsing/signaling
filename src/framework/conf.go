package framework

import (
	"signaling/third/goconfig"
)

type FrameworkConf struct {
	logDir      string
	logFile     string
	logLevel    string
	logToStderr bool

	httpPort         int
	httpStaticDir    string
	httpStaticPrefix string

	httpsPort int
	httpsCert string
	httpsKey  string
}

var configFile *goconfig.ConfigFile

func GetStaticDir() string {
	return gconf.httpStaticDir
}

func loadConf(confFile string) (*FrameworkConf, error) {
	var err error
	configFile, err = goconfig.LoadConfigFile(confFile)
	if err != nil {
		return nil, err
	}
	conf := &FrameworkConf{}
	// log config
	if conf.logDir, err = configFile.GetValue("log", "logDir"); err != nil {
		return nil, err
	}

	if conf.logFile, err = configFile.GetValue("log", "logFile"); err != nil {
		return nil, err
	}

	if conf.logLevel, err = configFile.GetValue("log", "logLevel"); err != nil {
		return nil, err
	}

	if conf.logToStderr, err = configFile.Bool("log", "logToStderr"); err != nil {
		return nil, err
	}

	//http config
	if conf.httpPort, err = configFile.Int("http", "port"); err != nil {
		return nil, err
	}

	if conf.httpStaticDir, err = configFile.GetValue("http", "staticDir"); err != nil {
		return nil, err
	}

	if conf.httpStaticPrefix, err = configFile.GetValue("http", "staticPrefix"); err != nil {
		return nil, err
	}

	//https config
	if conf.httpsPort, err = configFile.Int("https", "port"); err != nil {
		return nil, err
	}

	if conf.httpsCert, err = configFile.GetValue("https", "cert"); err != nil {
		return nil, err
	}

	if conf.httpsKey, err = configFile.GetValue("https", "key"); err != nil {
		return nil, err
	}

	return conf, nil
}
