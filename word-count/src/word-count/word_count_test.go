package wordcount

import "testing"

var tests = []struct {
    description string
    input       string
    expected    map[string]int
} {
    {
        "count one word",
        "word",
        map[string]int { "word": 1 },
    }, {
        "count one of each",
        "one of each",
        map[string]int {
            "each": 1,
            "of": 1,
            "one": 1 },
    }, {
        "count multiple occurrences",
        "one fish two fish red fish blue fish",
        map[string]int {
            "blue": 1,
            "fish": 4,
            "one": 1,
            "red": 1,
            "two": 1 },
    }, {
        "count everything just once",
        "all the kings horses and all the kings men",
        map[string]int {
            "all": 2,
            "the": 2,
            "kings": 2,
            "horses": 1,
            "and": 1,
            "men": 1 },
    }, {
        "ignore punctuation",
        "car : carpet as java : javascript!!&@$%^&",
        map[string]int {
            "as": 1,
            "car": 1,
            "carpet": 1,
            "java": 1,
            "javascript": 1 },
    }, {
        "handle cramped lists",
        "one,two,three",
        map[string]int {
            "one": 1,
            "two": 1,
            "three": 1 },
    }, {
        "include numbers",
        "testing, 1, 2 testing",
        map[string]int {
            "1": 1,
            "2": 1,
            "testing": 2 },
    }, {
        "normalize case",
        "go Go GO",
        map[string]int { "go": 3 },
    }, {
        "allow apostrophes",
        "First: don't laugh. Then: don't cry.",
        map[string]int {
            "first": 1,
            "don't": 2,
            "laugh": 1,
            "then": 1,
            "cry": 1 },
    }, {
        "treat symbols as separators",
        "hey,my_spacebar_is_broken.",
        map[string]int {
            "hey": 1,
            "my": 1,
            "spacebar": 1,
            "is": 1,
            "broken": 1 },
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

func TestWordCount(t *testing.T) {
    for _, test := range tests {
        var counts = Count(test.input)
        if equal(counts, test.expected) {
            t.Logf("Passed: %s", test.description)
        } else {
            t.Fatalf("Failed: %s, expected %v but it was %v", test.description, test.expected, counts)
        }
    }
}

func BenchmarkWordCount(b *testing.B) {
    for i := 0; i < b.N; i++ {
        for _, test := range tests {
            Count(test.input)
        }
    }
}
