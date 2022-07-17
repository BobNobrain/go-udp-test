package binpack

import (
	"encoding/binary"
	"math"
)

func (bp *Binpacker) MarshallFloat32(value float32) {
	n := math.Float32bits(value)
	binary.LittleEndian.PutUint32(bp.binary[bp.cursor:], n)
	bp.cursor += 4
}
func (bp *Binpacker) UnmarshallFloat32() (float32, error) {
	n := binary.LittleEndian.Uint32(bp.binary[bp.cursor:])
	f := math.Float32frombits(n)
	bp.cursor += 4
	return f, nil
}

func (bp *Binpacker) MarshallUInt32(value uint32) {
	binary.LittleEndian.PutUint32(bp.binary[bp.cursor:], value)
	bp.cursor += 4
}
func (bp *Binpacker) UnmarshallUInt32() (uint32, error) {
	n := binary.LittleEndian.Uint32(bp.binary[bp.cursor:])
	bp.cursor += 4
	return n, nil
}

func (bp *Binpacker) MarshallUInt16(value uint16) {
	binary.LittleEndian.PutUint16(bp.binary[bp.cursor:], value)
	bp.cursor += 2
}
func (bp *Binpacker) UnmarshallUInt16() (uint16, error) {
	n := binary.LittleEndian.Uint16(bp.binary[bp.cursor:])
	bp.cursor += 2
	return n, nil
}

func (bp *Binpacker) MarshallByte(value byte) {
	bp.binary[bp.cursor] = value
	bp.cursor += 1
}
func (bp *Binpacker) UnmarshallByte() (byte, error) {
	b := bp.binary[bp.cursor]
	bp.cursor += 1
	return b, nil
}
