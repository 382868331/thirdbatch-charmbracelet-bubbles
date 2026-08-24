package viewport

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles017SourceContract(t *testing.T) {
    source, err := os.ReadFile("viewport.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(subLines) > 1 {") {
        t.Fatalf("expected source contract is missing")
    }
}
