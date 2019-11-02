package decibel_goods

import (
	"github.com/go-xorm/xorm"
	"github.com/myafeier/log"
	"os"
	"path/filepath"
)

const (
	Yes YesOrNo = 1 << 0 //是
	No  YesOrNo = 1 << 1 //否
)

type YesOrNo int8 //布尔开关

//码生成观察者
type IGenerateWatcher interface {
	GenerateImage(merchantId int64, code, hash string, session *xorm.Session) (imageData []byte, err error)
}

var Engine *xorm.Engine
var GenerateWatchers map[CodeTypeBit]IGenerateWatcher
var CodeImageSavePath string  //码图片保存路径
var CodeImageUrlPrefix string //码图片访问路径前缀
var CodePrefix string         //码前缀

func init() {
	log.SetPrefix("Decibel_Goods")
	log.SetLogLevel(log.DEBUG)
}

//初始化环境
func InitEnviroment(codeImageSavePath,codeImageUrlPrefix,codePrefix,logPrefix string, logLevel log.Level, engine *xorm.Engine) (err error) {
	if logPrefix!=""{
		log.SetPrefix("Decibel_Goods")

	}
	if logLevel!=0{
		log.SetLogLevel(logLevel)
	}

	CodeImageSavePath=codeImageSavePath
	CodeImageUrlPrefix=codeImageUrlPrefix
	CodePrefix=codePrefix

	if engine != nil {
		structs := []interface{}{
			CodeBase{}, GoodsBasic{},
		}
		Engine = engine
		for _, v := range structs {
			exist, err := engine.IsTableExist(v)
			if err != nil {
				panic(err.Error())
			}
			if !exist {
				engine.CreateTables(v)
				engine.CreateIndexes(v)
			} else {
				engine.Sync2(v)
			}
		}
	}


	return

}

//注册观察者
func AddWatcher(codeType CodeTypeBit, watcher IGenerateWatcher) {
	if GenerateWatchers == nil {
		GenerateWatchers = make(map[CodeTypeBit]IGenerateWatcher)
	}
	GenerateWatchers[codeType] = watcher
}


func WritePngFile(filePath string,imageData []byte)(err error){
	if _,err=os.Stat(filepath.Dir(filePath));err!=nil{
		if err!=nil{
			err=os.MkdirAll(filepath.Dir(filePath),os.ModePerm)
			if err != nil {
				log.Error(err.Error())
				return
			}
		}
	}else{
		if err != nil {
			log.Error(err.Error())
			return
		}
	}
	var imageFile *os.File
	imageFile,err=os.OpenFile(filePath,os.O_RDWR|os.O_CREATE,os.ModePerm)
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer imageFile.Close()
	_,err=imageFile.Write(imageData)
	if err != nil {
		log.Error(err.Error())
		return
	}
	return
}