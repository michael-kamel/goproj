package repositories

import "../models"

type BuyerRequestRepository interface {
	AddBuyerRequest(request models.BuyerRequest) error
}