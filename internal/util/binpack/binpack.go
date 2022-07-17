package binpack

type BinaryPackable interface {
	Marshall(into *Binpacker)
	Unmarshall(from *Binpacker) error
}

type Binpacker struct {
	binary []byte
	Err    error
	cursor int
}

func NewBinpackerAround(buffer []byte) *Binpacker {
	return &Binpacker{
		binary: buffer,
	}
}
func NewBinpacker(size int) *Binpacker {
	return &Binpacker{
		binary: make([]byte, size),
	}
}

func (bp *Binpacker) GetData() []byte {
	return bp.binary[:bp.cursor]
}
