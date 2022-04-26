package objects

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Passwd    string `json:"password"`
	AuthToken string `json:"authToken"`
}

type Customer struct {
	Id          string `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Email       string `json:"email"`
	Nationality string `json:"nationality"`
	Address     string `json:"address"`
}

type Affiliate struct {
	Id            string `json:"id"`
	AffiliateName string `json:"affiliateName"`
	District      string `json:"district"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phoneNumber"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
}

type Transaction struct {
	Id           string `json:"id"`
	SenderName   string `json:"senderName"`
	ReceiverName string `json:"receiverName"`
	Date         string `json:"date"`
	Money        string `json:"money"`
	Message      string `json:"message"`
}

type Document struct {
	Id               string `json:"id"`
	DocType          string `json:"docType"`
	DocNumber        string `json:"docNumber"`
	IssuingAuthority string `json:"issuingAuthority"`
	ExpiryDate       string `json:"expiryDate"`
	Img              string `json:"img"`
}
