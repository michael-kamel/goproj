package repositories

import "../models"

type ListingRepository interface {
	getListings(listing ListingSpecification) []models.Listing
	addListing(listing models.Listing)
}

type ListingSpecification struct {
	Category string
	Location string
	Space int
	Price int
}