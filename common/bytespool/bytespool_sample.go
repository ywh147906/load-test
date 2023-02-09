package bytespool

import (
	"math"
	"math/bits"
	"sync"
)

var builtinPool Pool

// Pool consists of 32 sync.Pool, representing byte slices of length from 0 to 32 in powers of 2.
type Pool struct {
	pools [32]sync.Pool
}

// Get returns a byte slice with given length from the built-in pool.
func GetSample(size int) []byte {
	return builtinPool.Get(size)
}

// Put returns the byte slice to the built-in pool.
func PutSample(buf []byte) {
	builtinPool.Put(buf)
}

// Get retrieves a byte slice of the length requested by the caller from pool or allocates a new one.
func (p *Pool) Get(size int) []byte {
	if size <= 0 {
		return nil
	}
	if size > math.MaxInt32 {
		return make([]byte, size)
	}
	idx := index(uint32(size))
	if v := p.pools[idx].Get(); v != nil {
		bp := v.(*[]byte)
		return (*bp)[:size]
	}
	return make([]byte, 1<<idx)[:size]
}

// Put returns the byte slice to the pool.
func (p *Pool) Put(buf []byte) {
	size := cap(buf)
	if size == 0 || size > math.MaxInt32 {
		return
	}
	idx := index(uint32(size))
	if size != 1<<idx { // this byte slice is not from Pool.Get(), put it into the previous interval of idx
		idx--
	}
	p.pools[idx].Put(&buf)
}

func index(n uint32) uint32 {
	return uint32(bits.Len32(n - 1))
}
