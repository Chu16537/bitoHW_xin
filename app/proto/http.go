package proto

type AddPerson struct {
	Name        string `json:"name"`         // 名稱
	Height      int    `json:"height"`       // 身高
	Gender      int    `json:"gender"`       // 性別
	WantedDates int    `json:"wanted_dates"` // 可執行約會次數
}
