package bob

import "regexp"

var fine = regexp.MustCompile("^[\\s\\v\\x{00a0}\\x{2002}]*$")
var whoa = regexp.MustCompile("^[^\\p{Ll}]*[A-Z][^\\p{Ll}]*$")
var sure = regexp.MustCompile("^.*\\?[\\s\\v]*$")

func Hey(question string) string {
    if fine.MatchString(question) {
        return "Fine. Be that way!"
    }
    if whoa.MatchString(question) {
        return "Whoa, chill out!"
    }
    if sure.MatchString(question) {
        return "Sure."
    }
    return "Whatever."
}
