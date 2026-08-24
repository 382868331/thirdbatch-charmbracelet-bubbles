package textarea

import (
    "os"
    "strings"
    "testing"
)

func TestTaskBugfixBubbles015SourceContract(t *testing.T) {
    source, err := os.ReadFile("textarea.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "for end < len(line) && !unicode.IsSpace(line[end]) {") {
        t.Fatalf("expected source contract is missing")
    }
}
