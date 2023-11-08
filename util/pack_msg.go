package util

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/yangioc/bk_pack/proto/dtomsg"
	"github.com/yangioc/bk_pack/staticdata"
)

// MsgDecode 解析收到的封包

// @param data 來源的原始資料

// @return *DecodeResponse 解析完成的 code & payload (raw data)

// @return error
func MsgDecode(data []byte) (*dtomsg.Dto_Base, error) {
	dataLen := len(data)
	msg := &dtomsg.Dto_Base{}
	if dataLen < staticdata.UUIDLen {
		return msg, fmt.Errorf("data len=%v is invalid", dataLen)
	}

	uuid := string(data[:staticdata.UUIDLen])
	msg.UUID = strings.TrimSpace(uuid)
	if err := Unmarshal(data[staticdata.UUIDLen:], msg); err != nil {
		return msg, err
	}
	return msg, nil
}

// MsgEncode 將準備送出的訊息包裝成溝通好的格式

// @param f 訊息內容

// @return []byte 編碼後的位元組陣列

// @return error
func MsgEncode(data *dtomsg.Dto_Base) ([]byte, error) {
	if data.UUID == "" {
		return nil, errors.New("[MsgEncode] uuid is empty.")
	}

	uuidByte := make([]byte, staticdata.UUIDLen)
	copy(uuidByte, []byte(data.UUID))

	payload, err := Marshal(data)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(uuidByte)
	if _, err := buf.Write(payload); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}
