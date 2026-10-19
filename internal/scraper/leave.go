package scraper

import (
	"encoding/json"
	"fmt"
	"net/url"
)

func (s *Session) GetManager() (string, error) {
	managerURL := s.baseURL + "/Hostel/Student/StudApprMngr?rollno=" + s.Rollno

	response, err := s.Client.Get(managerURL)
	if err != nil {
		return "", fmt.Errorf("failed to get manager: %w", err)
	}
	defer response.Body.Close()

	var managers []ManagerResponse
	err = json.NewDecoder(response.Body).Decode(&managers)
	if err != nil {
		return "", fmt.Errorf("failed to decode manager response: %w", err)
	}

	if len(managers) == 0 {
		return "", fmt.Errorf("no manager found")
	}

	return managers[0].Id, nil
}

func (s *Session) ApplyLeave(payload LeavePayload) error {
	leaveURL := s.baseURL + "/Hostel/Student/StudentLeavApply"

	form := payload.toValues()
	response, err := s.Client.PostForm(leaveURL, form)
	if err != nil {
		return fmt.Errorf("failed to apply leave: %w", err)
	}
	defer response.Body.Close()

	var jsonResponse []LeaveResponse
	decoder := json.NewDecoder(response.Body)
	err = decoder.Decode(&jsonResponse)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 || len(jsonResponse) == 0 {
		return fmt.Errorf("FAILED to apply leave, status code: %d", response.StatusCode)
	}

	if jsonResponse[0].Oresult != 1 {
		return fmt.Errorf("FAILED to apply leave, oresult: %d", jsonResponse[0].Oresult)
	}

	return nil
}

func (p *LeavePayload) toValues() url.Values {
	values := url.Values{}
	values.Set("rollno", p.Rollno)
	values.Set("fromDate", p.FromDate)
	values.Set("toDate", p.ToDate)
	values.Set("fromTime", p.FromTime)
	values.Set("toTime", p.ToTime)
	values.Set("fromSession", p.FromSession)
	values.Set("toSession", p.ToSession)
	values.Set("leaveType", p.LeaveType)
	values.Set("reason", p.Reason)
	values.Set("manager", p.Manager)

	return values
}
