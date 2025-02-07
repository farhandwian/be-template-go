package model

import (
	"errors"
	"math/big"
)

type UserAccess string

func NewUserAccess() UserAccess {
	return UserAccess(DEFAULT_OPERATION)
}

// func NewUserAccessAdmin() UserAccess {
// 	userAccess := NewUserAccess()
// 	userAccess.AssignAccess(ADMIN_OPERATION)
// 	return userAccess
// }

func (ua UserAccess) Extract() []Access {
	// Convert UserAccess string to integer
	num, ok := new(big.Int).SetString(string(ua), 16)
	if !ok {
		return nil
	}

	// If the number is zero, return an empty slice instead of nil
	if num.Cmp(big.NewInt(0)) == 0 {
		return []Access{}
	}

	var result []Access

	// Iterate through each bit
	for i := 0; num.BitLen() > 0; i++ {
		// Check if the least significant bit is set
		if num.Bit(0) == 1 {
			// Calculate the corresponding power of 2
			access := new(big.Int).Lsh(big.NewInt(1), uint(i))
			// Convert to hexadecimal string
			hexStr := access.Text(16)
			result = append(result, Access(hexStr))
		}
		// Right shift by 1 bit
		num.Rsh(num, 1)
	}

	return result
}

func (ua UserAccess) HasAccess(access Access) bool {

	// Convert current access to big.Int
	current, ok := new(big.Int).SetString(string(ua), 16)
	if !ok {
		return false
	}

	// Convert access to check to big.Int
	checkAccess, ok := new(big.Int).SetString(string(access), 16)
	if !ok {
		return false
	}

	// Perform bitwise AND operation
	result := new(big.Int).And(current, checkAccess)

	// If the result equals the check access, it means the access is present
	return result.Cmp(checkAccess) == 0
}

func (ua *UserAccess) AssignAccess(newAccesses ...Access) error {
	// Convert current access to big.Int
	current, ok := new(big.Int).SetString(string(*ua), 16)
	if !ok {
		return ErrInvalidCurrentAccess
	}

	// Process each new access
	for _, access := range newAccesses {
		newAccess, ok := new(big.Int).SetString(string(access), 16)
		if !ok {
			return ErrInvalidNewAccess
		}

		// Perform binary OR operation
		current.Or(current, newAccess)
	}

	// Convert result back to UserAccess
	*ua = UserAccess(current.Text(16))
	return nil
}

func (ua *UserAccess) RevokeAccess(revokedAccesses ...Access) error {
	// Convert current access to big.Int
	current, ok := new(big.Int).SetString(string(*ua), 16)
	if !ok {
		return ErrInvalidCurrentAccess
	}

	// Create a big.Int to hold all revoked accesses
	revokedTotal := new(big.Int)

	// Process each revoked access
	for _, access := range revokedAccesses {
		revokedAccess, ok := new(big.Int).SetString(string(access), 16)
		if !ok {
			return ErrInvalidRevokedAccess
		}

		// Add this access to the total of revoked accesses
		revokedTotal.Or(revokedTotal, revokedAccess)
	}

	// Perform binary AND NOT operation to remove revoked accesses
	current.AndNot(current, revokedTotal)

	// Convert result back to UserAccess
	*ua = UserAccess(current.Text(16))
	return nil
}

func (ua *UserAccess) ResetAccess(newAccesses ...Access) error {
	// Create a new big.Int to hold the new access rights
	newAccess := new(big.Int)

	// Process each new access
	for _, access := range newAccesses {
		accessValue, ok := new(big.Int).SetString(string(access), 16)
		if !ok {
			return errors.New("invalid access value")
		}

		// Perform binary OR operation to add this access
		newAccess.Or(newAccess, accessValue)
	}

	// Convert the result back to UserAccess
	*ua = UserAccess(newAccess.Text(16))

	return nil
}

// Helper function to convert a slice of strings to UserAccess
func ToUserAccess(accesses ...string) UserAccess {
	total := big.NewInt(0)
	for _, access := range accesses {
		num, ok := new(big.Int).SetString(access, 16)
		if !ok {
			continue
		}
		total.Or(total, num)
	}
	return UserAccess(total.Text(16))
}

// Error types for better error handling
var (
	ErrInvalidCurrentAccess = errors.New("invalid current access")
	ErrInvalidNewAccess     = errors.New("invalid new access")
	ErrInvalidRevokedAccess = errors.New("invalid revoked access")
)
