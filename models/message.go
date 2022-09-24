package models

type Msg struct {
	User int64  `json:"user"`
	Data string `json:"text"`
	Chat int64  `json:"chat"`
}
type MsgDTO struct {
	User string `json:"user"`
	Data string `json:"text"`
	Chat string `json:"chat"`
}
