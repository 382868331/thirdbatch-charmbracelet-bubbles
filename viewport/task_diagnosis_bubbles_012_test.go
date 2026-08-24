package viewport

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBubbles012SourceContract(t *testing.T) {
    source, err := os.ReadFile("viewport.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if sh := m.Style.GetHeight(); sh != 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
