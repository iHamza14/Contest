package problem

import (
	"errors"
)

func ValidateContest(c ContestYAML) error {

	if c.ContestID == "" {
		return errors.New("contest_id required")
	}

	if c.Title == "" {
		return errors.New("contest title required")
	}

	if c.EndTime.Before(c.StartTime) {
		return errors.New("end time before start time")
	}

	for _, p := range c.Problems {

		if p.ProblemID == "" {
			return errors.New("problem_id required")
		}

		if p.TimeLimitMS <= 0 {
			return errors.New("invalid time limit")
		}

		if p.MemoryLimitMB <= 0 {
			return errors.New("invalid memory limit")
		}

		// if len(p.HiddenTestcases) == 0 {
		// 	return errors.New("hidden testcases required")
		// }
	}

	return nil
}
