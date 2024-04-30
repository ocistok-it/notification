package dingtalk

import (
	"encoding/json"
	"fmt"
	"github.com/ocistok-it/notification/internal/infrastructure/custerr"
	"github.com/ocistok-it/notification/internal/infrastructure/dto"
	"net/http"
	"net/url"
)

func (s *Service) robotSend(accessToken string, payload interface{}) error {
	var (
		response dto.DingMessageResponse
	)

	value := url.Values{}
	value.Set("access_token", accessToken)

	uri := fmt.Sprintf("%s?%s", s.config.Endpoint, value.Encode())

	resp, err := s.apicall.Post(uri, payload, http.Header{})
	if err != nil {
		return custerr.New("error_invoke_api", err.Error())
	}

	if err = json.Unmarshal(resp, &response); err != nil {
		return custerr.New("json_decode", err.Error())
	}

	if response.Errcode != 0 {
		return custerr.New("error_api", fmt.Sprintf("error with code %v: %s", response.Errcode, response.Errmsg))
	}

	return nil
}

func (s *Service) getFormatImagePost(image string) string {
	return fmt.Sprint("![screenshot](", image, ")")
}

func (s *Service) getFormatPost(imgPost string, title string, content string) string {
	return fmt.Sprint(imgPost, " \n **", title, "** \n > ", content)
}
