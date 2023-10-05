package helpers

import (
	"strconv"
	"strings"
)

func IntSliceToString(numbers []int64) string {
	var result []string
	for _, id := range numbers {
		result = append(result, strconv.FormatInt(id, 10))
	}
	return strings.Join(result, ",")
}
