package cli

import (
	"testing"
)

func TestFrameworksList(t *testing.T) {
	if len(Frameworks) == 0 {
		t.Fatal("expected at least one framework")
	}

	expected := []string{"Gin", "Fiber", "Chi", "Echo", "Gorilla/Mux", "Standard Library"}
	if len(Frameworks) != len(expected) {
		t.Errorf("expected %d frameworks, got %d", len(expected), len(Frameworks))
	}

	for i, f := range Frameworks {
		if i < len(expected) && f.Name != expected[i] {
			t.Errorf("expected framework %d to be %s, got %s", i, expected[i], f.Name)
		}
		if f.DirName == "" {
			t.Errorf("framework %s has empty directory name", f.Name)
		}
		if f.Port == "" {
			t.Errorf("framework %s has empty port", f.Name)
		}
	}
}

func TestFrameworksUniqueDirs(t *testing.T) {
	dirs := make(map[string]bool)
	for _, f := range Frameworks {
		if dirs[f.DirName] {
			t.Errorf("duplicate directory name: %s", f.DirName)
		}
		dirs[f.DirName] = true
	}
}
