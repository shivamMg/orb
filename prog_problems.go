package main

import (
	"io/ioutil"
	"strings"
)

func getNextProblem() string {
	x := 0
	lastProblem := readLastProblem()
	problems := readProblems()
	for i, p := range problems {
		if p == lastProblem {
			// Keep x in range(0, len(problems))
			x = (i + 1) % len(problems)
		}
	}
	nextProblem := problems[x]
	return nextProblem
}

func readLastProblem() string {
	filename := assetsDir + "last_problem.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}

	problem := strings.TrimRight(string(content), "\n")
	return problem
}

func readProblems() []string {
	dirname := assetsDir + "problems"
	content, err := ioutil.ReadDir(dirname)
	if err != nil {
		return []string{}
	}

	var problems []string
	for _, p := range content {
		problems = append(problems, p.Name())
	}
	return problems
}
