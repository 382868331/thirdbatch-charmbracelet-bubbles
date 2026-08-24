package runeutil

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles003SourceContract(t *testing.T) {
    source, err := os.ReadFile("runeutil.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "case r == '\\t':") {
        t.Fatalf("expected source contract is missing")
    }
}
