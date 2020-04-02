package toolbox

import (
	"math/big"
	"runtime"
	"testing"
)

func Benchmark_ClearBytes(b *testing.B) {
	bs := make([]byte, 1000*1000)
	for i := 0; i < b.N; i++ {
		ClearBytes(bs)
	}
}

func simpleClearBytes(bs []byte) {
	for i := range bs {
		bs[i] = 0
	}
	runtime.KeepAlive(bs)
}

func Benchmark_SimpleClearBytes(b *testing.B) {
	bs := make([]byte, 1000*1000)
	for i := 0; i < b.N; i++ {
		simpleClearBytes(bs)
	}
}

func Test_ClearBytes(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		bs := make([]byte, 1000*1000)
		ClearBytes(bs)
	})
	t.Run("nil", func(t *testing.T) {
		ClearBytes(nil)
	})
}

func Test_ClearBig(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		b := new(big.Int).SetInt64(100)
		ClearBig(b)
	})
	t.Run("more", func(t *testing.T) {
		b := new(big.Int).Lsh(big.NewInt(1), 200)
		ClearBig(b)
	})
	t.Run("zero", func(t *testing.T) {
		b := new(big.Int)
		ClearBig(b)
	})
	t.Run("nil", func(t *testing.T) {
		ClearBig(nil)
	})
}
