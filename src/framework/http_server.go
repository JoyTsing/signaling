package framework

import (
	"fmt"
	"net/http"

	"signaling/src/glog"
)

type ActionInterface interface {
	Execute(w http.ResponseWriter, cr *ComRequest)
}

var GActionRouter = make(map[string]ActionInterface)

type ComRequest struct {
	R      *http.Request
	Logger *ComLog
	LogId  uint32
}

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path == "/favicon.ico" {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(""))
			return
		}

		if action, ok := GActionRouter[r.URL.Path]; ok {
			if action != nil {
				// wrap for logging
				cr := &ComRequest{R: r, Logger: &ComLog{}, LogId: GetLogId32()}
				LogUser(cr)

				r.ParseForm() // 获取表单数据
				for k, v := range r.Form {
					cr.Logger.AddLogItem("[Req]: "+k, v[0])
				}
				cr.Logger.TimeBegin("TotalCost")
				action.Execute(w, cr)
				cr.Logger.TimeEnd("TotalCost")

				cr.Logger.Infof("") //flush
			} else {
				responseError(w, r, http.StatusInternalServerError, "Internal Server Error")
			}
		} else {
			responseError(w, r, http.StatusNotFound, "Not Found")
		}
	})

}

func LogUser(cr *ComRequest) {
	r := cr.R
	cr.Logger.AddLogItem("LogId", fmt.Sprintf("%d", cr.LogId))
	cr.Logger.AddLogItem("Url", r.URL.Path)
	cr.Logger.AddLogItem("Referer", r.Header.Get("Referer"))
	cr.Logger.AddLogItem("Cookie", r.Header.Get("Cookie"))
	cr.Logger.AddLogItem("UA", r.Header.Get("User-Agent"))
	cr.Logger.AddLogItem("ClientIP", r.RemoteAddr)
	cr.Logger.AddLogItem("RealClientIP", getRealClientIP(r))
}

func getRealClientIP(r *http.Request) string {
	ip := r.RemoteAddr
	if rip := r.Header.Get("X-Real-IP"); rip != "" {
		ip = rip
	} else if rip := r.Header.Get("X-Forwarded-For"); rip != "" {
		ip = rip
	}
	return ip
}

func responseError(w http.ResponseWriter, r *http.Request, status int, err string) {
	w.WriteHeader(status)
	w.Write([]byte(fmt.Sprintf("%d - %s", status, err)))
}

func StartHttpServer() error {
	glog.Infof("[Start] HttpServer on port %d", gconf.httpPort)
	return http.ListenAndServe(fmt.Sprintf(":%d", gconf.httpPort), nil)
}
