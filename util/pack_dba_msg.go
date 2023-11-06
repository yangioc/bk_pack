package util

import (
	"github.com/yangioc/bk_pack/dto"
)

func PackDBAReq(request string, data []byte) ([]byte, error) {
	req, err := Marshal(&dto.Dto_DBA_Req{
		Request: request,
		Data:    data,
	})
	if err != nil {
		return nil, err
	}

	return req, nil
}

func UnpackDBARes(resData []byte) (*dto.Dto_DBA_Res, error) {
	res := &dto.Dto_DBA_Res{}
	if err := Unmarshal(resData, res); err != nil {
		return nil, err
	}
	return res, nil
}
