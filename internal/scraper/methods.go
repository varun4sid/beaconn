package scraper

import "fmt"

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

	return response.StatusCode, nil
}
