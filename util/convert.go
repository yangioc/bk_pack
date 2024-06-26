package util

import (
	"encoding/hex"
	"time"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

func Big5ToUtf8(source []byte) ([]byte, error) {
	big5Toutf8 := traditionalchinese.Big5.NewDecoder()
	data, _, err := transform.Bytes(big5Toutf8, source)
	return data, err
}

func Utf8ToBig5(source []byte) ([]byte, error) {
	encoder := traditionalchinese.Big5.NewEncoder()
	data, _, err := transform.Bytes(encoder, source)
	return data, err
}

func Big5ToUtf8Str(source string) (string, error) {
	big5Toutf8 := traditionalchinese.Big5.NewDecoder()
	data, _, err := transform.String(big5Toutf8, source)
	return data, err
}

func Utf8ToBig5Str(source string) (string, error) {
	encoder := traditionalchinese.Big5.NewEncoder()
	data, _, err := transform.String(encoder, source)
	return data, err
}

// 為了與 Js 統一時間單位的處理

// 只能處理到秒數
// 帶入毫秒級會被處理到只剩下秒級。
// 帶入奈秒級會被處理到只剩下秒級。
func ParseUnixSec(t int64) time.Time {
	// utils.ServerTimeNow().UnixNano() => 1257894000000000000
	// utils.ServerTimeNow().Unix() =>     1257894000
	// js date => 1591069259005
	pTime := t
	if pTime >= 1e18 {
		pTime /= 1e8
	} else if pTime >= 1e12 {
		pTime /= 1e3
	}
	return time.Unix(pTime, 0).UTC()
}

// 統一時間單位長度 13碼
func ParseJavaUnixSec(t time.Time) int64 {
	return t.UnixNano() / 1000000
}

// 是否跨日
func IsTransDate(startTime, endTime time.Time) bool {
	if startTime.Year() != endTime.Year() {
		return true
	}
	if startTime.Month() != endTime.Month() {
		return true
	}
	if startTime.Day() != endTime.Day() {
		return true
	}
	return false
}

func HexToByte(hexString string) ([]byte, error) {
	return hex.DecodeString(hexString)
}
