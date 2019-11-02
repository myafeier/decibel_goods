package decibel_goods

//商品属性-不同的栏目有不同的属性集
type GoodsAttr struct {
	Id         int64 `json:"id"` //
	Uuid       string `json:"uuid" xorm:"varchar(100) default '' "`  //属性唯一码
	MerchantId int64    //商户id，为1时为平台所有
	CateId     int64    //栏目
	AttrName   string   //属性名称
	AttrValues []string //属性可选值
	Sequence   int      //排序
}
