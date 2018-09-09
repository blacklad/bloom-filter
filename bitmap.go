/*
@Time : 2018/9/5 下午8:29
@Author : heikexue
@File : bitmap
@Software: GoLand
*/
package bitmap

const BitmapSize = 0x01 << 32

type BitMap struct {
	data []uint64

	size uint64
}

func CreateBitMap() *BitMap {
	return CreateBitMapSize(BitmapSize)
}

func CreateBitMapSize(size uint64) *BitMap {
	if size > BitmapSize || size < 0 {
		size = BitmapSize
	}
	return &BitMap{
		data: make([]uint64, size>>6),
		size: uint64(size - 1),
	}
}

func (b *BitMap) GetMaps() []uint64 {
	return b.data
}

func (b *BitMap) SetBit(offset uint64, val uint8) bool {
	if b.size < offset {
		return false
	}
	index, pos := offset/64, offset%64

	if val == 0 {
		b.data[index] &^= 0x01 << pos
	} else {
		b.data[index] |= 0x01 << pos
	}
	return true
}

func (b *BitMap) GetBit(offset uint64) uint64 {
	if b.size < offset {
		return 0
	}
	index, pos := offset/64, offset%64
	return (b.data[index] >> pos) & 0x01
}
