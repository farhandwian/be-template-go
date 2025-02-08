package gateway

func ValidatePageSize(page, size int) (int, int) {

	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 10
	}

	return page, size
}
