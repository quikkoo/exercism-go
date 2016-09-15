package helloworld

import "testing"

var tests = []struct {
    description string
    name        string
    expected    string
} {
    {
        "With empty name",
        "",
        "Hello, World!",
    }, {
        "With name Alice",
        "Alice",
        "Hello, Alice!",
    },
}

func TestHelloWorld(t *testing.T) {
    for _, test := range tests {
        var greeting = Greet(test.name)
        if greeting == test.expected {
            t.Logf("Passed: %s", test.description)
        } else {
            t.Fatalf("Failed: %s, expected %v but it was %v", test.description, test.expected, greeting)
        }
    }
}

func BenchmarkHelloWorld(b *testing.B) {
    for _, test := range tests {
        for i := 0; i < b.N; i++ {
            Greet(test.name)
        }
    }
}
