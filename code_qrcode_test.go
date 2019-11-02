package decibel_goods

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestCodeQrcode_GetOne(t *testing.T) {

	qrcode:=new(CodeQrCode)
	qrcode.Session=session
	qrcode.MerchantId=1
	qrcode.GoodsId=2
	qrcode.CodeTypeBit=CodeTypeOfQrCode
	qrcode.Code="asdfasdfywehwersfdsf"
	qrcode.Hash=fmt.Sprintf("%x",sha256.Sum256([]byte(qrcode.Code)))

	qrcode.Insert()

}
