package action

import (
	"html/template"
	"net/http"

	"signaling/src/framework"
	"signaling/third/glog"
)

type xrtcClientPushAction struct{}

func NewXrtcClientPushAction() *xrtcClientPushAction {
	return &xrtcClientPushAction{}
}

func (x *xrtcClientPushAction) Execute(w http.ResponseWriter, cr *framework.ComRequest) {
	t, err := template.ParseFiles(framework.GetStaticDir() + "/template/push.html")
	if err != nil {
		glog.Error("template.ParseFiles error: ", err)
		writeHtmlErrorResponse(w, http.StatusNotFound, "404 - Not found")
		return
	}

	req := make(map[string]string)
	for k, v := range cr.R.Form {
		req[k] = v[0]
	}

	if err = t.Execute(w, req); err != nil {
		glog.Error("template.ParseFiles error: ", err)
		writeHtmlErrorResponse(w, http.StatusNotFound, "404 - Not found")
		return
	}
}

func writeHtmlErrorResponse(w http.ResponseWriter, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(err))
}
