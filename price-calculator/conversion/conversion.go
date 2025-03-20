package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(strings []string) ([]float64, error) {
	result := make([]float64, len(strings))

	for i, string := range strings {
		float, err := strconv.ParseFloat(string, 64)

		if err != nil {
			return nil, errors.New("Error converting string to float")
		}

		result[i] = float
	}

	return result, nil
}
