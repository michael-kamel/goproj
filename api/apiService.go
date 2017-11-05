package api

type APIService interface {
	GetData(url string, query map[string]string) ([]byte, error)
	PostData(url string, body []byte) ([]byte, error)
	Init()
}
