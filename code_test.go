package decibel_goods

import (
	"io/ioutil"
	"log"
	"sync"
	"testing"
)
var initMutex sync.Mutex
func initEnv()  {
	initMutex.Lock()
	defer initMutex.Unlock()
	var inited bool
	if !inited{
		err:=InitEnviroment("/Users/xiafei/test","http://locallhost","","",0,engine)
		if err!=nil{
			log.Panic(err)
		}
		log.Println(Engine)

		logo,err:=ioutil.ReadFile("/Users/xiafei/test/1080x1080.png")
		if err!=nil{
			log.Panic(err)
		}
		//注册默认的普通二维码观察者
		qrCodeWatcher,err:=NewWatcherGenerateQrCode(nil,logo,0,0,0,"","")
		if err != nil {
			log.Panic()
			return
		}
		AddWatcher(CodeTypeOfBarCode,qrCodeWatcher)
		inited=true
	}

}

func TestGenerateCode(t *testing.T) {
	initEnv()

	result,err:=NewCode(1,1,1,"","",CodeTypeOfBarCode,nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(result.GetBasicInfo())
}

func TestSearchCodeByHash(t *testing.T) {
	initEnv()
	result,has,err:=SearchCode("","f8f3b7671541c8ca115285312050742239cb3bf95f862b3a00020499c751ceaa",CodeTypeOfBarCode,nil)
	if err != nil {
		t.Error(err)
	}
	if has{
		t.Log(result.GetBasicInfo())
	}else{
		t.Log("no found")
	}
}