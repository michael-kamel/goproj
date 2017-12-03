package repositories

import (
	"goproj/models"
	"goproj/api"
	"strconv"
	"encoding/json"
)

type AddListingRequest struct {
	models.OwnerInfo `json:"ownerInfo"`
	Category string `json:"category"`
	Location string `json:"location"`
	Space int `json:"space"`
	Price int `json:"price"`
	Address string `json:"address"`
	Description string `json:"description"`
}
type GetListingBySpecificationResponse struct {
	Success bool `json:"success"`
	Listings []models.Listing `json:"listings"`
}
type AddListingResponse struct {
	Success bool
}

type ListingAPIService struct {
	ApiService api.APIService
	GetURL string
	PostURL string
}

func (this *ListingAPIService) GetListings(listingSpec models.ListingSpecification) ([]models.Listing, error){
	query := make(map[string]string)
	query["category"] = listingSpec.Category
	query["location"] = listingSpec.Location
	space := strconv.Itoa(listingSpec.Space)
	price := strconv.Itoa(listingSpec.Price)
	query["space"] = space
	query["price"] = price
	response, err := this.ApiService.GetData(this.GetURL, query)
	if err != nil {
		return nil, &unFulfilledRequest{
			"GetListings",
			err.Error(),
		}
	}
	data := GetListingBySpecificationResponse{}
	json.Unmarshal(response, &data)
	return data.Listings, nil
}
func (this *ListingAPIService) AddListing(listing models.Listing) error {
	request := AddListingRequest{
		listing.OwnerInfo,
		listing.Category,
		listing.Location,
		listing.Space,
		listing.Price,
		listing.Address,
		listing.Description,
	}
	data, _ := json.Marshal(request)
	_, err := this.ApiService.PostData(this.PostURL, data)
	if err != nil {
		return &unFulfilledRequest{
			"AddListings",
			err.Error(),
		}
	}
	return nil
}