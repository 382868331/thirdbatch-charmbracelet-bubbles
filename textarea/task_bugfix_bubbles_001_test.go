package textarea

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles001SourceContract(t *testing.T) {
    source, err := os.ReadFile("textarea.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "minWidth := reservedInner + reservedOuter + 1") {
        t.Fatalf("expected source contract is missing")
    }
}
