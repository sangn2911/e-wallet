package objects

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Passwd    string `json:"password"`
	AuthToken string `json:"authToken"`
}

type Customer struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	DateOfBirth string `json:"dateOfBirth"`
	Email       string `json:"email"`
	Nationality string `json:"nationality"`
	Address     string `json:"address"`
}

type Affiliate struct {
	Id            int    `json:"id"`
	AffiliateName string `json:"affiliateName"`
	District      string `json:"district"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phoneNumber"`
	Fax           string `json:"fax"`
	Email         string `json:"email"`
}

type Transaction struct {
	Id           int    `json:"id"`
	SenderName   string `json:"senderName"`
	ReceiverName string `json:"receiverName"`
	Date         string `json:"date"`
	Money        string `json:"money"`
	Message      string `json:"message"`
}

type Document struct {
	Id               int    `json:"id"`
	DocType          string `json:"docType"`
	DocNumber        string `json:"docNumber"`
	IssuingAuthority string `json:"issuingAuthority"`
	ExpiryDate       string `json:"expiryDate"`
	Img              string `json:"img"`
	UserId           string `json:"userid"`
}

func (customer *Customer) Clone(obj Customer) {
	customer.LastName = obj.LastName
	customer.FirstName = obj.FirstName
	customer.DateOfBirth = obj.DateOfBirth
	customer.Email = obj.Email
	customer.Nationality = obj.Nationality
	customer.Address = obj.Address
}

func (user *User) Clone(obj User) {
	user.Username = obj.Username
	user.Email = obj.Email
	user.Passwd = obj.Passwd
}

func (tran *Transaction) Clone(obj Transaction) {
	tran.SenderName = obj.SenderName
	tran.ReceiverName = obj.ReceiverName
	tran.Date = obj.Date
	tran.Money = obj.Money
	tran.Message = obj.Message
}

func (doc *Document) Clone(obj Document) {
	doc.DocType = obj.DocType
	doc.DocNumber = obj.DocNumber
	doc.IssuingAuthority = obj.DocNumber
	doc.ExpiryDate = obj.ExpiryDate
	doc.Img = obj.Img
	doc.UserId = obj.UserId
}

func (affi *Affiliate) Clone(obj Affiliate) {
	affi.AffiliateName = obj.AffiliateName
	affi.District = obj.District
	affi.Address = obj.Address
	affi.PhoneNumber = obj.PhoneNumber
	affi.Fax = obj.Fax
	affi.Email = obj.Email
}
