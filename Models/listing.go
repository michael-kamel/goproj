package Models

type Listing struct {
	ID string
	OwnerInfo
	Category string
	Location string
	Space int
	Price int
	Description string
}