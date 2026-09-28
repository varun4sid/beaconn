package scraper

import (
	"net/http"
)

type Session struct {
	Client   *http.Client
	baseURL  string
	Rollno   string
	Password string
}

type StudentDetails struct {
	Name string `json:"name"`
}

type ManagerResponse struct {
	Id   string `json:"staff_id"`
	Name string `json:"staff_name"`
}

type LeavePayload struct {
	Rollno      string `json:"rollno"`
	FromDate    string `json:"fromDate"`
	FromTime    string `json:"fromTime"`
	FromSession string `json:"fromSession"`
	ToDate      string `json:"toDate"`
	ToTime      string `json:"toTime"`
	ToSession   string `json:"toSession"`
	LeaveType   string `json:"leaveType"`
	Reason      string `json:"reason"`
	Manager     string `json:"manager"`
}

type LeaveResponse struct {
	Oresult int `json:"oresult"`
}

type TokenDetails struct {
	Name     string
	Quantity string
}

type TokenData struct {
	Status       string         `json:"status"`
	QRCode       string         `json:"qrcode"`
	TokenDetails []TokenDetails `json:"tokenDetails"`
}
