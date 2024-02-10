// Code generated by ogen, DO NOT EDIT.

package lokiapi

import (
	"bytes"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	ht "github.com/ogen-go/ogen/http"
)

func encodePushRequest(
	req PushReq,
	r *http.Request,
) error {
	switch req := req.(type) {
	case *Push:
		const contentType = "application/json"
		e := new(jx.Encoder)
		{
			req.Encode(e)
		}
		encoded := e.Bytes()
		ht.SetBody(r, bytes.NewReader(encoded), contentType)
		return nil
	case *PushReqApplicationXProtobuf:
		const contentType = "application/x-protobuf"
		body := req
		ht.SetBody(r, body, contentType)
		return nil
	default:
		return errors.Errorf("unexpected request type: %T", req)
	}
}
