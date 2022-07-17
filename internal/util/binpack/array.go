package binpack

func MarshallArray16[T BinaryPackable](bp *Binpacker, items []T) {
	l := len(items)
	bp.MarshallUInt16(uint16(l))

	for _, item := range items {
		item.Marshall(bp)
	}
}
func UnmarshallArray16[T BinaryPackable](bp *Binpacker) ([]T, error) {
	n, err := bp.UnmarshallUInt16()
	if err != nil {
		bp.Err = err
		return nil, err
	}

	l := int(n)
	result := make([]T, l)

	for i := 0; i < l; i++ {
		ierr := result[i].Unmarshall(bp)
		if ierr != nil {
			bp.Err = ierr
			return nil, ierr
		}
	}

	return result, nil
}
