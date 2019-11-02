package decibel_goods

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"math/rand"
	"testing"
	"time"
)

const host= "127.0.0.1"
const port= "3306"
const user= "test"
const passwd= "test"
const dbName= "test"
var session *xorm.Session
var engine *xorm.Engine
func init()  {

	var err error
	engine, err = xorm.NewEngine("mysql", user+":"+passwd+"@tcp("+host+":"+port+")/"+dbName+"?charset=utf8mb4")
	if err != nil {
		panic("error when connect database!,err:" + err.Error())
	}
	engine.ShowSQL(true)

	strs:=[]interface{}{
		&CodeBase{},&GoodsBasic{},
	}

	session=engine.NewSession()
	for _,v:=range strs{
		exist,err:=engine.IsTableExist(v)
		if err != nil {
			panic( err.Error())
		}
		if !exist{
			engine.CreateTables(v)
			engine.CreateIndexes(v)
		}else{
			engine.Sync2(v)
		}
	}

}

func TestGoodsStandard_Insert(t *testing.T) {
	rand.Seed(time.Now().Unix())
	good:=new(GoodsStandard)
	good.Session=session
	good.GoodsBasic.Name=fmt.Sprintf("测试商品%d",rand.Int31n(10000))
	good.GoodsBasic.MerchantId=1
	good.GoodsBasic.Uuid=fmt.Sprintf("%10d",rand.Int63n(23423423434))
	good.Gallery=[]*GoodsImage{
		{
				Url:"https://via.placeholder.com/1080x1080.png",
				ThumbUrl:"https://via.placeholder.com/680x680.png",
				Desc:"测试图片",
		},
	}
	switch rand.Int31n(2){
	case 0:
		good.GoodsBasic.FeatureBit=GoodsFeatureOfHot|GoodsFeatureOfNew
	case 1:
		good.GoodsBasic.FeatureBit=GoodsFeatureOfRecommend
	case 2:
		good.GoodsBasic.FeatureBit=GoodsFeatureOfNew
	}

	err:=good.Insert()
	if err != nil {
		t.Error(err)
	}
}

func TestGoodsStandard_GetGood(t *testing.T) {
	good:=new(GoodsStandard)
	good.Session=session
	rand.Seed(time.Now().Unix())
	switch rand.Int31n(2) {
		case 0:
			has,err:=good.GetGood(0,0,0,0,"","",0,0,GoodsFeatureOfNew,false)
			if err != nil {
				t.Fatal(err)
			}
			if has{
				t.Logf("%+v",*good)
			}
	case 1:
		has,err:=good.GetGood(0,0,0,0,"","",0,0,GoodsFeatureOfHot|GoodsFeatureOfNew,false)
		if err != nil {
			t.Fatal(err)
		}
		if has{
			t.Logf("%+v",*good)
		}
	case 2:
		has,err:=good.GetGood(0,0,0,0,"","",0,0,GoodsFeatureOfRecommend,false)
		if err != nil {
			t.Fatal(err)
		}
		if has{
			t.Logf("%+v",*good)
		}

	}


}
