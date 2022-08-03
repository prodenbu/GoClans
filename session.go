package GoClans

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Session struct {
	ApiKey string
}

func (s *Session) get(Url string) ([]byte, error) {
	url := "https://api.clashofclans.com/v1/" + Url

	req, _ := http.NewRequest(http.MethodGet, url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		body, _ := ioutil.ReadAll(res.Body)
		errBody := ClientError{}
		json.Unmarshal(body, &errBody)

		return []byte{}, fmt.Errorf("%d, %+v", res.StatusCode, errBody)
	}
	body, _ := ioutil.ReadAll(res.Body)

	return body, nil
}

func (s *Session) post(Url string, payload interface{}) ([]byte, error) {
	url := "https://api.clashofclans.com/v1/" + Url

	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(payload)
	if err != nil {
		return nil, err
	}
	req, _ := http.NewRequest(http.MethodPost, url, &buf)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+s.ApiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		return body, fmt.Errorf("got error with status %d", res.StatusCode)
	}
	return body, nil
}
