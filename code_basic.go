package decibel_goods

import (
	"github.com/go-xorm/xorm"
	"github.com/myafeier/log"
	"time"
)

type CodeTypeBit int8

const (
	CodeTypeOfBarCode    CodeTypeBit = 1 << 0 //条码
	CodeTypeOfQrCode     CodeTypeBit = 1 << 1 //二维码
	CodeTypeOfWxMicroApp CodeTypeBit = 1 << 2 //微信小程序码
)


type ICodeModel interface {
	TableName()string
	GetOne(id, merchantId, goodsId, skuId int64, codeTypeBit CodeTypeBit, code string, hash string, lockForUpdate bool) (has bool, err error)
	Insert() (err error)
}



type CodeBase struct {
	Id          int64         `json:"id"`
	CodeTypeBit CodeTypeBit   `json:"code_type_bit" xorm:"default 0 index"`      //二维码类型
	MerchantId  int64         `json:"merchant_id" xorm:"default 0 index"`        // 商户ID
	GoodsId     int64         `json:"goods_id" xorm:"default 0 index"`           //商品id
	SkuId       int64         `json:"sku_id" xorm:"default 0 index"`             //skuId
	Code        string        `json:"code" xorm:"varchar(100) default '' index"` //码值
	Hash        string        `json:"hash" xorm:"varchar(100) default '' index"` //码值hash
	ImageUrl    string        `json:"image_url" xorm:"varchar(200) default ''"`  //码图片url
	Created     time.Time     `json:"created" xorm:"created"`                    //
	Updated     time.Time     `json:"updated" xorm:"updated"`                    //
	Session     *xorm.Session `json:"-" xorm:"-"`
}

func (self *CodeBase)TableName()string{
	return "goods_code"
}

func (self *CodeBase) GetOne(id, merchantId, goodsId, skuId int64, codeTypeBit CodeTypeBit, code string, hash string, lockForUpdate bool) (has bool, err error) {
	if id > 0 {
		self.Session.ID(id)
	}
	if merchantId > 0 {
		self.Session.Where("merchant_id = ?", merchantId)
	}
	if goodsId > 0 {
		self.Session.Where("goods_id = ?", goodsId)
	}
	if skuId > 0 {
		self.Session.Where("sku_id = ?", skuId)
	}

	if code != "" {
		self.Session.Where("code =?", code)
	}
	if hash != "" {
		self.Session.Where("hash =?", hash)
	}
	if codeTypeBit != 0 {
		self.Session.Where("code_type_bit & ? !=0", codeTypeBit)
	}
	if lockForUpdate {
		self.Session.ForUpdate()
	}

	has, err = self.Session.Get(self)
	if err != nil {
		log.Error(err.Error())
	}
	return
}

func (self *CodeBase) Insert() (err error) {
	_, err = self.Session.Insert(self)
	if err != nil {
		log.Error(err.Error())
	}
	return
}
