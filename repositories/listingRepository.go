package repositories

import "../models"

type ListingRepository interface {
	GetListings(listing models.ListingSpecification) ([]models.Listing, error)
	AddListing(listing models.Listing) error
}