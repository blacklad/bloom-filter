/*
@Time : 2018/9/5 下午9:35
@Author : heikexue
@File : bithash
@Software: GoLand
*/
package bitmap

import (
	"github.com/spaolacci/murmur3"
)

// 生成不同的hash值
func GetBitHash(str string) []uint64 {
	hashSeed := []uint{5, 11, 13, 61, 99}
	return GetBitHashWithSeed(str, hashSeed)
}

func GetBitHashWithSeed(str string, hashSeed []uint) []uint64 {
	hashRes := []uint64{}
	for _, val := range hashSeed {
		bithash := murmur3.New32WithSeed(uint32(val))
		bithash.Write([]byte(str))
		bithash.Sum32()
		hashRes = append(hashRes, uint64(bithash.Sum32()))
	}
	return hashRes
}

func GetBitHashWithSize(str string, size uint64) []uint64 {
	hasheRes := GetBitHash(str)
	for key, val := range hasheRes {
		hasheRes[key] = val % size
	}
	return hasheRes
}

func GetBitHashWithSizeAndSeed(str string, size uint64, seed []uint) []uint64 {
	hasheRes := GetBitHashWithSeed(str, seed)
	for key, val := range hasheRes {
		hasheRes[key] = val % size
	}
	return hasheRes
}
