package api

import (
	"net/http"
	"encoding/json"
	"bytes"
)

type MinJsonApiService struct{
	client *http.Client
	Headers map[string]string
}
func (this *MinJsonApiService) Init() {
	this.client = &http.Client{}
}
func (this *MinJsonApiService) GetData(url string, query map[string]interface{}, data interface{}) {
	req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        panic(err)
	}
	this.prepare(req)
    q := req.URL.Query()
	req.URL.RawQuery = q.Encode()
	resp, err := this.client.Do(req)
	if err != nil {
        panic(err)
	}
	json.NewDecoder(resp.Body).Decode(data)
	resp.Body.Close()
}
func (this *MinJsonApiService) PostData(url string, body map[string]interface{}, data interface{}) {
	jsonString, err := json.Marshal(body)
	if err != nil {
        panic(err)
    }
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonString))
	this.prepare(req)
    if err != nil {
        panic(err)
    }
	resp, err := this.client.Do(req)
	if err != nil {
        panic(err)
	}
	json.NewDecoder(resp.Body).Decode(data)
	resp.Body.Close()
}
func (this *MinJsonApiService) prepare(request *http.Request) {
	for key, value := range this.Headers { 
		request.Header.Add(key, value)
	}
}
