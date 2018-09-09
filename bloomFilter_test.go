/*
@Time : 2018/9/6 上午9:57
@Author : heikexue
@File : bloomFilter_test
@Software: GoLand
*/
package bitmap

import (
	"bufio"
	"code.byted.org/gopkg/logs"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCreateBloom(t *testing.T) {
	bitmap := CreateBitMap()
	bloom := CreateBloom(bitmap)
	if bloom.bitMap.size > 0 {
		fmt.Println("bloom create succ")
	}
}

func TestBloom_IsContainer(t *testing.T) {
	bitmap := CreateBitMap()
	bloom := CreateBloom(bitmap)

	if bloom.IsContainer("ssss") {
		fmt.Println("sss in bloom")
	} else {
		fmt.Println("sss not in bloom")
	}
}

func TestBloom_Insert(t *testing.T) {
	bitmap := CreateBitMapSize(2 << 25)
	bloom := CreateBloom(bitmap)

	input, err := os.Open("gids")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer input.Close()

	inputReader := bufio.NewReader(input)
	for {
		inputString, readerError := inputReader.ReadString('\n')
		// 按行读取  inputString末尾有\n
		if bloom.Insert(inputString) {
			//插入成功
			//fmt.Println("insert succ", inputString)
		} else {
			fmt.Println("insert err", inputString)
		}
		if readerError == io.EOF {
			break
		}
	}

	logs.Flush()
}
