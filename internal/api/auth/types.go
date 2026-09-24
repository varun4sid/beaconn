package auth

type AuthRequest struct {
	Rollno string `json:"rollno"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status  string `json:"status"`
	Username string `json:"username,omitempty"`
}