package proto

// 配對條件
type MatchFilter struct {
	ID                int              `json:"id"`                   // 編號
	Height            int              `json:"height"`               // 身高
	Gender            int              `json:"gender"`               // 性別 1男 2女
	Count             int              `json:"count"`                // 數量
	NotMatchPersonIDs map[int]struct{} `json:"not_match_person_ids"` // 不要配對的id
}

func NewMatchFilter(id int, name string, height int, gender int, count int, notMatchPersonIDs map[int]struct{}) *MatchFilter {
	return &MatchFilter{
		ID:                id,
		Height:            height,
		Gender:            gender,
		Count:             count,
		NotMatchPersonIDs: notMatchPersonIDs,
	}
}
