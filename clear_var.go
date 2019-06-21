package toolbox

import "math/big"

// ClearBig clear a big.Int from memory
func ClearBig(b *big.Int) {
	words := b.Bits()
	clearBigWords(words)
}

func clearBigWords(w []big.Word) {
	if len(w) == 0 {
		return
	}
	w[0] = 0
	for wp := 1; wp < len(w); wp *= 2 {
		copy(w[wp:], w[:wp])
	}
}

// ClearBytes clear a byte slice from memory
func ClearBytes(b []byte) {
	if len(b) == 0 {
		return
	}
	b[0] = 0
	for bp := 1; bp < len(b); bp *= 2 {
		copy(b[bp:], b[:bp])
	}
}
