package action

import (
	"encoding/json"
	"net/http"
	"strconv"

	"signaling/src/comerrors"
	"signaling/src/framework"
)

type pushAction struct{}

type pushData struct {
	Type string `json:"type"`
	Sdp  string `json:"sdp"`
}

func NewPushAction() *pushAction {
	return &pushAction{}
}

type xrpcPushRequest struct {
	Cmdno      int    `json:"cmd_no"`
	Uid        uint64 `json:"uid"`
	StreamName string `json:"stream_name"`
	Audio      int    `json:"audio"`
	Video      int    `json:"video"`
}

type xrpcPushResponse struct {
	Errno  int    `json:"err_no"`
	ErrMsg string `json:"err_msg"`
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
	req := xrpcPushRequest{
		Cmdno:      CMDNO_PUSH,
		Uid:        uint64(uid),
		StreamName: streamName,
		Audio:      audio,
		Video:      video,
	}
	var resp xrpcPushResponse
	if err = framework.Call("xrtc", req, &resp, cr.LogId); err != nil {
		cerr := comerrors.NewError(comerrors.NetworkErr, "backend process error"+err.Error())
		writeJsonErrorResponse(cerr, w, cr)
		return
	}
	//fmt.Printf("%+v\n", resp)
	if resp.Errno != 0 {
		cerr := comerrors.NewError(comerrors.NetworkErr, "backend process errno: "+string(rune(resp.Errno)))
		writeJsonErrorResponse(cerr, w, cr)
		return
	}
	//make response
	httpResp := comHttpResp{
		ErrNo:  0,
		ErrMsg: "",
		Data: pushData{
			Type: "offer",
			Sdp:  resp.Offer,
		},
	}
	b, _ := json.Marshal(httpResp)
	cr.Logger.AddLogItem("resp", string(b))
	w.Write(b)
}
