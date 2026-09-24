package scraper

import (
	"fmt"
	"encoding/json"
)

func (s *Session) Login(rollno, password string) (int, error) {
	indexURL := s.baseURL + "/Hostel/Home/Index"
	authURL := s.baseURL + "/Hostel/Login/Authenticate"

	response, err := s.Client.Get(indexURL)
	if err != nil {
		return 0, fmt.Errorf("failed to get index page: %w", err)
	}
	response.Body.Close()

	payload := map[string][]string{
		"name":     {rollno},
		"password": {password},
	}

	response, err = s.Client.PostForm(authURL, payload)
	if err != nil {
		return 0, fmt.Errorf("failed to post login form: %w", err)
	}
	response.Body.Close()

	s.Rollno = rollno
	s.Password = password
	return response.StatusCode, nil
}

func (s *Session) GetStudentName() (string, error) {
	studURL := s.baseURL + "/Hostel/Student/studDetails?rollno=" + s.Rollno
	response, err := s.Client.Get(studURL)
	if err != nil {
		return "", fmt.Errorf("failed to get student details: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return "", fmt.Errorf("failed to get student details: status code %d", response.StatusCode)
	}

	var studentDetails []StudentDetails
	err = json.NewDecoder(response.Body).Decode(&studentDetails)
	if err != nil {
		return "", fmt.Errorf("failed to decode student details: %w", err)
	}

	if len(studentDetails) == 0 {
		return "", fmt.Errorf("no student details found for roll number %s", s.Rollno)
	}

	return studentDetails[0].Name, nil
}