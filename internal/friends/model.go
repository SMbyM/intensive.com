package friends

type FriendDTO struct {
	Fst int32 `json:"fst"`
	Snd int32 `json:"snd"`
}

type FriendsDTO struct {
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
	Nickname string `json:"nickname"`
}
