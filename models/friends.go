package models

type FriendDTO struct {
	Fst int `json:"fst"`
	Snd int `json:"snd"`
}

type FriendsDTO struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Nickname string `json:"nickname"`
}
