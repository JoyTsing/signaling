package action

import (
	"fmt"
	"net/http"

	"signaling/src/framework"
)

type xrtcClientPushAction struct{}

func NewXrtcClientPushAction() *xrtcClientPushAction {
	return &xrtcClientPushAction{}
}

func (x *xrtcClientPushAction) Execute(w http.ResponseWriter, cr *framework.ComRequest) {
	fmt.Println("[Action] xrtcClientPushAction.Execute")
}
