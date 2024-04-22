package apicall

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"io"
	"net/http"
)

func (a *Apicall) encodeRequest(request interface{}) (*bytes.Buffer, error) {
	var reqBody bytes.Buffer
	if err := json.NewEncoder(&reqBody).Encode(request); err != nil {
		return nil, custerr.New("APICALL_ENCODE_REQ", "error encoding request").WithStacktrace(err)
	}
	return &reqBody, nil
}

func (a *Apicall) consumeResponse(resp *http.Response) ([]byte, error) {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, custerr.New("API_500", "error read body").WithStacktrace(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, custerr.New(
			fmt.Sprintf("API_%v", resp.StatusCode),
			fmt.Sprintf("error api call with code: %v", resp.StatusCode)).
			WithStacktrace(errors.New(string(respBody)))
	}

	return respBody, nil
}
