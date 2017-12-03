package models

type Listing struct {
	Id string
	OwnerInfo
	Category string
	Location string
	Space int
	Price int
	Address string
	Description string
}