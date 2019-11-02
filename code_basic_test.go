package decibel_goods

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestCodeBase_Insert(t *testing.T) {
	codebase:=new(CodeBase)
	codebase.Session=session
	codebase.MerchantId=1
	codebase.GoodsId=2
	codebase.CodeTypeBit=CodeTypeOfQrCode
	codebase.Code="asdfasdfywehwersfdsf"
	codebase.Hash=fmt.Sprintf("%x",sha256.Sum256([]byte(codebase.Code)))
	codebase.Insert()
}