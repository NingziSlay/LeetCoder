package main

import (
	"fmt"
	"sync"
)

// Bitmap 每个元素可以存储 32 个 bit 位，每个 bit 位表示一个 int32 类型的整数
// 如果使用 int32, 最高位是符号位，所以只能使用 31 个 bit 位
//
// uint32(0) 应该存放在 bitmap[0] 中的第一个 bit 位上
//     0 / 32 = 0 (确定索引位置)
//	   1 << (0 % 32) = 1 (确定 bit 位)
//
// uint32(3) 应该存放在 bitmap[0] 中的第四个 bit 位上
//     3 / 32 = 0 (确定索引位置)
//	   1 << (3 % 32) = 1000 (确定 bit 位)
//
// uint32(34) 应该存放在 bitmap[1] 中的第三个 bit 位上
//     34 / 32 = 1 (确定索引位置)
//	   1 << (34 % 32) = 100 (确定 bit 位)
//
// 整除是为了分区，每 32 位是一个区 (int32 类型最多存储 32 个 bit 位)
type Bitmap struct {
	lock   sync.Mutex
	bitmap []uint32
}

// BitLength uint32 的 bit 长度
const BitLength uint32 = 32

// NewBitmap 创建一个新的 Bitmap 对象，length 为 Bitmap 的初始大小
// 随着 item 的不断增加，Bitmap 会自动扩容
func NewBitmap(length uint) *Bitmap {
	return &Bitmap{
		bitmap: make([]uint32, length),
	}
}

// 计算 n 在对应索引位置中，应该存在哪个 bit 位上
//
// fun fact: 如果 x = 2 ^ y, 则 n % x 等价于 n & (x - 1)
// 如果 BitLength = 32 的话，可以使用:
// 		bitIndex := n & (BitLength - 1)
func (b *Bitmap) coordinate(n uint32) (uint32, uint32) {
	return n / BitLength, 1 << (n & (BitLength - 1))
}

// Add 将一个元素添加到 Bitmap 中
func (b *Bitmap) Add(n uint32) {
	x, y := b.coordinate(n)
	for x >= uint32(len(b.bitmap)) {
		b.bitmap = append(b.bitmap, 0)
	}
	// benchmark 的结果中，不加锁的速度要比加锁的速度快了 5 倍还多！
	b.lock.Lock()
	defer b.lock.Unlock()
	b.bitmap[x] |= y
}

// Del 将一个元素从 Bitmap 中移除
func (b *Bitmap) Del(n uint32) {
	x, y := b.coordinate(n)
	// 等价于 bitmap[bitmapIndex] = bitmap[bitmapIndex] &^ (1 << bitIndex)
	// x &= ^ y 表示 x 清空 y
	// x &= ^y 表示对 y 取反后与 x
	//b.lock.Lock()
	//defer b.lock.Unlock()
	b.bitmap[x] &= ^ y
}

// Contain 验证一个元素是否存在于 Bitmap 中
func (b *Bitmap) Contain(n uint32) bool {
	x, y := b.coordinate(n)
	return b.bitmap[x]&y != 0
}

func (b *Bitmap) String() string {
	var r string
	for _, x := range b.bitmap {
		r = fmt.Sprintf("%032b", x) + r
	}
	return r
}
