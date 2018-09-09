/*
@Time : 2018/9/5 下午9:54
@Author : heikexue
@File : bithash_test
@Software: GoLand
*/
package bitmap

import (
	"fmt"
	"testing"
)

func TestGetBitHash(t *testing.T) {
	res := GetBitHash("abcd")
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestGetBitHashWithSize(t *testing.T) {
	res := GetBitHashWithSize("abcd", 64)
	for _, v := range res {
		fmt.Println(v)
	}
}

func TestGetBitHashWithSizeAndSeed(t *testing.T) {
	seed := []uint{1, 2, 3, 4}
	res := GetBitHashWithSizeAndSeed("abcd", 64, seed)
	for _, v := range res {
		fmt.Println(v)
	}
}
