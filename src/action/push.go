package action

import (
	"net/http"
	"strconv"

	"signaling/src/comerrors"
	"signaling/src/framework"
)

type pushAction struct{}

func NewPushAction() *pushAction {
	return &pushAction{}
}

type xrtcPushRequest struct {
	Cmdno      int    `json:"cmd_no"`
	Uid        uint64 `json:"uid"`
	StreamName string `json:"stream_name"`
	Audio      int    `json:"audio"`
	Video      int    `json:"video"`
}

type xrtcPushResponse struct {
	Errno  int    `json:"err_no"`
	ErrMsg int    `json:"err_msg"`
	Offer  string `json:"offer"`
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

	// log：uid, streamName, audio, video
	// fmt.Println("uid", uid, "streamname", streamName, "audio:", audio, "video:", video)
	// 通过rpc拉流
	req := xrtcPushRequest{
		Cmdno:      CMDNO_PUSH,
		Uid:        uint64(uid),
		StreamName: streamName,
		Audio:      audio,
		Video:      video,
	}
	var resp xrtcPushResponse
	if err = framework.Call("xrtc", req, resp, cr.LogId); err != nil {
		cerr := comerrors.NewError(comerrors.NetworkErr, "backend process error")
		writeJsonErrorResponse(cerr, w, cr)
		return
	}
}
