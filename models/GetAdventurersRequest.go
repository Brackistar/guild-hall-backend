package models

type GetAdventurersRequest struct {
	IdList []uint64 `json:"id_list"`
}
