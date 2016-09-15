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
        description: "detect no matches",
        subject: "diaper",
        candidates: []string {
            "hello", "world", "zombies", "pants",
        },
        expected: []string {},
    }, {
        description: "detect simple anagram",
        subject: "ant",
        candidates: []string {
            "tan", "stand", "at",
        },
        expected: []string { "tan" },
    }, {
        description: "detect multiple anagrams",
        subject: "master",
        candidates: []string {
            "stream", "pigeon", "maters",
        },
        expected: []string { "stream", "maters"},
    }, {
        description: "does not confuse different duplicates",
        subject: "galea",
        candidates: []string {
            "eagle",
        },
        expected: []string {},
    }, {
        description: "does not include identical words",
        subject: "corn",
        candidates: []string {
            "corn", "dark", "Corn", "rank", "CORN", "cron", "park",
        },
        expected: []string { "cron" },
    }, {
        description: "eliminate anagrams with same checksum",
        subject: "mass",
        candidates: []string {
            "last",
        },
        expected: []string {},
    }, {
        description: "eliminate anagram subsets",
        subject: "good",
        candidates: []string {
            "dog", "goody",
        },
        expected: []string {},
    }, {
        description: "detect anagrams",
        subject: "listen",
        candidates: []string {
            "enlists", "google", "inlets", "banana",
        },
        expected: []string { "inlets" },
    }, {
        description: "detect more anagrams",
        subject: "allergy",
        candidates: []string {
            "gallery", "ballerina", "regally", "clergy", "largely", "leading",
        },
        expected: []string { "gallery", "regally", "largely" },
    }, {
        description: "treat subject anagrams as case insensitive",
        subject: "Orchestra",
        candidates: []string {
            "cashregiser", "carthorse", "radishes",
        },
        expected: []string { "carthorse" },
    }, {
        description: "treat candidates anagrams as case insensitive",
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
        for i := 0; result && i < len(actual) && i < len(expected); i++ {
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
