package action

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"signaling/src/comerrors"
	"signaling/src/framework"
)

type comHttpResp struct {
	ErrNo  int         `json:"errNo"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

func writeJsonErrorResponse(cerr *comerrors.Errors, w http.ResponseWriter, cr *framework.ComRequest) {
	//log
	cr.Logger.AddLogItem("errNo", strconv.Itoa(cerr.Errno()))
	cr.Logger.AddLogItem("errMsg", cerr.Error())
	cr.Logger.Warningf("[Request] process failed: %s", cerr.Error())
	//response
	rsp := &comHttpResp{ErrNo: cerr.Errno(), ErrMsg: "process error"}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(rsp)
	w.Write(buffer.Bytes())
}
