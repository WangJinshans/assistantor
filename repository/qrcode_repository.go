package repository

import (
	"assistantor/model"
	"encoding/json"
	"time"
)

func GenerateQrcodeContent(qrcodeId string, qrcodeName string) (bs []byte, err error) {
	info := make(map[string]interface{})
	info["code_id"] = qrcodeId
	info["code_name"] = qrcodeName
	info["time_stamp"] = time.Now().Unix()
	bs, err = json.Marshal(info)
	if err != nil {
		return
	}
	return
}

func SaveQrCode(qr model.QrCodeInfo) (err error) {
	err = engine.Save(&qr).Error
	return
}

func GetQrCode(qrCodeId string) (qrcode model.QrCodeInfo, err error) {
	err = engine.Where("qrcode_id = ?", qrCodeId).First(&qrcode).Error
	return
}
