package anagram

import (
    "sort"
    "strings"
)

type RuneArray []rune

func (runes RuneArray) Len() int {
    return len(runes)
}

func (runes RuneArray) Swap(lhs int, rhs int) {
    runes[lhs], runes[rhs] = runes[rhs], runes[lhs]
}

func (runes RuneArray) Less(lhs int, rhs int) bool {
    return runes[lhs] < runes[rhs]
}

func filter(array []string, f func(string) bool) []string {
    var filtered = make([]string, 0)
    for _, value := range array {
        if f(value) {
            filtered = append(filtered, value)
        }
    }

    return filtered
}

func isSame(lhs string, rhs string) bool {
    return strings.Compare(lhs, rhs) == 0
}

func hasSameSize(lhs string, rhs string) bool {
    return len(lhs) == len(rhs)
}

func hasSameElements(lhs RuneArray, rhs RuneArray) bool {
    var result = true
    for i := 0; result && i < len(lhs) && i < len(rhs); i++ {
        result = lhs[i] == rhs[i]
    }
    return result
}

func Matches(subject string, candidates []string) []string {
    var lsub = strings.ToLower(subject)
    var ssub = RuneArray(lsub)
    sort.Sort(ssub)

    return filter(candidates, func(candidate string) bool {
        var lcan = strings.ToLower(candidate)
        var scan = RuneArray(lcan)
        sort.Sort(scan)

        return hasSameSize(lsub, lcan) && !isSame(lsub, lcan) && hasSameElements(ssub, scan)
    })
}
