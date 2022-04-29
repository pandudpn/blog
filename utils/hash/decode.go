package hash

func (h *hash) Decode(v string) (int, error) {
	if h == nil {
		return 0, ErrData
	}
	
	d, err := h.h.DecodeWithError(v)
	if err != nil {
		return 0, err
	}
	
	return d[0], nil
}
