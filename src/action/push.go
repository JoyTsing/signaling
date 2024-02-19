package action

import (
	"fmt"
	"net/http"
	"strconv"

	"signaling/src/comerrors"
	"signaling/src/framework"
)

type pushAction struct{}

func NewPushAction() *pushAction {
	return &pushAction{}
}

func (p *pushAction) Execute(w http.ResponseWriter, cr *framework.ComRequest) {
	r := cr.R
	//uid
	var strUid string
	if values, ok := r.Form["uid"]; ok {
		strUid = values[0]
	}

	uid, err := strconv.ParseInt(strUid, 10, 64)
	if err != nil || uid <= 0 {
		cerr := comerrors.NewError(comerrors.ParamErr, "uid is invalid: "+err.Error())
		writeJsonErrorResponse(cerr, w, cr)
		return
	}
	// streamName
	var streamName string
	if values, ok := r.Form["streamName"]; ok {
		streamName = values[0]
	}
	if streamName == "" {
		cerr := comerrors.NewError(comerrors.ParamErr, "streamName is empty")
		writeJsonErrorResponse(cerr, w, cr)
		return
	}

	//audio & video
	var strAudio, strVideo string
	var audio, video int

	if values, ok := r.Form["audio"]; ok {
		strAudio = values[0]
	}

	if strAudio == "" || strAudio == "0" {
		audio = 0
	} else {
		audio = 1
	}

	if values, ok := r.Form["video"]; ok {
		strVideo = values[0]
	}

	if strVideo == "" || strVideo == "0" {
		video = 0
	} else {
		video = 1
	}
	fmt.Println("uid", uid, "streamname", streamName, "audio:", audio, "video:", video)
}
