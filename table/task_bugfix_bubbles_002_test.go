package table

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles002SourceContract(t *testing.T) {
    source, err := os.ReadFile("table.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if m.cols[i].Width <= 0 {") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "if m.cols[i].Width < 0 {") {
        t.Fatalf("mutated source contract is still present")
    }
}
