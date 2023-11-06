package util

import (
	"github.com/yangioc/bk_pack/dto"
	"github.com/yangioc/bk_pack/msgpack"
)

func PackDBAReq(msgUid, request string, data []byte) ([]byte, error) {
	req, err := Marshal(&dto.Dto_DBA_Req{
		Request: request,
		Data:    data,
	})
	if err != nil {
		return nil, err
	}

	msg, err := MsgEncode(&msgpack.BasePack{
		UUID:    msgUid,
		Payload: req,
	})
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func UnpackDBARes(resData []byte) (*dto.Dto_DBA_Res, error) {
	msg, err := MsgDecode(resData)
	if err != nil {
		return nil, err
	}

	res := &dto.Dto_DBA_Res{}
	if err = Unmarshal(msg.Payload, res); err != nil {
		return nil, err
	}
	return res, nil
}
