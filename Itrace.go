package decibel_goods

import "time"

type ITrace interface {
	GetTraceInfo(goodsId int64)(brief string, created time.Time ,desc []Description,err error)
	AddTrace(goodsId int64,brief string,desc []Description)error
	CheckInfoValid()bool
}
