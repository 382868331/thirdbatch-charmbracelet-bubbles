package textinput

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles009SourceContract(t *testing.T) {
    source, err := os.ReadFile("textinput.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(value) < len(suggestion) {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if len(value) <= len(suggestion) {") {
        t.Fatalf("mutated source contract is still present")
    }
}
