package decibel_goods

type GoodsNonStandard struct {
	*GoodsBasic `xorm:"extends"`
}


//称重
func (self *GoodsNonStandard)Weigh(){

}