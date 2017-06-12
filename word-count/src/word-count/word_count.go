package wordcount

import (
    "regexp"
    "strings"
)

var pattern = regexp.MustCompile("[']?[^'a-zA-Z0-9]+[']?")

func Calculate(phrase string) map[string]int {
    var words = pattern.Split(strings.ToLower(phrase), -1)

    var result = make(map[string]int)
    for _, word := range words {
        if word != "" {
            var count = result[word] + 1
            result[word] = count
        }
    }

    return result
}
