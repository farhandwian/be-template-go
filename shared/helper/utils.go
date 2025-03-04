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

func ValidateSortParamsWithForeignKey(allowedSortBy map[string]bool, allowedForeignSortBy map[string]string, sortBy, sortOrder string, defaultSortBy string) (string, string, error) {
	// Default values
	if sortBy == "" {
		sortBy = defaultSortBy
	}
	if sortOrder == "" {
		sortOrder = "desc"
	}

	sortOrder = strings.ToLower(sortOrder)
	if sortOrder != "asc" && sortOrder != "desc" {
		return "", "", fmt.Errorf("invalid sort_order parameter: %s", sortOrder)
	}

	// Check if the sort field is from the main table
	if allowedSortBy[sortBy] {
		return sortBy, sortOrder, nil
	}

	// Check if the sort field is from a foreign key table
	if columnName, exists := allowedForeignSortBy[sortBy]; exists {
		return columnName, sortOrder, nil
	}

	return "", "", fmt.Errorf("invalid sort_by parameter: %s", sortBy)
}
