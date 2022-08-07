package _721

import (
	"sort"
)

func accountsMerge(accountList [][]string) [][]string {
	solver := NewSolver(accountList)
	return solver.Solve(accountList)
}

type Solver struct {
	visited  HashSet
	adjacent Adjacent
}

func NewSolver(accountList [][]string) *Solver {
	var visited = make(HashSet, 0)
	var adjacent = make(Adjacent, 0)
	var accountFirstEmail string
	var accountEmail string

	for _, account := range accountList {
		accountFirstEmail = account[1]
		for j := 2; j < len(account); j++ {
			accountEmail = account[j]

			adjacent.Add(accountFirstEmail, accountEmail)
			adjacent.Add(accountEmail, accountFirstEmail)
		}
	}

	return &Solver{visited: visited, adjacent: adjacent}
}

func (s *Solver) Solve(accountList [][]string) [][]string {
	var (
		accountName       string
		accountFirstEmail string
		mergedAccounts    = make([][]string, 0)
	)

	for _, account := range accountList {
		accountName = account[0]
		accountFirstEmail = account[1]

		if s.visited.Contains(accountFirstEmail) {
			continue
		}

		mergedAccounts = append(mergedAccounts, s.getMergedAccount(accountName, accountFirstEmail))
	}

	return mergedAccounts
}

func (s *Solver) getMergedAccount(accountName string, accountFirstEmail string) []string {
	mergedAccount := make([]string, 1)
	mergedAccount[0] = accountName

	stack := Stack{accountFirstEmail}
	set := HashSet{}
	for !stack.Empty() {
		email := stack.Pop()
		s.visited.Add(email)

		if !set.Contains(email) {
			mergedAccount = append(mergedAccount, email)
			set.Add(email)
		}

		if !s.adjacent.Contains(email) {
			break
		}

		for _, edge := range s.adjacent[email] {
			if !s.visited.Contains(edge) {
				stack.Push(edge)
			}
		}
	}

	emails := mergedAccount[1:]
	sort.Slice(emails, func(i, j int) bool {
		return emails[i] < emails[j]
	})
	return append([]string{accountName}, emails...)
}

type HashSet map[string]struct{}

func (s *HashSet) Contains(key string) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *HashSet) Add(value string) {
	if _, ok := (*s)[value]; !ok {
		(*s)[value] = struct{}{}
	}
}

type Adjacent map[string][]string

func (s *Adjacent) Contains(key string) bool {
	_, ok := (*s)[key]
	return ok
}

func (s *Adjacent) Add(key, value string) {
	if _, ok := (*s)[key]; !ok {
		(*s)[key] = make([]string, 0)
	}
	(*s)[key] = append((*s)[key], value)
}

type Stack []string

func (s *Stack) Push(value string) {
	*s = append(*s, value)
}

func (s *Stack) Pop() string {
	last := s.Size() - 1
	value := (*s)[last]
	*s = (*s)[:last]
	return value
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}

func (s *Stack) Size() int {
	return len(*s)
}
