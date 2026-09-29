package account

type User struct {
	ID       int32    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Male     string `json:"male"`
	Birthday string `json:"birthday"`
}

type Friendship struct {
	Fst int32 `json:"fst"`
	Snd int32 `json:"snd"`
}
