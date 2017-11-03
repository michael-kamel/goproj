package repositories

import (
	"../models"
	"../api"
	"strconv"
)

type ListingAPIService struct {
	apiService api.APIService
	getURL string
	postURL string
}
type ListingResponse struct {
    Collection []models.Listing
}


func (this *ListingAPIService) getListings(listingSpec ListingSpecification) []models.Listing{
	query := make(map[string]interface{})
	query["category"] = listingSpec.Category
	query["location"] = listingSpec.Location
	space := strconv.Itoa(listingSpec.Space)
	price := strconv.Itoa(listingSpec.Price)
	query["space"] = space
	query["price"] = price
	listings := ListingResponse{}
	this.apiService.GetData(this.getURL, query, &listings)
	return listings.Collection
}
func (this *ListingAPIService) postListings(listing models.Listing) {
	body := make(map[string]interface{})
	body["category"] = listing.Category
	body["location"] = listing.Location
	body["space"] = strconv.Itoa(listing.Space)
	body["price"] = strconv.Itoa(listing.Price)
	body["description"] = listing.Description
	ownerInfo := make(map[string]string)
	ownerInfo["name"] = listing.OwnerInfo.Name
	ownerInfo["phone"] = listing.OwnerInfo.Phone
	ownerInfo["email"] = listing.OwnerInfo.Email
	body["ownerInfo"] = ownerInfo
	this.apiService.PostData(this.postURL, body, nil)
}