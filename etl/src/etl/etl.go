package etl

import "strings"

func Transform(data map[int][]string) map[string]int {
    var result = make(map[string]int)
    for key, values := range data {
        for _, value := range values {
            result[strings.ToLower(value)] = key
        }
    }

    return result
}
