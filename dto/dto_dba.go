package dto

// 對 DBA 請求並等待回應
type Dto_DBA_Req struct {
	Request string `json:"request"`
	Type    string `json:"type"`
	Data    []byte `json:"data"`
}

// DBA 回應
type Dto_DBA_Res struct {
	Request string `json:"request"`
	Data    []byte `json:"data"`
	State   int    `json:"state"`
	Message string `json:"message"`
}
