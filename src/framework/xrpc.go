package framework

import "fmt"

func Call(serviceName string, request interface{}, response interface{}, logId uint32) error {
	fmt.Println("Call:[", serviceName, request, response, logId, "]")
	return nil
}
