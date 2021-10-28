package fsfire

func InSlice(source []interface{}, find interface{}) bool {
	for _, item := range source {
		if item == find {
			return true
		}
	}

	return false
}
