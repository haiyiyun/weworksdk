package logic

import wework "github.com/haiyiyun/weworksdk"

func CancelAuthEventLogic(data []byte, ww wework.IWeWork) {
	ww.Logger().Sugar().Info(string(data))
}
