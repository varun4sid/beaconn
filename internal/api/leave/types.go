package leave

type LeaveRequest struct {
	Rollno   string  `json:"rollno"`
	Password string  `json:"password"`
	NumDays  *int    `json:"numDays"`
	FromDate *string `json:"fromDate"`
	ToDate   *string `json:"toDate"`
	Reason   *string `json:"reason"`
}

