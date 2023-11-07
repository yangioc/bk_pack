package log

// Log 輸出設置
// 數值越大log越詳細 max:255
var Level uint8 = 255

// Log 標籤設定
// 0: Test log 不管輸出等級一定會輸出
var TagMap map[uint8]struct{} = map[uint8]struct{}{}

var LevelToStringMap = map[string]uint8{
	"error": Level_Error,
	"info":  Level_Info,
	"debug": Level_Debug,
}

const (
	Level_Error = 0
	Level_Info  = 50
	Level_Debug = 100
)

const (
	Level_Error_str = "error"
	Level_Info_str  = "info"
	Level_Debug_str = "debug"
)

const (
	Tag_Test = 0 // 測試用標籤
)

type logFormat struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"msg"`
}
