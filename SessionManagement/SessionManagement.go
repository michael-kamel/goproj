package SessionManagement

type UserSession struct {
	UUID string
	State string
	Data struct {
		BuyerOrSeller string //enum "Buyer", "Seller"
		
		//contact info
		Name string
		Phone string
		Email string

		//data
		Category string
		Location string
		Space int
		Price int
		Address string
		//seller
		Description string

		//buyer
	}
	RejectMessages []string //previous state needs to set this
	ReceivedItems []ReceivedItem
	ReceivedItemsMessage string
	ChosenItem string
}

type ReceivedItem struct {
	Number int
	ID string
	/*Name string
	Phone string
	Email string
	Category string
	Location string
	Space int
	Price int
	Description string*/
}

var UserSessions map[string]*UserSession = make(map[string]*UserSession)

/*func GetUserSession(uuid string) UserSession {
	return UserSessions[uuid]
}*/

func GenerateNewUserSession(uuid string) {
	UserSessions[uuid] = &UserSession{UUID:uuid, State:"phase1", RejectMessages:[]string{"Please type buy or sell."}}
}

func DeleteSession(uuid string) {
	delete(UserSessions, uuid)
}