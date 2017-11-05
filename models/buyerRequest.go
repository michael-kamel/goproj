package models

type BuyerRequest struct {
	OwnerInfo
	Listings []string
}