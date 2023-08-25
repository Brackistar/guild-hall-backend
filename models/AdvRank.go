package models

type AdvRank struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Order       uint16 `json:"order"`
	PointsToGet uint32 `json:"points"`
}
