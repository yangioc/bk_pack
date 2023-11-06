package msgpack

type BasePack struct {
	UUID           string        `json:"-"`           // 訊號唯一編號
	From           ServiceName   `json:"from"`        // 訊號來源
	Router         []ServiceName `json:"router"`      // 訊號經過的服務
	StartTime      int64         `json:"datetime"`    // 封包創建時間 *13碼時間
	ExpirationTime int64         `json:"expdatetime"` // 任務過期時間 *13碼時間
	Payload        []byte        `json:"payload"`     // 內容
}

type ServiceName string
