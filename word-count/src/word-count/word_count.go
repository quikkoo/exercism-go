package wordcount

import (
    "regexp"
    "strings"
)

var pattern = regexp.MustCompile("[^'a-zA-Z0-9]+")

func Count(phrase string) map[string]int {
    var words = pattern.Split(phrase, -1)

    var result = make(map[string]int)
    for _, wrd := range words {
        if wrd != "" {
            var word = strings.ToLower(wrd)
            var count = result[word] + 1
            result[word] = count
        }
    }

    return result
}
