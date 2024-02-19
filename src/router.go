package main

import (
	"signaling/src/action"
	"signaling/src/framework"
)

func init() {
	framework.GActionRouter["/xrtcclient/push"] = action.NewXrtcClientPushAction()
}
