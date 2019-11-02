package decibel_goods

import (
	"crypto/sha256"
	"fmt"
	"github.com/myafeier/log"
	"os"
	"time"
)

type CodeBarCode struct {
	*CodeBase
}

func (self *CodeBarCode) GetBasicInfo() (id, merchantId, goodsId, skuId int64, codeTypeBit CodeTypeBit) {
	fmt.Println(self.CodeBase)
	return self.CodeBase.Id, self.MerchantId, self.GoodsId, self.SkuId, CodeTypeOfBarCode
}

// code,hash 可以为空，为空时自动生成
func (self *CodeBarCode) GenerateCode(merchantId, goodsId, skuId int64, code string, hash string) (err error) {
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
	self.CodeTypeBit = CodeTypeOfBarCode

	if watcher, ok := GenerateWatchers[CodeTypeOfBarCode]; ok {
		var imageData []byte
		imageData, err = watcher.GenerateImage(merchantId, code, hash, self.Session)
		if err != nil {
			log.Error(err.Error())
			return
		}
		fileShortPath := fmt.Sprintf("%c%d%c%s.png", os.PathSeparator, merchantId, os.PathSeparator, hash)
		fileFullPath := fmt.Sprintf("%s%s", CodeImageSavePath,fileShortPath)

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
