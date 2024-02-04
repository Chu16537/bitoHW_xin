package proto

import "bitohw_xin/app/consts"

type Person struct {
	ID                int              `json:"id"`                   // 編號
	Name              string           `json:"name"`                 // 名稱
	Height            int              `json:"height"`               // 身高
	Gender            int              `json:"gender"`               // 性別 1男 2女
	WantedDates       int              `json:"wanted_dates"`         // 可執行約會次數
	IsSingle          bool             `json:"is_single"`            // 是否單身
	MatchStatus       int              `json:"match_status"`         // 配對狀態
	NotMatchPersonIDs map[int]struct{} `json:"not_match_person_ids"` // 不要配對的id
}

func NewPerson(id int, name string, height int, gender int, wantedDates int) *Person {
	return &Person{
		ID:                id,
		Name:              name,
		Height:            height,
		Gender:            gender,
		WantedDates:       wantedDates,
		IsSingle:          true,
		MatchStatus:       consts.MatchStatus_Match,
		NotMatchPersonIDs: make(map[int]struct{}),
	}
}

// 設定 IsSingle
func (p *Person) SetSingle(b bool) {
	p.IsSingle = b
}

// 設定 MatchStatus
func (p *Person) SetMatchStatus(s int) {
	p.MatchStatus = s
}

// 添加已經配對的id
func (p *Person) AddnotMatchPersonIDs(ids []int) {
	for _, v := range ids {
		p.NotMatchPersonIDs[v] = struct{}{}
	}
}

// 是否可以配對
func (p *Person) IsMatch() bool {
	return p.MatchStatus == consts.MatchStatus_Match && p.WantedDates > 0 && p.IsSingle
}

// 進行配對
func (p *Person) Match(count int) {
	p.WantedDates -= count

	if p.WantedDates < 0 {
		p.WantedDates = 0
	}

	if p.WantedDates == 0 {
		p.MatchStatus = consts.MatchStatus_NotMatch
	}
}
