package etl

import "testing"

type letterPerScore map[int][]string
type scorePerLetter map[string]int

var tests = []struct {
    description string
    input       letterPerScore
    expected    scorePerLetter
} {
    {
        "should transform one value",
        letterPerScore { 1: { "WORLD" } },
        scorePerLetter { "world": 1 },
    }, {
        "should transform more values",
        letterPerScore {
            1: { "WORLD", "GSCHOOLERS" },
        },
        scorePerLetter {
            "world": 1,
            "gschoolers": 1,
        },
    }, {
        "should transform more keys",
        letterPerScore {
            1: { "APPLE", "ARTICHOKE" },
            2: { "BOAT", "BALLERINA" },
        },
        scorePerLetter {
            "apple": 1,
            "artichoke": 1,
            "boat": 2,
            "ballerina": 2,
        },
    }, {
        "should transform full dataset",
        letterPerScore {
            1:  { "A", "E", "I", "O", "U", "L", "N", "R", "S", "T" },
            2:  { "D", "G" },
            3:  { "B", "C", "M", "P" },
            4:  { "F", "H", "V", "W", "Y" },
            5:  { "K" },
            8:  { "J", "X" },
            10: { "Q", "Z" },
        },
        scorePerLetter {
            "a": 1, "b":  3, "c": 3, "d": 2, "e": 1,
            "f": 4, "g":  2, "h": 4, "i": 1, "j": 8,
            "k": 5, "l":  1, "m": 3, "n": 1, "o": 1,
            "p": 3, "q": 10, "r": 1, "s": 1, "t": 1,
            "u": 1, "v":  4, "w": 4, "x": 8, "y": 4,
            "z": 10,
        },
    },
}

func equal(actual map[string]int, expected map[string]int) bool {
    if len(actual) != len(expected) {
        return false
    }

    for key, actualValue := range actual {
        expectedValue, present := expected[key]
        if !present || actualValue != expectedValue {
            return false
        }
    }

    return true
}

func TestETL(t *testing.T) {
    for _, test := range tests {
        var actual = Transform(test.input)
        if equal(actual, test.expected) {
            t.Logf("Passed: %s", test.description)
        } else {
            t.Fatalf("Failed: %s, expected %v but it was %v", test.description, test.expected, actual)
        }
    }
}

func BenchmarkETL(b *testing.B) {
    for _, test := range tests {
        for i := 0; i < b.N; i++ {
            Transform(test.input)
        }
    }
}
