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
        "transform one value",
        letterPerScore { 1: { "A" } },
        scorePerLetter { "a": 1 },
    }, {
        "transform more values",
        letterPerScore { 1: { "A", "E", "I", "O", "U" } },
        scorePerLetter { "a": 1, "e": 1, "i": 1, "o": 1, "u": 1 },
    }, {
        "transform more keys",
        letterPerScore {
            1: { "A", "E" },
            2: { "D", "G" },
        },
        scorePerLetter {
            "a": 1,
            "e": 1,
            "d": 2,
            "g": 2,
        },
    }, {
        "transform full dataset",
        letterPerScore {
            1:    { "A", "E", "I", "O", "U", "L", "N", "R", "S", "T" },
            2:    { "D", "G" },
            3:    { "B", "C", "M", "P" },
            4:    { "F", "H", "V", "W", "Y" },
            5:    { "K" },
            8:    { "J", "X" },
            10: { "Q", "Z" },
        },
        scorePerLetter {
            "a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1,
            "d": 2, "g": 2,
            "b": 3, "c": 3, "m": 3, "p": 3,
            "f": 4, "h": 4, "v": 4, "w": 4, "y": 4,
            "k": 5,
            "j": 8, "x": 8,
            "q": 10, "z": 10,
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
