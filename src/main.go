package main

import (
	"signaling/src/framework"
)

func main() {
	err := framework.Init("./conf/framework.conf")
	if err != nil {
		panic(err)
	}
	err = framework.StartHttpServer()
	if err != nil {
		panic(err)
	}
}
