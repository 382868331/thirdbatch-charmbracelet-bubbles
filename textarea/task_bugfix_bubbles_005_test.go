package textarea

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles005SourceContract(t *testing.T) {
    source, err := os.ReadFile("textarea.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(m.Value()) == 0 && m.row == 0 && m.col == 0 && m.Placeholder != \"\" {") {
        t.Fatalf("expected source contract is missing")
    }
}
