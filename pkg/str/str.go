package str

import (
	"encoding/json"
	"strconv"
)

// Strings2Int64s converts a slice of strings to a slice of int64s.
func Strings2Int64s(strs []string) ([]int64, error) {
	var int64s []int64
	for _, str := range strs {
		i, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		int64s = append(int64s, i)
	}
	return int64s, nil
}

// Int64s2String converts a slice of int64s to a slice of string.
func Int64s2String(int64s []int64) (string, error) {
	b, err := json.Marshal(int64s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
