package repositories

import "goproj/models"

type BuyerRequestRepository interface {
	AddBuyerRequest(request models.BuyerRequest) error
}