package decibel_goods

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/myafeier/log"
)

//码接口
type ICode interface {
	ICodeModel
	ICodeService
}


//码对外服务接口
type ICodeService interface {
	GetBasicInfo()(id, merchantId, goodsId, skuId int64,codeTypeBit CodeTypeBit)
	GenerateCode(merchantId, goodsId, skuId int64, code string, hash string)(err error)
}

//根据hash 或 code 查询码
func SearchCode(code,hash string,codeType CodeTypeBit,session *xorm.Session)(result ICode,has bool,err error){
	if session==nil{
		session=Engine.NewSession()
	}


	result=GenerateCodeProvider(codeType,session)
	if result==nil{
		err=fmt.Errorf("invalid codeType: %d",codeType)
		return
	}
	has,err=result.GetOne(0,0,0,0,codeType,code,hash,false)
	if err != nil {
		log.Error(err.Error())
		return
	}

	return
}

//生成码
func NewCode(merchantId,goodsId,skuId int64,code string,hashStr string,codeType CodeTypeBit,session *xorm.Session)(result ICode,err error){
	if session==nil{
		session=Engine.NewSession()
	}

	result=GenerateCodeProvider(codeType,session)
	if result==nil{
		err=fmt.Errorf("invalid codeType: %d",codeType)
		return
	}

	err=result.GenerateCode(merchantId,goodsId,skuId,code,hashStr)
	if err != nil {
		log.Error(err.Error())
	}
	return
}

func GenerateCodeProvider(bit CodeTypeBit,session *xorm.Session)ICode{
	switch bit {
	case CodeTypeOfBarCode:
		return &CodeBarCode{&CodeBase{Session:session}}
	case CodeTypeOfQrCode:
		return &CodeQrCode{&CodeBase{Session:session}}
	case CodeTypeOfWxMicroApp:
		return &CodeWxMicroApp{&CodeBase{Session:session}}
	default:
		return nil
	}
}