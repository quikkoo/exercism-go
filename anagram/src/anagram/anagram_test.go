package anagram

import (
    "strings"
    "testing"
)

var tests = []struct {
    description string
    subject     string
    candidates  []string
    expected    []string
} {
    {
        description: "should detect no matches",
        subject: "diaper",
        candidates: []string {
            "hello", "world", "zombies", "pants",
        },
        expected: []string {},
    }, {
        description: "should detect simple anagram",
        subject: "ant",
        candidates: []string {
            "tan", "stand", "at",
        },
        expected: []string { "tan" },
    }, {
        description: "should detect multiple anagrams",
        subject: "master",
        candidates: []string {
            "stream", "pigeon", "maters",
        },
        expected: []string { "stream", "maters"},
    }, {
        description: "should not confuse different duplicates",
        subject: "galea",
        candidates: []string {
            "eagle",
        },
        expected: []string {},
    }, {
        description: "should does not include identical words",
        subject: "corn",
        candidates: []string {
            "corn", "dark", "Corn", "rank", "CORN", "cron", "park",
        },
        expected: []string { "cron" },
    }, {
        description: "should eliminate anagram subsets",
        subject: "good",
        candidates: []string {
            "dog", "goody",
        },
        expected: []string {},
    }, {
        description: "should eliminate anagrams with same checksum",
        subject: "mass",
        candidates: []string {
            "last",
        },
        expected: []string {},
    }, {
        description: "should detect anagrams",
        subject: "listen",
        candidates: []string {
            "enlists", "google", "inlets", "banana",
        },
        expected: []string { "inlets" },
    }, {
        description: "should detect more multiple anagrams",
        subject: "allergy",
        candidates: []string {
            "gallery", "ballerina", "regally", "clergy", "largely", "leading",
        },
        expected: []string { "gallery", "regally", "largely" },
    }, {
        description: "should treat subject anagrams as case insensitive",
        subject: "Orchestra",
        candidates: []string {
            "cashregiser", "carthorse", "radishes",
        },
        expected: []string { "carthorse" },
    }, {
        description: "should treat candidates anagrams as case insensitive",
        subject: "orchestra",
        candidates: []string {
            "Cashregiser", "Carthorse", "Radishes",
        },
        expected: []string { "Carthorse" },
    },
}

func equal(actual []string, expected []string) bool {
    if len(actual) != len(expected) {
        return false
    }

    var result = true;
    for i := 0; result && i < len(actual); i++ {
            result = strings.Compare(actual[i], expected[i]) == 0
    }

    return result;
}

func TestAnagram(t *testing.T) {
    for _, test := range tests {
        var anagrams = Matches(test.subject, test.candidates)
        if equal(anagrams, test.expected) {
            t.Logf("Passed: %s", test.description)
        } else {
            t.Fatalf("Failed: %s, expected %v but it was %v", test.description, test.expected, anagrams)
        }
    }
}

func BenchmarkAnagram(b *testing.B) {
    for _, test := range tests {
        for i := 0; i < b.N; i++ {
            Matches(test.subject, test.candidates)
        }
    }
}
