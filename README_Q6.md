## API 文件

## querySinglePerson

動作: GET

參數

| key   | 說明     | 類型 |
| ----- | -------- | ---- |
| id    | 會員編號 | int  |
| count | 配對數量 | int  |

網址 http://127.0.0.1:8080/person/:id?count=3

## addSinglePersonAndMatch

動作: POST

參數

| key   | 說明     | 類型 |
| ----- | -------- | ---- |
| id    | 會員編號 | int  |
| count | 配對數量 | int  |

網址 http://127.0.0.1:8080/addSinglePersonAndMatch

參數格式

```
type AddPerson struct {
	Name        string `json:"name"`         // 名稱
	Height      int    `json:"height"`       // 身高
	Gender      int    `json:"gender"`       // 性別
	WantedDates int    `json:"wanted_dates"` // 可執行約會次數
}
```

## removeSinglePerson

動作: DELETE

參數

| key | 說明     | 類型 |
| --- | -------- | ---- |
| id  | 會員編號 | int  |

網址 http://127.0.0.1:8080/person/:id
