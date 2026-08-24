package textarea

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles006SourceContract(t *testing.T) {
    source, err := os.ReadFile("textarea.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if unicode.IsSpace(m.value[m.row][m.col-1]) {") {
        t.Fatalf("expected source contract is missing")
    }
}
