package helper

import (
	"fmt"
	"strings"
)

func ValidatePageSize(page, size int) (int, int) {

	if page <= 0 {
		page = 1
	}

	if size <= 0 {
		size = 10
	}

	return page, size
}

// ValidateSortParams memvalidasi parameter sorting
func ValidateSortParams(allowedSortBy map[string]bool, sortBy, sortOrder string, defaultSortBy string) (string, string, error) {
	// Default values
	if sortBy == "" {
		sortBy = defaultSortBy
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	if !allowedSortBy[sortBy] {
		return "", "", fmt.Errorf("invalid sort_by parameter: %s", sortBy)
	}

	// Validate sortOrder
	sortOrder = strings.ToLower(sortOrder)
	if sortOrder != "asc" && sortOrder != "desc" {
		return "", "", fmt.Errorf("invalid sort_order parameter: %s", sortOrder)
	}

	return sortBy, sortOrder, nil
}
