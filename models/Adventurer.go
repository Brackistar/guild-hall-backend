package models

type Adventurer struct {
	Id    uint64   `json:"id"`
	Name  string   `json:"name"`
	Image []byte   `json:"image"`
	Rank  AdvRank  `json:"rank"`
	Class AdvClass `json:"class"`
}
