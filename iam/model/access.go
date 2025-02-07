package model

import (
	"errors"
	"strings"
)

type Access string // access function in hex

func NewAccess(access string) Access {
	return Access(access)
}

func (a Access) Validate() error {
	// Convert to lowercase for easier comparison
	s := strings.ToLower(string(a))

	// Check if the string is empty
	if s == "" {
		return errors.New("access value cannot be empty")
	}

	// Check if the first character is valid
	if s[0] != '1' && s[0] != '2' && s[0] != '4' && s[0] != '8' {
		return errors.New("access value must start with 1, 2, 4, or 8")
	}

	// Check if all remaining characters are '0'
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			return errors.New("access value must only contain zeros after the first digit")
		}
	}

	return nil
}

// 1 -> 1
//
// 2 -> 2
//
// 3 -> 4
//
// 4 -> 8
//
// 5 -> 10
//
// 6 -> 20
//
// 7 -> 40
//
// 8 -> 80
//
// 9 -> 100
//
// 10 -> 200
//
// 11 -> 400
//
// 12 -> 800
//
// and continued..

// func GetAccess(sequenceIndex int) (Access, error) {
// 	if sequenceIndex <= 0 {
// 		return Access(""), errors.New("sequenced index must be a positive integer")
// 	}

// 	// Initialize with "1"
// 	value := "1"

// 	// Determine the first digit
// 	switch (sequenceIndex - 1) % 4 {
// 	case 0:
// 		value = "1"
// 	case 1:
// 		value = "2"
// 	case 2:
// 		value = "4"
// 	case 3:
// 		value = "8"
// 	}

// 	// Add the appropriate number of zeros
// 	zeros := (sequenceIndex - 1) / 4
// 	for i := 0; i < zeros; i++ {
// 		value += "0"
// 	}

// 	return Access(value), nil
// }

// 1 -> 1
//
// 2 -> 2
//
// 4 -> 3
//
// 8 -> 4
//
// 10 -> 5
//
// 20 -> 6
//
// 40 -> 7
//
// 80 -> 8
//
// 100 -> 9
//
// 200 -> 10
//
// 400 -> 11
//
// 800 -> 12
//
// and continued..
// func GetSequencedIndex(access Access) (int, error) {
// 	// Validate the access value using the existing Validate method
// 	if err := access.Validate(); err != nil {
// 		return 0, err
// 	}

// 	// Convert to uppercase for consistency
// 	s := strings.ToUpper(string(access))

// 	// Determine the base index from the first character
// 	var baseIndex int
// 	switch s[0] {
// 	case '1':
// 		baseIndex = 1
// 	case '2':
// 		baseIndex = 2
// 	case '4':
// 		baseIndex = 3
// 	case '8':
// 		baseIndex = 4
// 	}

// 	// Calculate the sequenced index
// 	zeroCount := len(s) - 1
// 	sequencedIndex := baseIndex + (zeroCount * 4)

// 	return sequencedIndex, nil
// }

// func (a Access) Validate() error {

// 	// Convert the hexadecimal string to a big.Int
// 	num, ok := new(big.Int).SetString(string(a), 16)
// 	if !ok {
// 		return errors.New("invalid hexadecimal format")
// 	}

// 	// Check if the number is positive
// 	if num.Sign() <= 0 {
// 		return errors.New("access value must be positive")
// 	}

// 	// Check if it's a power of two
// 	// A number is a power of two if and only if it has exactly one bit set to 1
// 	// We can check this by ANDing the number with (number - 1)
// 	// If the result is 0, it's a power of two

// 	// Subtract 1 from the number
// 	one := big.NewInt(1)
// 	numMinusOne := new(big.Int).Sub(num, one)

// 	// Perform bitwise AND
// 	result := new(big.Int).And(num, numMinusOne)

// 	if result.Cmp(big.NewInt(0)) != 0 {
// 		return errors.New("access value must be a power of two")
// 	}

// 	return nil
// }

// Helper function to test the Validate method
func TestAccess(access string) error {
	return Access(strings.ToLower(access)).Validate()
}

// "1": "function_1"
// "2": "function_2"
// "4": "function_3"
// "8": "function_4"

// "10": "function_5"
// "20": "function_6"
// "40": "function_7"
// "80": "function_8"

// "100": "function_9"
// "200": "function_10"
// "400": "function_11"
// "800": "function_12"

// "1000": "function_13"
// "2000": "function_14"
// "4000": "function_15"
// "8000": "function_16"

// "10000": "function_17"
// "20000": "function_18"
// "40000": "function_19"
// "80000": "function_20"

// "100000": "function_21"
// "200000": "function_22"
// "400000": "function_23"
// "800000": "function_24"

// "1000000": "function_25"
// "2000000": "function_26"
// "4000000": "function_27"
// "8000000": "function_28"

// "10000000": "function_29"
// "20000000": "function_30"
// "40000000": "function_31"
// "80000000": "function_32"

// "100000000": "function_33"
// "200000000": "function_34"
// "400000000": "function_35"
// "800000000": "function_36"

// "1000000000": "function_37"
// "2000000000": "function_38"
// "4000000000": "function_39"
// "8000000000": "function_40

// func GetAccesses(accessesIndex []int) ([]Access, error) {
// 	var accesses []Access
// 	for _, v := range accessesIndex {
// 		a, err := GetAccess(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		accesses = append(accesses, a)
// 	}
// 	return accesses, nil
// }

// type AccessMapID string
// type AccessMap struct {
// 	ID          AccessMapID
// 	Access      Access
// 	Description string
// }

// var accessMaps []AccessMap = []AccessMap{
// 	{
// 		ID:          "a1",
// 		Access:      Access("10"),
// 		Description: "Function one",
// 	},
// 	{
// 		ID:          "a2",
// 		Access:      Access("10"),
// 		Description: "Function one",
// 	},
// 	{
// 		ID:          "a3",
// 		Access:      Access("10"),
// 		Description: "Function one",
// 	},
// }

// func GetAccessesByMap(accessIDs []AccessMapID) ([]Access, error) {
// 	// use accessMap to map between
// 	// return error if the id is not exist
// }
