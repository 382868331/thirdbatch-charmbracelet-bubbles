package textarea

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles019SourceContract(t *testing.T) {
    source, err := os.ReadFile("selection.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(m.value) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
