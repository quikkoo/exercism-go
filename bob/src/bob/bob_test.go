package bob

import "testing"

var tests = []struct {
    description string
    question    string
    expected    string
} {
    {
        "should respond to a statement",
        "Tom-ay-to, tom-aaaah-to.",
        "Whatever.",
    }, {
        "should respond to shouting",
        "WATCH OUT!",
        "Whoa, chill out!",
    }, {
        "should respond to questions",
        "Does this cryogenic chamber make me look fat?",
        "Sure.",
    }, {
        "should respond to questions ending with numbers",
        "You are, what, like 15?",
        "Sure.",
    }, {
        "should respond to forceful talking",
        "Let's go make out behind the gym!",
        "Whatever.",
    }, {
        "should respond to acronyms in regular speech",
        "It's OK if you don't want to go to the DMV.",
        "Whatever.",
    }, {
        "should respond to forceful questions as shouting",
        "WHAT THE HELL WERE YOU THINKING?",
        "Whoa, chill out!",
    }, {
        "should respond to shouting with special characters",
        "ZOMG THE %^*@#$(*^ ZOMBIES ARE COMING!!11!!1!",
        "Whoa, chill out!",
    }, {
        "should respond to numbers when shouting",
        "1, 2, 3 GO!",
        "Whoa, chill out!",
    }, {
        "should respond to numbers as speech",
        "1, 2, 3",
        "Whatever.",
    }, {
        "should respond to questions with only numbers",
        "4?",
        "Sure.",
    }, {
        "should respond to shouting with no exclamation mark",
        "I HATE YOU",
        "Whoa, chill out!",
    }, {
        "should respond to statement containing question mark",
        "Ending with ? means a question.",
        "Whatever.",
    }, {
        "should respond to prattling on",
        "Wait! Hang on. Are you going to be OK?",
        "Sure.",
    }, {
        "should respond to silence",
        "",
        "Fine. Be that way!",
    }, {
        "should respond to prolonged silence",
        "          ",
        "Fine. Be that way!",
    }, {
        "should respond to others blank characters",
        "\n\r \t\v\u00A0\u2002",
        "Fine. Be that way!",
    }, {
        "should respond to multiple line question",
        "\nDoes this cryogenic chamber make me look fat?\nno",
        "Whatever.",
    }, {
        "should respond to non-letters with question",
        ":) ?",
        "Sure.",
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
