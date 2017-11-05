package repositories

import (
	"../models"
	"../api"
	"encoding/json"
)

type AddBuyerRequest struct {
	BuyerInfo models.OwnerInfo `json:"buyerInfo"`
	Listings []string `json:"listings"`
}

type AddBuyerRequestResponse struct {
	Success bool
}

type BuyerRequestAPIService struct {
	ApiService api.APIService
	PostURL string
}


func (this *BuyerRequestAPIService) AddBuyerRequest(request models.BuyerRequest) error {
	buyerRequest := AddBuyerRequest {
		request.OwnerInfo,
		request.Listings,
	}
	data, _ := json.Marshal(&buyerRequest)
	_, err := this.ApiService.PostData(this.PostURL, data)
	if err != nil {
		return &unFulfilledRequest{
			"AddBuyerRequest",
			err.Error(),
		}
	}
	return nil
}