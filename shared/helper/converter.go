package helper

import (
	"encoding/json"

	"gorm.io/datatypes"
)

func ToDataTypeJSON[T any](a ...T) datatypes.JSON {
	if len(a) == 0 {
		return datatypes.JSON{}
	}

	jsonData, err := json.Marshal(a)
	if err != nil {
		// Handle error, perhaps log it
		return datatypes.JSON{}
	}

	return datatypes.JSON(jsonData)
}

func ToDataTypeJSONPtr[T any](a ...T) *datatypes.JSON {
	if len(a) == 0 {
		return nil
	}

	jsonData, err := json.Marshal(a)
	if err != nil {

		return nil
	}

	jsonValue := datatypes.JSON(jsonData)
	return &jsonValue
}
