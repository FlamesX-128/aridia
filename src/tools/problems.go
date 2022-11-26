package tools

import models "github.com/FlamesX-128/aridia/src/models/database"

func ExtendPostProblem(problem models.PostProblem) models.Problem {
	return models.Problem{
		Description: problem.Description,
		Explanation: problem.Explanation,
		AuthorID:    problem.AuthorID,
		Answer:      problem.Answer,
		Community:   true,
		Verified:    false,
		CreatedAt:   GetCurrentTime(),
	}
}

func ExtendPutProblem(problem models.PutProblem) models.Problem {
	return models.Problem{
		Description: problem.Description,
		Explanation: problem.Explanation,
		AuthorID:    problem.AuthorID,
		Answer:      problem.Answer,
		Community:   true,
		Verified:    false,
		CreatedAt:   GetCurrentTime(),
	}
}
