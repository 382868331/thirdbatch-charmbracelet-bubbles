package textinput

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBubbles004SourceContract(t *testing.T) {
    source, err := os.ReadFile("textinput.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "m.SetCursor(m.pos - 1)") {
        t.Fatalf("expected source contract is missing")
    }
    if strings.Contains(string(source), "m.SetCursor(m.pos + 1)") {
        t.Fatalf("mutated source contract is still present")
    }
}
