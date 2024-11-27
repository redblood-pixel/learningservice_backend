package domain

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGroupRequest struct {
	Name  string `json:"name"`
	Words []int  `json:"words"` // id of words in group
}

type DeleteGroupRequest struct {
	ID int `json:"id"`
}
