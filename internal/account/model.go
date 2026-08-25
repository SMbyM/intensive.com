package account

type UserProfile struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Male     string `json:"male"`
	Birthday string `json:"birthday"`
}

type Friendship struct {
	Fst int `json:"fst"`
	Snd int `json:"snd"`
}
