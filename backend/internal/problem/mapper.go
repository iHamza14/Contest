package problem

import "contest-platform/internal/models"

func ToContestModel(c ContestYAML) models.Contest {

	var problems []models.Problem

	for _, p := range c.Problems {

		var testcases []models.TestCase

		// Visible testcases
		for _, tc := range p.VisibleTestcases {

			testcases = append(testcases, models.TestCase{
				InputFile:  tc.InputFile,
				OutputFile: tc.OutputFile,
				IsHidden:   false,
			})
		}

		// Hidden testcases
		for _, tc := range p.HiddenTestcases {

			testcases = append(testcases, models.TestCase{
				InputFile:  tc.InputFile,
				OutputFile: tc.OutputFile,
				IsHidden:   true,
			})
		}

		problems = append(problems, models.Problem{
			ProblemID:        p.ProblemID,
			ProblemTitle:     p.ProblemTitle,
			ProblemStatement: p.ProblemStatement,
			TimeLimitMS:      p.TimeLimitMS,
			MemoryLimitMB:    p.MemoryLimitMB,
			TestCases:        testcases,
		})
	}

	return models.Contest{
		ContestID: c.ContestID,
		Title:     c.Title,
		StartTime: c.StartTime,
		EndTime:   c.EndTime,
		Problems:  problems,
	}
}
