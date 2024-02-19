package main

import (
	"signaling/src/framework"
)

func startHttp() {
	if err := framework.StartHttpServer(); err != nil {
		panic(err)
	}
}

func startHttps() {
	if err := framework.StartHttpsServer(); err != nil {
		panic(err)
	}
}

func main() {
	if err := framework.Init("./conf/framework.conf"); err != nil {
		panic(err)
	}
	// 静态资源/static
	framework.RegisterStaticUrl()

	go startHttp()
	startHttps()
}
