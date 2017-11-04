package api

type APIService interface {
	GetData(url string, query map[string]interface{}, data interface{})
	PostData(url string, body map[string]interface{}, data interface{})
	Init()
}
