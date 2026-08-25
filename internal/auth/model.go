package auth

type UserProfile struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Male     string `json:"male"`
	Day      string `json:"day"`
	Month    string `json:"month"`
	Year     string `json:"year"`
}
