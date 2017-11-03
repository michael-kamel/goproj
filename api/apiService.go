package api

type APIService interface {
	getData(url string, query map[string]string, data interface{})
	postData(url string, body map[string]string, data interface{})
	init()
}
