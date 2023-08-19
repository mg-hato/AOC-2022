package solver

import (
	c "aoc/common"
	m "aoc/d21/models"
	"fmt"
)

func create_cached_solver(graph map[string]m.MonkeyJob) func(string) (int64, error) {
	cache := make(map[string]int64)

	id_resolver := func(id string) (int64, error) {
		val, is_cached := cache[id]
		if !is_cached {
			return 0, fmt.Errorf(`identifier "%s" is not cached`, id)
		}
		return val, nil
	}

	return func(request_id string) (int64, error) {
		stack := c.MakeStack(request_id)
		unresolved_ancestors := make(map[string]bool)
		for !stack.IsEmpty() {
			id, _ := stack.Pop()

			// If id is already in cache, nothing to do
			if _, contains := cache[id]; contains {
				unresolved_ancestors[id] = false
				continue
			}

			unresolved := c.Filter(func(dep_id string) bool {
				_, resolved := cache[dep_id]
				return !resolved
			}, graph[id].GetIdentifiers())

			// if id is not in cache, but its dependents are all resolved
			if len(unresolved) == 0 {
				value, err := graph[id].Calculate(id_resolver)
				if err != nil {
					return 0, err
				}
				cache[id] = value
				unresolved_ancestors[id] = false
				continue
			}

			// if id was already marked as an ancestor that is not resolved
			// we have a circular dependency
			if unresolved_ancestors[id] {
				return 0, fmt.Errorf(`graph resolution failed; circular dependency on "%s"`, id)
			}

			// otherwise, prioritise resolution of dependents
			// but mark current id as ancestor that is not resolved
			// in order to detect circular dependencies
			unresolved_ancestors[id] = true
			stack.Push(id, unresolved...)
		}
		return cache[request_id], nil
	}
}
