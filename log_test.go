package gelf

import (
	"bytes"
	"encoding/json"
	"testing"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"
)

const (
	debugMsg = "Debug msg"
)

func TestUDPWriter(t *testing.T) {
	r, err := gelf.NewReader("127.0.0.1:0")
	if err != nil {
		t.Error(err)
	}
	cfg := map[string]interface{}{
		Namespace: map[string]interface{}{
			"addr":       r.Addr(),
			"enable_tcp": false,
		},
	}
	wr, err := NewWriter(cfg)
	if err != nil {
		t.Error(err)
	}

	wr.Write([]byte(debugMsg))
	msg, err := r.ReadMessage()
	if err != nil {
		t.Error(err)
	}
	buff := &bytes.Buffer{}
	if err := msg.MarshalJSONBuf(buff); err != nil {
		t.Error(err)
	}
	result := buff.Bytes()

	var message map[string]interface{}
	if err = json.Unmarshal(result, &message); err != nil {
		t.Errorf("Unable to Unmarshal message: %s", err.Error())
	}

	v, ok := message["short_message"]
	if !ok {
		t.Error("unable to decode message")
	}
	resultMsg, ok := v.(string)
	if !ok {
		t.Error("invalid short_message format")
	}
	if resultMsg != debugMsg {
		t.Errorf("the short_message field %s should be equal to %s", resultMsg, debugMsg)
	}
}

func TestInvalidConfig(t *testing.T) {
	cfg := map[string]interface{}{
		"Invalid Namespace": map[string]interface{}{},
	}
	_, err := NewWriter(cfg)
	if err != ErrWrongConfig {
		t.Errorf("The error should be %s, not %s", ErrWrongConfig.Error(), err.Error())
	}

	cfg = map[string]interface{}{
		Namespace: false,
	}
	_, err = NewWriter(cfg)
	if err != ErrWrongConfig {
		t.Errorf("The error should be %s, not %s", ErrWrongConfig.Error(), err.Error())
	}
}

func TestMissingAddr(t *testing.T) {
	cfg := map[string]interface{}{
		Namespace: map[string]interface{}{
			"enable_tcp": false,
		},
	}
	_, err := NewWriter(cfg)
	if err != ErrMissingAddr {
		t.Errorf("The error should be %s, not %s", ErrMissingAddr.Error(), err.Error())
	}
}
