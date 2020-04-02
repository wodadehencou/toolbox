package toolbox

import (
	"math/big"
	"runtime"
)

// ClearBig clear a big.Int from memory
func ClearBig(b *big.Int) {
	if b == nil {
		return
	}
	words := b.Bits()
	clearBigWords(words)
}

func clearBigWords(w []big.Word) {
	for i := range w {
		w[i] = 0
	}
	runtime.KeepAlive(w)
}

// ClearBytes clear a byte slice from memory
func ClearBytes(b []byte) {
	for i := range b {
		b[i] = 0
	}
	runtime.KeepAlive(b)
}
