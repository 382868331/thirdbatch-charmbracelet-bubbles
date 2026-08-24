package help

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles011SourceContract(t *testing.T) {
    source, err := os.ReadFile("help.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if totalWidth > 0 && i < len(groups) {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if totalWidth > 0 || i < len(groups) {") {
        t.Fatalf("mutated source contract is still present")
    }
}
