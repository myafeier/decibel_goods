package decibel_goods

import "time"

type GoodsSku struct {
	Id         int64
	Uuid       string                     //库存唯一码
	GoodsId    int64                      //商品id
	MerchantId int64                      //商户id
	StoreId    int64                      //如果storeId=0 就说明是商户的总库存
	GoodsAttr  []*GoodsSkuAttr            //商品特性
	Price      int                        //售价
	Amount     int                        //库存
	Created    time.Time `xorm:"created"` //
	Updated    time.Time `xorm:"updated"` //
}

type GoodsSkuAttr struct {
	GoodsAttrId int64
	Value       string
}
