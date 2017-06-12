package wordcount

import "testing"

var tests = []struct {
    description string
    input       string
    expected    map[string]int
} {
    {
        "should count one word",
        "word",
        map[string]int { "word": 1 },
    }, {
        "should count one of each",
        "one of each",
        map[string]int {
            "each": 1,
            "of": 1,
            "one": 1,
        },
    }, {
        "should count multiple occurrences",
        "one fish two fish red fish blue fish",
        map[string]int {
            "blue": 1,
            "fish": 4,
            "one": 1,
            "red": 1,
            "two": 1,
        },
    }, {
        "should count everything just once",
        "all the kings horses and all the kings men",
        map[string]int {
            "all": 2,
            "the": 2,
            "kings": 2,
            "horses": 1,
            "and": 1,
            "men": 1,
        },
    }, {
        "should ignore punctuation",
        "car : carpet as java : javascript!!&@$%^&",
        map[string]int {
            "as": 1,
            "car": 1,
            "carpet": 1,
            "java": 1,
            "javascript": 1,
        },
    }, {
        "should handle cramped lists",
        "one,two,three",
        map[string]int {
            "one": 1,
            "two": 1,
            "three": 1,
        },
    }, {
        "should include numbers",
        "testing, 1, 2 testing",
        map[string]int {
            "1": 1,
            "2": 1,
            "testing": 2,
        },
    }, {
        "should normalize case",
        "go Go GO",
        map[string]int { "go": 3 },
    }, {
        "should allow apostrophes",
        "First: don't laugh. Then: don't cry.",
        map[string]int {
            "first": 1,
            "don't": 2,
            "laugh": 1,
            "then": 1,
            "cry": 1,
        },
    }, {
        "should treat symbols as separators",
        "hey,my_spacebar_is_broken.",
        map[string]int {
            "hey": 1,
            "my": 1,
            "spacebar": 1,
            "is": 1,
            "broken": 1,
        },
    }, {
        "should counts words with quotations",
        "Joe can't tell between 'large' and large.",
        map[string]int {
            "joe": 1,
            "can't": 1,
            "tell": 1,
            "between": 1,
            "large": 2,
            "and": 1,
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

func TestWordCount(t *testing.T) {
    for _, test := range tests {
        var counts = Calculate(test.input)
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
            Calculate(test.input)
        }
    }
}
