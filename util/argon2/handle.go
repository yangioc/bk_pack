package argon2

import (
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

// 推薦的加密選項
// 優點: 安全性高, 沒有不安全的係數
//
// 缺點: 處理時間較長

type Params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	salt        []byte
	keyLength   uint32
}

// 預設加密參數
var defaultParams = &Params{
	memory:      64 * 1024,          // 使用記憶體大小需要以 KiB 為單為增減
	iterations:  5,                  // 雜湊次數(過少的話有機率被暴力破解)
	parallelism: 2,                  // 執行併發數量(併發處理幣面)
	salt:        []byte("00000000"), // 雜訊
	keyLength:   32,                 // 最後雜湊數長度
}

func Encode(data []byte, params *Params) string {
	if params == nil {
		params = defaultParams
	}
	hash := argon2.IDKey(data, params.salt, params.iterations, params.memory, params.parallelism, params.keyLength)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	return b64Hash
}

func Compare(target, hash []byte, params *Params) bool {
	if params == nil {
		params = defaultParams
	}
	otherHash := argon2.IDKey(target, params.salt, params.iterations, params.memory, params.parallelism, params.keyLength)
	return subtle.ConstantTimeCompare(hash, otherHash) == 1
}
