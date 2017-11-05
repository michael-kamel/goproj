package SessionManagement

type UserSession struct {
	UUID string
	State string
	Data struct {
		//general info
		BuyerOrSeller string //enum "Buyer", "Seller"
		MobileNumber string
		Address string
		
		//buyer specific data
		ItemLocation string

		//seller specific data

		//common data (for both buyers and sellers)
		ItemPrice int
	}
	RejectMessages []string //previous state needs to set this
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