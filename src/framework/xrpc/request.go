package xrpc

import (
	"bytes"
	"io"
)

type Request struct {
	Header Header
	Body   io.Reader
}

func NewRequest(body io.Reader, logId uint32) *Request {
	req := new(Request)
	req.Header.LogId = logId
	req.Header.MagicNum = HEADER_MAGIC_NUM

	if body != nil {
		switch v := body.(type) {
		case *bytes.Buffer:
			req.Header.BodyLen = uint32(v.Len())
		case *bytes.Reader:
			req.Header.BodyLen = uint32(v.Len())
		default:
			return nil
		}
		req.Body = io.LimitReader(body, int64(req.Header.BodyLen))
	}
	return req
}
