package toolbox

import (
	"math/big"
	"math/bits"
)

// magic: code from math.big
const bigWordSize = (bits.UintSize / 8) // word size in bits

func BigToSlice(in *big.Int, slice []byte) {
	i := len(slice)
	for _, d := range in.Bits() {
		for j := 0; j < bigWordSize && i > 0; j++ {
			i--
			slice[i] = byte(d)
			d >>= 8
		}
	}
}
