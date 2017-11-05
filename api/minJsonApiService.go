package api

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"errors"
)

type MinJsonApiService struct{
	client *http.Client
	Headers map[string]string
}
func (this *MinJsonApiService) Init() {
	this.client = &http.Client{}
}
func (this *MinJsonApiService) GetData(url string, query map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
	}
	this.prepare(req)
	q := req.URL.Query()
	for key, value := range query {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := this.client.Do(req)
	if err != nil {
        return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Failed to get data from url:" + url)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
         return nil, err
	}
	resp.Body.Close()
	return body, nil
}
func (this *MinJsonApiService) PostData(url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
         return nil, err
	}
	this.prepare(req)
	resp, err := this.client.Do(req)
	if err != nil {
         return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("Failed to post data to url:" + url)
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
         return nil, err
	}
	resp.Body.Close()
	return responseBody, nil
}
func (this *MinJsonApiService) prepare(request *http.Request) {
	for key, value := range this.Headers {
		request.Header.Add(key, value)
	}
	request.Header.Set("Content-Type", "application/json")
}