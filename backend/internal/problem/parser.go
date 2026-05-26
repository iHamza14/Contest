package problem

import (
	"io"
	"time"

	"go.yaml.in/yaml/v4"
)

type ContestYAML struct {
	Problems []ProblemYAML `yaml:"problems"`

	ContestID string `yaml:"contest_id"`
	Title     string `yaml:"title"`

	StartTime time.Time `yaml:"start_time"`
	EndTime   time.Time `yaml:"end_time"`
}

type ProblemYAML struct {
	ProblemID        string `yaml:"problem_id"`
	ProblemTitle     string `yaml:"problem_title"`
	ProblemStatement string `yaml:"problem_statement"`

	VisibleTestcases []TestCaseYAML `yaml:"visible_testcases"`
	HiddenTestcases  []TestCaseYAML `yaml:"hidden_testcases"`

	TimeLimitMS   int `yaml:"time_limit_ms"`
	MemoryLimitMB int `yaml:"memory_limit_mb"`
}

type TestCaseYAML struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

func ParseProblem(yamlReader io.Reader) (ContestYAML, error) {

	var c ContestYAML

	decoder := yaml.NewDecoder(yamlReader)
	decoder.KnownFields(true)

	err := decoder.Decode(&c)

	if err != nil {
		return ContestYAML{}, err
	}

	return c, nil
}
