package model

import (
	"testing"
)

func TestAccess_Validate(t *testing.T) {
	tests := []struct {
		name    string
		access  string
		wantErr bool
	}{
		{"Valid_1", "1", false},
		{"Valid_2", "2", false},
		{"Valid_4", "4", false},
		{"Valid_8", "8", false},
		{"Valid_10", "10", false},
		{"Valid_20", "20", false},
		{"Valid_40", "40", false},
		{"Valid_80", "80", false},
		{"Valid_100", "100", false},
		{"Valid_200", "200", false},
		{"Valid_1000", "1000", false},
		{"Valid_LargeNumber", "10000000000000000", false}, // 2^64
		{"Invalid_Empty", "", true},
		{"Invalid_0", "0", true},
		{"Invalid_3", "3", true},
		{"Invalid_5", "5", true},
		{"Invalid_A", "A", true},
		{"Invalid_F", "F", true},
		{"Invalid_11", "11", true},
		{"Invalid_1A", "1A", true},
		{"Invalid_Negative", "-4", true},
		{"Invalid_NonHex", "G", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Access(tt.access)
			err := a.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Access.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTestAccess(t *testing.T) {
	tests := []struct {
		name    string
		access  string
		wantErr bool
	}{
		{"LowerCase", "8", false},
		{"UpperCase", "80", false},
		{"MixedCase", "100A", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := TestAccess(tt.access)
			if (err != nil) != tt.wantErr {
				t.Errorf("TestAccess() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// func TestGetAccess(t *testing.T) {
// 	tests := []struct {
// 		name           string
// 		sequencedIndex int
// 		wantAccess     Access
// 		wantErr        bool
// 	}{
// 		{"Valid_1", 1, "1", false},
// 		{"Valid_2", 2, "2", false},
// 		{"Valid_3", 3, "4", false},
// 		{"Valid_4", 4, "8", false},
// 		{"Valid_5", 5, "10", false},
// 		{"Valid_6", 6, "20", false},
// 		{"Valid_7", 7, "40", false},
// 		{"Valid_8", 8, "80", false},
// 		{"Valid_9", 9, "100", false},
// 		{"Valid_10", 10, "200", false},
// 		{"Valid_16", 16, "8000", false},
// 		{"Valid_20", 20, "80000", false},
// 		{"Valid_32", 32, "80000000", false},
// 		{"Invalid_0", 0, "", true},
// 		{"Invalid_Negative", -1, "", true},
// 		{"Valid_LargeNumber", 100, "8000000000000000000000000", false},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotAccess, err := GetAccess(tt.sequencedIndex)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetAccess() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotAccess != tt.wantAccess {
// 				t.Errorf("GetAccess() = %v, want %v", gotAccess, tt.wantAccess)
// 			}
// 		})
// 	}
// }

// func TestGetSequencedIndex(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		access    Access
// 		wantIndex int
// 		wantErr   bool
// 	}{
// 		{"Valid_1", "1", 1, false},
// 		{"Valid_2", "2", 2, false},
// 		{"Valid_4", "4", 3, false},
// 		{"Valid_8", "8", 4, false},
// 		{"Valid_10", "10", 5, false},
// 		{"Valid_20", "20", 6, false},
// 		{"Valid_40", "40", 7, false},
// 		{"Valid_80", "80", 8, false},
// 		{"Valid_100", "100", 9, false},
// 		{"Valid_200", "200", 10, false},
// 		{"Valid_400", "400", 11, false},
// 		{"Valid_800", "800", 12, false},
// 		{"Valid_1000", "1000", 13, false},
// 		{"Valid_8000", "8000", 16, false},
// 		{"Valid_10000", "10000", 17, false},
// 		{"Valid_LargeNumber", "1000000000000000", 61, false},
// 		{"Valid_SecondLargeNumber", "8000000000000000000000000", 100, false},
// 		{"Invalid_Empty", "", 0, true},
// 		{"Invalid_3", "3", 0, true},
// 		{"Invalid_5", "5", 0, true},
// 		{"Invalid_11", "11", 0, true},
// 		{"Invalid_1A", "1A", 0, true},
// 		{"Invalid_101", "101", 0, true},
// 		{"Invalid_Negative", "-4", 0, true},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotIndex, err := GetSequencedIndex(tt.access)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetSequencedIndex() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if gotIndex != tt.wantIndex {
// 				t.Errorf("GetSequencedIndex() = %v, want %v", gotIndex, tt.wantIndex)
// 			}
// 		})
// 	}
// }
