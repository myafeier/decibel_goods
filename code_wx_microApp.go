package decibel_goods

import (
	"crypto/sha256"
	"fmt"
	"github.com/myafeier/log"
	"os"
	"time"
)

type CodeWxMicroApp struct {
	*CodeBase
}

func (self *CodeWxMicroApp)GetBasicInfo()(id, merchantId, goodsId, skuId int64,codeTypeBit CodeTypeBit){
	return self.Id,self.MerchantId,self.GoodsId,self.SkuId,CodeTypeOfBarCode
}


// code,hash 可以为空，为空时自动生成
func (self *CodeWxMicroApp) GenerateCode(merchantId, goodsId, skuId int64, code string, hash string) (err error) {
	if code == "" {
		code = fmt.Sprintf("%s%d%d%d%d", CodePrefix, merchantId, goodsId, skuId, time.Now().Unix())
	}
	if hash == "" {
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(code)))
	}
	self.Code = code
	self.Hash = hash
	self.MerchantId = merchantId
	self.GoodsId = goodsId
	self.SkuId = skuId
	self.CodeTypeBit = CodeTypeOfQrCode

	if watcher, ok := GenerateWatchers[CodeTypeOfQrCode]; ok {
		var imageData []byte
		imageData, err = watcher.GenerateImage(merchantId, code, hash, self.Session)
		if err != nil {
			log.Error(err.Error())
			return
		}
		fileFullPath := fmt.Sprintf("%s%c%d%c%s", CodeImageSavePath, os.PathSeparator, merchantId, os.PathSeparator, hash)
		fileShortPath := fmt.Sprintf("%c%d%c%s", os.PathSeparator, merchantId, os.PathSeparator, hash)
		err = WritePngFile(fileFullPath, imageData)
		if err != nil {
			log.Error(err.Error())
			return
		}
		self.ImageUrl = CodeImageUrlPrefix + fileShortPath
	}

	err = self.Insert()
	if err != nil {
		log.Error(err.Error())
		return
	}

	return
}
