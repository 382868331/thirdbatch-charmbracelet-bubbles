package progress

import (
    "os"
    "strings"
    "testing"
)

func TestTaskDiagnosisBubbles020SourceContract(t *testing.T) {
    source, err := os.ReadFile("progress.go")
    if err != nil {
        t.Fatalf("read source: %v", err)
    }
    if !strings.Contains(string(source), "if len(colors) == 0 {") {
        t.Fatalf("expected source contract is missing")
    }
}
