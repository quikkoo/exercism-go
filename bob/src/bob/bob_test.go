package bob

import "testing"

var tests = []struct {
    description string
    question    string
    expected    string
} {
    {
        "stating something",
        "Tom-ay-to, tom-aaaah-to.",
        "Whatever.",
    }, {
        "shouting",
        "WATCH OUT!",
        "Whoa, chill out!",
    }, {
        "shouting gibberish",
        "FCECDFCAAB",
        "Whoa, chill out!",
    }, {
        "asking a question",
        "Does this cryogenic chamber make me look fat?",
        "Sure.",
    }, {
        "asking a numeric question",
        "You are, what, like 15?",
        "Sure.",
    }, {
        "asking gibberish",
        "fffbbcbeab?",
        "Sure.",
    }, {
        "talking forcefully",
        "Let's go make out behind the gym!",
        "Whatever.",
    }, {
        "using acronyms in regular speech",
        "It's OK if you don't want to go to the DMV.",
        "Whatever.",
    }, {
        "forceful question",
        "WHAT THE HELL WERE YOU THINKING?",
        "Whoa, chill out!",
    }, {
        "shouting numbers",
        "1, 2, 3 GO!",
        "Whoa, chill out!",
    }, {
        "only numbers",
        "1, 2, 3",
        "Whatever.",
    }, {
        "question with only numbers",
        "4?",
        "Sure.",
    }, {
        "shouting with special characters",
        "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
        "Whoa, chill out!",
    }, {
        "shouting with umlauts",
        "ÜMLÄÜTS!",
        "Whoa, chill out!",
    }, {
        "calmly speaking with umlauts",
        "ÜMLäÜTS!",
        "Whatever.",
    }, {
        "shouting with no exclamation mark",
        "I HATE YOU",
        "Whoa, chill out!",
    }, {
        "statement containing question mark",
        "Ending with ? means a question.",
        "Whatever.",
    }, {
        "non-letters with question",
        ":) ?",
        "Sure.",
    }, {
        "prattling on",
        "Wait! Hang on. Are you going to be OK?",
        "Sure.",
    }, {
        "silence",
        "",
        "Fine. Be that way!",
    }, {
        "prolonged silence",
        "                    ",
        "Fine. Be that way!",
    }, {
        "alternate silence",
        "\t\t\t\t\t\t\t\t\t\t",
        "Fine. Be that way!",
    }, {
        "multiple line question",
        "\nDoes this cryogenic chamber make me look fat?\nno",
        "Whatever.",
    },     {
        "starting with whitespace",
        "                 hmmmmmmm...",
        "Whatever.",
    }, {
        "ending with whitespace",
        "Okay if like my    spacebar    quite a bit?     ",
        "Sure.",
    }, {
        "other whitespace",
        "\n\r \t\v\u00a0\u2002",
        "Fine. Be that way!",
    },
}

func TestBob(t *testing.T) {
    for _, test := range tests {
        var response = Hey(test.question)
        if response == test.expected {
            t.Logf("Passed: %s", test.description)
        } else {
            t.Fatalf("Failed: %s, expected %v but it was %v", test.description, test.expected, response)
        }
    }
}

func BenchmarkBob(b *testing.B) {
    for _, test := range tests {
        for i := 0; i < b.N; i++ {
            Hey(test.question)
        }
    }
}
