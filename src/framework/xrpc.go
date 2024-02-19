package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"signaling/src/framework/xrpc"
)

var xrpcClients = make(map[string]*xrpc.Client)

func loadXrpcClient() error {
	sections := configFile.GetSectionList()
	for _, v := range sections {
		if !strings.HasPrefix(v, "xrpc.") {
			continue
		}
		mSection, err := configFile.GetSection(v)
		if err != nil {
			return err
		}
		//server
		values, ok := mSection["serverAddr"]
		if !ok {
			return errors.New("in config file, serverAddr not found")
		}
		//format
		arrServer := strings.Split(values, ",")
		for i, server := range arrServer {
			arrServer[i] = strings.TrimSpace(server)
		}
		//client
		client := xrpc.NewClient(arrServer)
		// timeout
		setTimeConf(mSection, client)

		xrpcClients[v] = client
	}
	return nil
}

func setTimeConf(mSection map[string]string, client *xrpc.Client) {
	if values, ok := mSection["connectTimeout"]; ok {
		if connectTimeout, err := strconv.Atoi(values); err == nil {
			client.ConnectTimeout = time.Duration(connectTimeout) * time.Millisecond
		}
	}

	if values, ok := mSection["readTimeout"]; ok {
		if readTimeout, err := strconv.Atoi(values); err == nil {
			client.ReadTimeout = time.Duration(readTimeout) * time.Millisecond
		}
	}

	if values, ok := mSection["writeTimeout"]; ok {
		if writeTimeout, err := strconv.Atoi(values); err == nil {
			client.WriteTimeout = time.Duration(writeTimeout) * time.Millisecond
		}
	}
}

func Call(serviceName string, request interface{}, response interface{}, logId uint32) error {
	//fmt.Println("Call:[", serviceName, "{", request, "}", response, logId, "]")
	client, ok := xrpcClients["xrpc."+serviceName]
	if !ok {
		return fmt.Errorf("xrpc client not found for %s", serviceName)
	}
	content, err := json.Marshal(request)
	if err != nil {
		return err
	}
	req := xrpc.NewRequest(bytes.NewReader(content), logId)
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	fmt.Println("response:", resp)
	return nil
}
