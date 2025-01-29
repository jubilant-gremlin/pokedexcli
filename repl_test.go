package main

import(
    "testing"
)

func TestCleanInput(t *testing.T) {
    //create a slice of test case structs
    cases := []struct {
        input string
        expected []string
    }{
        {
            input: "  hello world  ",
            expected: []string{"hello", "world"},
        },
        {
            input: "TONY CARrotS ",
            expected: []string{"tony", "carrots"},
        },
        {
            input: " farmer Cole is a Big DOO doo heaD  ",
            expected: []string{"farmer", "cole", "is", "a", "big", "doo", "doo", "head"},
        },
    }

    //loop over the cases and run the tests
    for _, c := range cases {
        actual := CleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("actual: %v, expected: %v", actual, c.expected)
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("actual: %v, expected: %v", word, expectedWord)
            }
        }
    }
}

