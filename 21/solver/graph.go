package solver

import (
	c "aoc/common"
	m "aoc/d21/models"
)

func make_graph(input []c.Pair[string, m.MonkeyJob]) map[string]m.MonkeyJob {
	return c.CreateKeyValueMap(
		input,
		c.GetFirst[string, m.MonkeyJob],
		c.GetSecond[string, m.MonkeyJob],
	)
}

func build_dependents(graph map[string]m.MonkeyJob) map[string][]string {
	dependents := c.AssociateWith(
		c.GetKeys(graph),
		func(id string) []string {
			return []string{}
		},
	)
	for job_id, job := range graph {
		for _, id := range job.GetIdentifiers() {
			dependents[id] = append(dependents[id], job_id)
		}
	}
	return dependents
}

func get_dependents_of(requested_id string, graph map[string]m.MonkeyJob) map[string]bool {
	direct_dependents := build_dependents(graph)
	is_dependent := make(map[string]bool)
	stack := c.MakeStack(requested_id)
	for !stack.IsEmpty() {
		id, _ := stack.Pop()
		if !is_dependent[id] {
			is_dependent[id] = true
			if len(direct_dependents[id]) > 0 {
				stack.Push(direct_dependents[id][0], direct_dependents[id][1:]...)
			}
		}
	}
	return is_dependent
}
