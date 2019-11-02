package decibel_goods

import (
	"github.com/go-xorm/xorm"
	"github.com/myafeier/log"
	"github.com/myafeier/qrcode"
)

//生成标准二维码的观察者

type WatcherGenerateQrCode struct {
	*qrcode.QrCode
}

func NewWatcherGenerateQrCode(fontByte,logoByte []byte, size int,fontSize float64,fontDPI float64 ,frontColor string,backColor string)(result *WatcherGenerateQrCode,err error){
	result=&WatcherGenerateQrCode{}
	result.QrCode,err=qrcode.Init(fontByte,logoByte,size,fontSize,fontDPI,frontColor,backColor)
	if err != nil {
		log.Error(err.Error())
	}
	return
}

func (self *WatcherGenerateQrCode) GenerateImage(merchantId int64, code, hash string, session *xorm.Session) (data []byte, err error) {
	if code!=""{
		data,err=self.GenerateQRCodeWithLogo(code)
	}else if hash==""{
		data,err=self.GenerateQRCodeWithLogo(hash)
	}
	if err != nil {
		log.Error(err.Error())
	}
	return
}
