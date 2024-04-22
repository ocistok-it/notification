package apicall

import (
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"net/http"
)

func (a *Apicall) Post(url string, request interface{}, headers http.Header) ([]byte, error) {
	body, err := a.encodeRequest(request)
	if err != nil {
		return nil, err
	}

	headers.Add("Content-Type", "application/json")

	res, err := a.client.Post(url, body, headers)
	if err != nil {
		return nil, custerr.New("API_500", err.Error())
	}

	return a.consumeResponse(res)
}
