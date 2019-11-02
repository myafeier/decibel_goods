package decibel_goods

import (
	"github.com/go-xorm/xorm"
	"github.com/myafeier/log"
)

type IGoods interface {
	IGoodsModel
	ITrace
}


func GenerateGoodsProvider(bit GoodsTypeBit,session *xorm.Session)IGoods{
	switch bit {
	case GoodsTypeOfStandard:
		return &GoodsStandard{&GoodsBasic{Session:session}}
	default:
		log.Error("invalid goodsTypeBit : %d",bit)
		return nil
	}
}











