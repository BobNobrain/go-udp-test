package binpack

func MarshallTagged(bp *Binpacker, tag byte, item BinaryPackable) {
	bp.MarshallByte(tag)
	item.Marshall(bp)
}
func UnmarshallTagged[T BinaryPackable](bp *Binpacker, create func(byte) (T, error)) (*T, error) {
	tag, tagErr := bp.UnmarshallByte()
	if tagErr != nil {
		bp.Err = tagErr
		return nil, tagErr
	}

	result, err := create(tag)
	if err != nil {
		bp.Err = err
		return nil, err
	}

	result.Unmarshall(bp)

	return &result, nil
}
