package decibel_goods

import "time"

//商品跟踪记录
type GoodsTraceBlock struct {
	Id          int64          `json:"id"`
	GoodsId     int64          `json:"goods_id" xorm:"default 0 index"`         //商品id
	Brief       string         `json:"brief" xorm:"varchar(2000) default ''"`    //记录简介
	Description []*Description `json:"description" xorm:"json"`                 //记录内容
	ByUid       int64          `json:"by_uid" xorm:"default 0"`                 //上传用户
	Created     time.Time      `json:"created" xorm:"created"`                  //创建时间
	CurHash     string         `json:"cur_hash" xorm:"varchar(200) default ''"` //用户私钥后的HASH
	PreHash     string         `json:"pre_hash" xorm:"varchar(100) default ''"` //上个区块系统私钥后的HASH
}
