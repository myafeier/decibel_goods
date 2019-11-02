package decibel_goods

import (
	"github.com/go-xorm/xorm"
	"github.com/myafeier/log"
	"time"
)

type GoodsTypeBit int8    //商品类型
type GoodsStateBit int8   //商品状态
type GoodsFeatureBit int8 //商品特性

const (
	GoodsTypeOfStandard    GoodsTypeBit = 1 << 0 //标准品
	GoodsTypeOfNonStandard GoodsTypeBit = 1 << 1 //非标品
	GoodsTypeOfVirtual     GoodsTypeBit = 1 << 2 //虚拟商品
)

const (
	GoodsStateOn  GoodsStateBit = 1 << 0 //上架
	GoodsStateOff GoodsStateBit = 1 << 1 //下架
)
const (
	GoodsFeatureOfNew       GoodsFeatureBit = 1 << 0 //商品特性-新品
	GoodsFeatureOfHot       GoodsFeatureBit = 1 << 1 //商品特性-热销
	GoodsFeatureOfRecommend GoodsFeatureBit = 1 << 2 //商品特性-推荐商品
)

type IGoodsModel interface {
	GetOne(id, merchantId, categoryId int64, tagBit GoodsTagBit, uuid string, name string, stateBit GoodsStateBit, typeBit GoodsTagBit, featureBit GoodsFeatureBit, lockForUpdate bool) (has bool, err error)
	Insert() (err error)
}

type GoodsBasic struct {
	Id            int64           `json:"id"`                                            //
	MerchantId    int64           `json:"merchant_id" xorm:"default 0 index"`            //商户ID
	CategoryId    int64           `json:"category_id" xorm:"default 0 index"`            //栏目ID
	BrandId       int64           `json:"brand_id" xorm:"default 0 index"`               //品牌id
	TagsBit       GoodsTagBit     `json:"tags_bit" xorm:"default 0 index"`               //拥有的标签
	Uuid          string          `json:"uuid" xorm:"varchar(100) default '' index"`     //商品唯一码
	Name          string          `json:"name" xorm:"varchar(500) default '' index"`     //商品名称
	StateBit      GoodsStateBit   `json:"state_bit" xorm:"tinyint(2) default 0 index"`   //商品状态
	UnitName      string          `json:"unit_name" xorm:"varchar(20) default ''"`       //单位名称
	TypeBit       GoodsTypeBit    `json:"type_bit" xorm:"tinyint(2) default 0 index"`    //商品类型
	FeatureBit    GoodsFeatureBit `json:"feature_bit" xorm:"tinyint(2) default 0 index"` //商品特性
	Sequence      int             `json:"sequence" xorm:" default 0 index"`              //排序
	Gallery       []*GoodsImage   `json:"gallery" xorm:"json"`                           //商品图片集
	Videos        []*GoodsVideo   `json:"videos" xorm:"json"`                            //商品视频集,打通腾讯视频
	Brief         string          `json:"brief" xorm:"varchar(2000) default ''"`         //简介
	Description   []*Description  `json:"description" xorm:"json"`                       //描述
	ClickCount    int64           `json:"click_count" xorm:"default 0 index"`            //点击量
	FavoriteCount int64           `json:"favorite_count" xorm:"default 0 index"`         //收藏量
	BuyCount      int64           `json:"buy_count" xorm:"default 0 index"`              //购买量
	ReturnCount   int64           `json:"return_count" xorm:"default 0 index"`           //退货量
	MarketPrice   int             `json:"market_price" xorm:"default 0"`                 //市场价
	MallPrice     int             `json:"mall_price" xorm:"default 0 index"`             //商城价格
	Attrs         []*GoodsAttr    `xorm:"-"`                                             //拥有的属性
	Session       *xorm.Session   `json:"-" xorm:"-"`
}

func (self *GoodsBasic) TableName() string {
	return "goods"
}

func (self *GoodsBasic) GetOne(id, merchantId, categoryId int64, tagBit GoodsTagBit, uuid string, name string, stateBit GoodsStateBit, typeBit GoodsTagBit, featureBit GoodsFeatureBit, lockForUpdate bool) (has bool, err error) {
	if id > 0 {
		self.Session.ID(id)
	}
	if merchantId > 0 {
		self.Session.Where("merchant_id = ?", merchantId)
	}
	if categoryId > 0 {
		self.Session.Where("category_id = ?", categoryId)
	}
	if tagBit > 0 {
		self.Session.Where("tags_bit & ? !=0", tagBit)
	}
	if uuid != "" {
		self.Session.Where("uuid =?", uuid)
	}
	if name != "" {
		self.Session.Where("name =?", name)
	}
	if stateBit != 0 {
		self.Session.Where("stat_bit & ? !=0", stateBit)
	}
	if typeBit != 0 {
		self.Session.Where("type_bit & ? !=0", typeBit)
	}
	if featureBit != 0 {
		self.Session.Where("feature_bit & ? !=0", featureBit)
	}
	if lockForUpdate {
		self.Session.ForUpdate()
	}
	has, err = self.Session.Get(&self)
	if err != nil {
		log.Error(err.Error())
	}
	return
}

func (self *GoodsBasic) Insert() (err error) {
	_, err = self.Session.Insert()
	if err != nil {
		log.Error(err.Error())
	}
	return
}

func (self *GoodsBasic) GetTraceInfo(goodsId int64) (brief string, created time.Time, desc []Description, err error) {

	return
}
func (self *GoodsBasic) AddTrace(goodsId int64, brief string, desc []Description) (err error) {

	return
}
func (self *GoodsBasic) CheckInfoValid() (valid bool) {

	return
}

type GoodsImage struct {
	Url      string `json:"url"`
	ThumbUrl string `json:"thumb_url"`
	Desc     string `json:"desc"`
}

type GoodsVideo struct {
	Url  string `json:"url"`
	Desc string `json:"desc"`
}

type Description struct {
	Type     string `json:"type"`      //可选值 text|image|video
	Content  string `json:"content"`   //内容， image 和 video 时都为 url
	Color    string `json:"color"`     //字体颜色，仅为text时有效
	FontSize int    `json:"font_size"` //字号，仅为text时有效
}
