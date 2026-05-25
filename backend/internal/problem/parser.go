package problem

import (
	"io"
	"time"

	"go.yaml.in/yaml/v4"
)

type Contest struct {
	Problems  []Problem `yaml:"problems"`
	ContestId string    `yaml:"contest_id"`
	Title     string    `yaml:"title"`
	StartTime time.Time `yaml:"start_time"`
	EndTime   time.Time `yaml:"end_time"`
}

type Problem struct {
	ProblemId        string     `yaml:"problem_id"`
	ProblemTitle     string     `yaml:"problem_title"`
	ProblemStatement string     `yaml:"problem_statement"`
	TestcaseV        []TestCase `yaml:"testcase_v"`
	TestcaseH        []TestCase `yaml:"testcase_h"`
	TimeLimitMs      int        `yaml:"time_limit_ms"`
	MemoryLimitMB    int        `yaml:"memory_limit_mb"`
}

type TestCase struct {
	InputFile  string `yaml:"input_file"`
	OutputFile string `yaml:"output_file"`
}

func ParseProblem(yamlReader io.Reader) (contest Contest, e error) {
	var c Contest

	decoder := yaml.NewDecoder(yamlReader)
	decoder.KnownFields(true)
	err := decoder.Decode(&c)

	if err != nil {
		return contest, err
	}

	return c, nil
}
