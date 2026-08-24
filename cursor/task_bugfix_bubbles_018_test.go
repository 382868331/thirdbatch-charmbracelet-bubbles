package cursor

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles018SourceContract(t *testing.T) {
    source, err := os.ReadFile("cursor.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if m.blinkCtx != nil && m.blinkCtx.cancel != nil {") {
        t.Fatalf("expected source contract is missing")
    }
}
