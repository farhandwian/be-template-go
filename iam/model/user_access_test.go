package model

import (
	"reflect"
	"testing"
)

func TestUserAccess_Extract(t *testing.T) {
	tests := []struct {
		name string
		ua   UserAccess
		want []Access
	}{
		{"Single_Access", "8", []Access{"8"}},
		{"Two_Accesses", "a", []Access{"2", "8"}},
		{"Three_Accesses", "d", []Access{"1", "4", "8"}},
		{"Complex_Case", "12f", []Access{"1", "2", "4", "8", "20", "100"}},
		{"Zero_Case", "0", []Access{}},
		{"Large_Number", "10000000000000000", []Access{"10000000000000000"}},
		{"Invalid_Input", "g", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ua.Extract()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserAccess.Extract() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToUserAccess(t *testing.T) {
	tests := []struct {
		name     string
		accesses []string
		want     UserAccess
	}{
		{"Single_Access", []string{"8"}, "8"},
		{"Two_Accesses", []string{"2", "8"}, "a"},
		{"Three_Accesses", []string{"1", "4", "8"}, "d"},
		{"Complex_Case", []string{"1", "2", "4", "8", "20", "100"}, "12f"},
		{"Empty_Case", []string{}, "0"},
		{"Invalid_Input", []string{"g"}, "0"},
		{"Mixed_Valid_Invalid", []string{"8", "g", "2"}, "a"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToUserAccess(tt.accesses...); got != tt.want {
				t.Errorf("ToUserAccess() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserAccessRoundTrip(t *testing.T) {
	tests := []struct {
		name  string
		input UserAccess
	}{
		{"Simple", "a"},
		{"Complex", "12f"},
		{"Large", "10000000000000000"},
		{"Zero", "0"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			extracted := tt.input.Extract()
			reconstructed := ToUserAccess(extractedToStrings(extracted)...)
			if reconstructed != tt.input {
				t.Errorf("Round trip failed. Input: %v, Reconstructed: %v", tt.input, reconstructed)
			}
		})
	}
}

// Helper function to convert []Access to []string
func extractedToStrings(accesses []Access) []string {
	result := make([]string, len(accesses))
	for i, access := range accesses {
		result[i] = string(access)
	}
	return result
}

func TestAssignAccess(t *testing.T) {
	tests := []struct {
		name           string
		currentAccess  UserAccess
		newAccesses    []Access
		expectedAccess UserAccess
		expectError    bool
	}{
		{
			name:           "Assign single access 1",
			currentAccess:  "8",
			newAccesses:    []Access{"1", "2"},
			expectedAccess: "b",
			expectError:    false,
		},
		{
			name:           "Assign single access 2",
			currentAccess:  "a",
			newAccesses:    []Access{"1", "2"},
			expectedAccess: "b",
			expectError:    false,
		},
		{
			name:           "Assign single access 3",
			currentAccess:  "8",
			newAccesses:    []Access{"1"},
			expectedAccess: "9",
			expectError:    false,
		},
		{
			name:           "Assign multiple accesses",
			currentAccess:  "8",
			newAccesses:    []Access{"1", "2"},
			expectedAccess: "b",
			expectError:    false,
		},
		{
			name:           "Assign existing access",
			currentAccess:  "a",
			newAccesses:    []Access{"8"},
			expectedAccess: "a",
			expectError:    false,
		},
		{
			name:           "Assign to zero",
			currentAccess:  "0",
			newAccesses:    []Access{"1", "2", "4"},
			expectedAccess: "7",
			expectError:    false,
		},
		{
			name:           "Assign zero",
			currentAccess:  "f",
			newAccesses:    []Access{"0"},
			expectedAccess: "f",
			expectError:    false,
		},
		{
			name:           "Invalid current access",
			currentAccess:  "g",
			newAccesses:    []Access{"1"},
			expectedAccess: "g",
			expectError:    true,
		},
		{
			name:           "Invalid new access",
			currentAccess:  "8",
			newAccesses:    []Access{"g"},
			expectedAccess: "8",
			expectError:    true,
		},
		{
			name:          "Assign access sample",
			currentAccess: "1", // 0000 0000 0000 0000 0000 0000 0000 0001
			newAccesses: []Access{
				"20000000", // 0010 0000 0000 0000 0000 0000 0000 0000
				"800000",   // 0000 0000 1000 0000 0000 0000 0000 0000
			},
			expectedAccess: "20800001", // 0010 0000 1000 0000 0000 0000 0000 0001
			expectError:    false,
		},
		{
			name:          "Experiment",
			currentAccess: "0",
			newAccesses: []Access{
				"20000000000000",
				"40000000000000",
				"80000000000000",
				"100000000000000",
				"200000000000000",
				"400000000000000",
				"800000000000000",
			},
			expectedAccess: "8",
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentAccess := tt.currentAccess
			err := currentAccess.AssignAccess(tt.newAccesses...)

			if (err != nil) != tt.expectError {
				t.Errorf("AssignAccess() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError && currentAccess != tt.expectedAccess {
				t.Errorf("AssignAccess() got = %v, want %v", currentAccess, tt.expectedAccess)
			}
		})
	}
}

func TestRevokeAccess(t *testing.T) {
	tests := []struct {
		name            string
		currentAccess   UserAccess
		revokedAccesses []Access
		expectedAccess  UserAccess
		expectError     bool
	}{
		{
			name:            "Revoke single access",
			currentAccess:   "a",           // 1010
			revokedAccesses: []Access{"2"}, // 10
			expectedAccess:  "8",           // 1000
			expectError:     false,
		},
		{
			name:            "Revoke multiple accesses",
			currentAccess:   "a",                // 1010
			revokedAccesses: []Access{"1", "2"}, // 1, 10
			expectedAccess:  "8",                // 1000
			expectError:     false,
		},
		{
			name:            "Revoke non-existent access",
			currentAccess:   "a",           // 1010
			revokedAccesses: []Access{"4"}, // 100
			expectedAccess:  "a",           // 1010
			expectError:     false,
		},
		{
			name:            "Revoke from zero",
			currentAccess:   "0",
			revokedAccesses: []Access{"1", "2", "4"},
			expectedAccess:  "0",
			expectError:     false,
		},
		{
			name:            "Revoke all accesses",
			currentAccess:   "f",
			revokedAccesses: []Access{"1", "2", "4", "8"},
			expectedAccess:  "0",
			expectError:     false,
		},
		{
			name:            "Invalid current access",
			currentAccess:   "g",
			revokedAccesses: []Access{"1"},
			expectedAccess:  "g",
			expectError:     true,
		},
		{
			name:            "Invalid revoked access",
			currentAccess:   "8",
			revokedAccesses: []Access{"g"},
			expectedAccess:  "8",
			expectError:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			currentAccess := tt.currentAccess
			err := currentAccess.RevokeAccess(tt.revokedAccesses...)

			if (err != nil) != tt.expectError {
				t.Errorf("RevokeAccess() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError && currentAccess != tt.expectedAccess {
				t.Errorf("RevokeAccess() got = %v, want %v", currentAccess, tt.expectedAccess)
			}
		})
	}
}

func TestHasAccess(t *testing.T) {
	tests := []struct {
		name           string
		currentAccess  UserAccess
		accessToCheck  Access
		expectedResult bool
	}{
		{
			name:           "Has default access",
			currentAccess:  "f",
			accessToCheck:  "1",
			expectedResult: true,
		},
		{
			name:           "Has single access",
			currentAccess:  "a",
			accessToCheck:  "2",
			expectedResult: true,
		},
		{
			name:           "Does not have access",
			currentAccess:  "a",
			accessToCheck:  "4",
			expectedResult: false,
		},
		{
			name:           "Has access in multiple accesses",
			currentAccess:  "f",
			accessToCheck:  "4",
			expectedResult: true,
		},
		{
			name:           "Zero current access",
			currentAccess:  "0",
			accessToCheck:  "1",
			expectedResult: false,
		},
		{
			name:           "Zero access to check",
			currentAccess:  "f",
			accessToCheck:  "0",
			expectedResult: true,
		},
		{
			name:           "Invalid current access",
			currentAccess:  "g",
			accessToCheck:  "1",
			expectedResult: false,
		},
		{
			name:           "Invalid access to check",
			currentAccess:  "f",
			accessToCheck:  "g",
			expectedResult: false,
		},
		{
			name:           "Large numbers",
			currentAccess:  "10000000000000000",
			accessToCheck:  "10000000000000000",
			expectedResult: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.currentAccess.HasAccess(tt.accessToCheck)
			if result != tt.expectedResult {
				t.Errorf("HasAccess() got = %v, want %v", result, tt.expectedResult)
			}
		})
	}
}

func TestUserAccess_ResetAccess(t *testing.T) {
	tests := []struct {
		name           string
		initialAccess  UserAccess
		newAccesses    []Access
		expectedAccess UserAccess
		expectError    bool
	}{
		{
			name:           "Reset to single access",
			initialAccess:  "f",           // 1111
			newAccesses:    []Access{"2"}, // 10
			expectedAccess: "2",           // 10
			expectError:    false,
		},
		{
			name:           "Reset to multiple accesses",
			initialAccess:  "f",                // 1111
			newAccesses:    []Access{"2", "8"}, // 10, 1000
			expectedAccess: "a",                // 1010
			expectError:    false,
		},
		{
			name:           "Reset to no access",
			initialAccess:  "f", // 1111
			newAccesses:    []Access{},
			expectedAccess: "0", // 0
			expectError:    false,
		},
		{
			name:           "Reset with existing and new accesses",
			initialAccess:  "f",                      // 1111
			newAccesses:    []Access{"1", "4", "20"}, // 1, 100, 100000
			expectedAccess: "25",                     // 100101
			expectError:    false,
		},
		{
			name:           "Reset with invalid access",
			initialAccess:  "f",
			newAccesses:    []Access{"g"},
			expectedAccess: "f", // Should not change on error
			expectError:    true,
		},
		{
			name:           "Reset with large numbers",
			initialAccess:  "10000000000000000",                                // 2^64
			newAccesses:    []Access{"10000000000000000", "20000000000000000"}, // 2^64, 2^65
			expectedAccess: "30000000000000000",                                // 2^64 + 2^65
			expectError:    false,
		},
		{
			name:           "Assign and revoke simultaneously",
			initialAccess:  "2f",                           // 101111
			newAccesses:    []Access{"1", "8", "40", "80"}, // 1, 1000, 1000000, 10000000
			expectedAccess: "c9",                           // 11001001
			expectError:    false,
		},
		{
			name:           "Add new accesses while keeping existing ones",
			initialAccess:  "1f",                                                 // 11111 (a, b, c, d, e)
			newAccesses:    []Access{"1", "2", "4", "8", "10", "20", "40", "80"}, // a, b, c, d, e, x, y, z
			expectedAccess: "ff",                                                 // 11111111
			expectError:    false,
		},
		{
			name:           "Revoke specific accesses",
			initialAccess:  "1f",                     // 11111 (a, b, c, d, e)
			newAccesses:    []Access{"2", "8", "10"}, // b, d, e
			expectedAccess: "1a",                     // 11010
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ua := tt.initialAccess
			err := ua.ResetAccess(tt.newAccesses...)

			if (err != nil) != tt.expectError {
				t.Errorf("ResetAccess() error = %v, expectError %v", err, tt.expectError)
				return
			}

			if !tt.expectError && ua != tt.expectedAccess {
				t.Errorf("ResetAccess() got = %v, want %v", ua, tt.expectedAccess)
			}
		})
	}
}
