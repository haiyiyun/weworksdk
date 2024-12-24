package logic

import (
	"encoding/xml"

	wework "github.com/haiyiyun/weworksdk"
	"github.com/haiyiyun/weworksdk/pkg/svr/models"
)

type ResetPermanentCodeEvent struct {
	BizEvent
	AuthCode string `xml:"AuthCode"`
}

func ResetPermanentCodeEventLogic(data []byte, ww wework.IWeWork) {
	var event CreateAuthEvent
	if err := xml.Unmarshal(data, &event); err != nil {
		ww.Logger().Sugar().Error(err)
		return
	}
	ww.Logger().Sugar().Info(string(data))
	resp := ww.GetPermanentCode(event.AuthCode)
	var corpPermanentCode models.CorpPermanentCode // 授权企业永久授权码
	engine.Model(models.CorpPermanentCode{}).
		Where(models.CorpPermanentCode{AuthCorpId: resp.AuthCorpInfo.CorpId}).
		First(&corpPermanentCode)
	corpPermanentCode.PermanentCode = resp.PermanentCode
	engine.Save(&corpPermanentCode)
	//TODO:将授权企业的首次access token 写入缓存
	ww.Logger().Sugar().Info(resp)
}
