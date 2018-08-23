package math

const (
	PrimePoly4 = 023
	PrimePoly8 = 0435
)

type ByteGaloisAlgebraImpl struct {
	gflog  []byte
	gfilog []byte
	size   byte
}

func New(w byte) *ByteGaloisAlgebraImpl {
	var primePoly int16
	switch w {
	case 4:
		primePoly = PrimePoly4
		break
	case 8:
		primePoly = PrimePoly8
		break
	}
	var size byte = 1 << w
	gflog := make([]byte, size)
	gfilog := make([]byte, size)

	var log, b int16

	for log, b = 0, 1; log < int16(size); log++ {
		gflog[b] = byte(log)
		gfilog[log] = byte(b)
		b = b << 1
		if b&int16(size) != 0 {
			b = b ^ primePoly
		}
	}
	return &ByteGaloisAlgebraImpl{gflog, gfilog, size}
}

func (alg *ByteGaloisAlgebraImpl) Add(x, y byte) byte {
	return x ^ y
}

func (alg *ByteGaloisAlgebraImpl) Sub(x, y byte) byte {
	return x ^ y
}

func (alg *ByteGaloisAlgebraImpl) Mul(x, y byte) byte {
	return alg.gfilog[(alg.gflog[x]+alg.gflog[y])%(alg.size-1)]
}

func (alg *ByteGaloisAlgebraImpl) Div(x, y byte) byte {
	return byte(alg.gfilog[(alg.gflog[x]-alg.gflog[y]+alg.size-1)%(alg.size-1)])
}
