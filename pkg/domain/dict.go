package domain

type Word struct {
	Id          int    `json:"id"`
	RusWord     string `json:"rus_word" db:"rus_word"`
	Translation string `json:"translation" db:"translation"`
}

type CreateWordRequest struct {
	RusWord     string `json:"rus_word"`
	Translation string `json:"translation"`
}

type DeleteWordRequest struct {
	Id int `json:"id"`
}
