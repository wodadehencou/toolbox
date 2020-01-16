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

// FromBig converts a *big.Int into a []uint64 in little endian
func FromBig(out []uint64, big *big.Int) { //little endian
	for i := range out {
		out[i] = 0
	}

	if bits.UintSize == 32 {
		for i, v := range big.Bits() {
			if i&0x01 == 0x00 {
				out[i/2] |= uint64(v)
			} else {
				out[i/2] |= (uint64(v) << 32)
			}
		}
	} else { // bits.UintSize == 64
		for i, v := range big.Bits() {
			out[i] = uint64(v)
		}
	}
}
