package binpack

func (bp *Binpacker) MarshallString16(value string) {
	bs := []byte(value)
	l := len(bs)

	bp.MarshallUInt16(uint16(l))

	for i, b := range bs {
		bp.binary[bp.cursor+i] = b
	}

	bp.cursor += l
}
func (bp *Binpacker) UnmarshallString16() (string, error) {
	n, err := bp.UnmarshallUInt16()
	if err != nil {
		bp.Err = err
		return "", err
	}

	l := int(n)

	s := string(bp.binary[bp.cursor : bp.cursor+l])
	bp.cursor += l

	return s, nil
}

func (bp *Binpacker) MarshallString32(value string) {
	bs := []byte(value)
	l := len(bs)

	bp.MarshallUInt32(uint32(l))

	for i, b := range bs {
		bp.binary[bp.cursor+i] = b
	}

	bp.cursor += l
}
func (bp *Binpacker) UnmarshallString32() (string, error) {
	n, err := bp.UnmarshallUInt32()
	if err != nil {
		bp.Err = err
		return "", err
	}

	l := int(n)

	s := string(bp.binary[bp.cursor : bp.cursor+l])
	bp.cursor += l

	return s, nil
}
