package toolbox

func CopyConcat(sl1, sl2 []byte) []byte {
	out := make([]byte, len(sl1)+len(sl2))
	copy(out[0:], sl1)
	copy(out[len(sl1):], sl2)
	return out
}

func AppendConcat(sl1, sl2 []byte) []byte {
	return append(sl1, sl2...)
}

func CopyConcatBytes(sl ...[]byte) []byte {
	var totalLen int
	for _, s := range sl {
		totalLen += len(s)
	}
	tmp := make([]byte, totalLen)
	var i int
	for _, s := range sl {
		i += copy(tmp[i:], s)
	}
	return tmp
}

func AppendConcatBytes(sl ...[]byte) []byte {
	if len(sl) == 0 {
		return nil
	}
	if len(sl) == 1 {
		return append(sl[0][:0:0], sl[0]...)
	}
	var out []byte
	out = append(sl[0], sl[1]...)
	for i := 2; i < len(sl); i++ {
		out = append(out, sl[i]...)
	}
	return out
}
