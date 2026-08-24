package textinput

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles014SourceContract(t *testing.T) {
    source, err := os.ReadFile("textinput.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "Prompt:           \"> \",") {
        t.Fatalf("expected source contract is missing")
    }
}
