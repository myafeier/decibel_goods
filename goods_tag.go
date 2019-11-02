package decibel_goods

import (
	"github.com/go-xorm/xorm"
	"time"
)

type GoodsTagBit int64 //限制了一个商户的最大标签数不超过64，如果要更长此处需要改成varchar

func (self GoodsTagBit) GetTagList(session *xorm.Session) (result []*GoodsTag, err error) {
	err = session.Where("tagBit &? !=0 ", self).Find(&result)
	return
}

//商品标签
type GoodsTag struct {
	Id         int64       `json:"id"`                                         //
	MerchantId int64       `json:"merchant_id" xorm:"default 0 index"`         //商户id
	TagBit     GoodsTagBit `json:"tag_bit" xorm:"default 0 index"`             //标示位,限制了一个商户的最大标签数不超过64
	Sequence   int8        `json:"sequence" xorm:"tinyint(2) default 0 index"` //最大256，越大越靠前
	Name       string      `json:"name" xorm:"varchar(200) default ''"`        //名称
	Icon       string      `json:"icon" xorm:"varchar(500) default ''"`        //图标
	Desc       string      `json:"desc" xorm:"varchar(2000) default ''"`       //简介
	Created    time.Time   `json:"created" xorm:"created"`                     //
	Updated    time.Time   `json:"updated" xorm:"updated"`                     //
}

func (self *GoodsTag) TableName() string {
	return "goods_tag"
}
