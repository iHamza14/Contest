package problem

import (
	"os"
	"testing"
)

func TestParseProblem(t *testing.T) {
	file, err := os.Open("./test.yaml") // Adjust path as needed
	if err != nil {
		t.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	contest, err := ParseProblem(file)

	if err != nil {
		t.Fatalf("ParseProblem returned an error: %v", err)
	}
	t.Logf("Successfully parsed Contest ID: %s", contest.ContestId)
	if len(contest.Problems) == 0 {
		t.Errorf("Expected problems to be parsed, got 0")
	}
}
