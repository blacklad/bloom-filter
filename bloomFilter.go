/*
@Time : 2018/9/6 上午9:56
@Author : heikexue
@File : bloomFilter
@Software: GoLand
*/
package bitmap

import (
	"code.byted.org/gopkg/logs"
)

//type BloomFilter interface {
//	IsContainer(s string) bool
//	Insert(str string) bool
//}

type Bloom struct {
	bitMap *BitMap
}

func CreateBloom(bitMap *BitMap) *Bloom {
	return &Bloom{
		bitMap: bitMap,
	}
}

func (b *Bloom) IsContainer(s string) bool {
	bitHash := GetBitHashWithSize(s, b.bitMap.size)
	for _, v := range bitHash {
		if b.bitMap.GetBit(uint64(v)) == 0 {
			return false
		}
	}
	return true
}

func (b *Bloom) Insert(s string) bool {
	if !b.IsContainer(s) {
		bitHash := GetBitHashWithSize(s, b.bitMap.size)
		for _, v := range bitHash {
			b.bitMap.SetBit(uint64(v), 1)
		}
		return true
	} else {
		logs.Info("Bloom has %s", s)
	}
	return false
}
