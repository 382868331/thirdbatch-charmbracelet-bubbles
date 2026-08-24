package paginator

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles010SourceContract(t *testing.T) {
    source, err := os.ReadFile("paginator.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "return m.Page == m.TotalPages-1") {
        t.Fatalf("expected source contract is missing")
    }
}
