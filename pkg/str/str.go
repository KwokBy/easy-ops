package str

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
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
	return string(b), err
}

func String2Int64s(str string) ([]int64, error) {
	var int64s []int64
	err := json.Unmarshal([]byte(str), &int64s)
	return int64s, err
}

// VersionIncrease increases the version of the given string.
func VersionIncrease(str string) string {
	if str == "" {
		return "1.0.0"
	}
	var (
		major, minor, patch int
		err                 error
	)
	if major, minor, patch, err = ParseVersionStringWithSeparator(str, "."); err != nil {
		return "1.0.0"
	}
	return strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch+1)
}

// ParseVersionStringWithSeparator parses the given string to major, minor and patch.
func ParseVersionStringWithSeparator(str, separator string) (major, minor, patch int, err error) {
	if str == "" {
		return 0, 0, 0, nil
	}
	strs := strings.Split(str, separator)
	if len(strs) != 3 {
		return 0, 0, 0, errors.New("invalid version string")
	}
	if major, err = strconv.Atoi(strs[0]); err != nil {
		return 0, 0, 0, err
	}
	if minor, err = strconv.Atoi(strs[1]); err != nil {
		return 0, 0, 0, err
	}
	if patch, err = strconv.Atoi(strs[2]); err != nil {
		return 0, 0, 0, err
	}
	return major, minor, patch, nil
}
