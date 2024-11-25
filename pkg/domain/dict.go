package domain

type Word struct {
	ID          int    `json:"word_id"`
	RusWord     string `json:"rus_word" db:"rus_word"`
	Translation string `json:"translation" db:"translation"`
}

type CreateWordRequest struct {
	RusWord     string `json:"rus_word"`
	Translation string `json:"translation"`
}

type DeleteWordRequest struct {
	ID int `json:"id"`
}
