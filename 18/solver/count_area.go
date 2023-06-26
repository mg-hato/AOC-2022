package solver

import (
	c "aoc/common"
	m "aoc/d18/models"
)

func CountAreaOfDropletSurfaces(envelope m.SolverInput) (int, error) {
	droplets := c.CreateSet(envelope.Get(), c.Identity[m.Droplet])

	area := 0
	for droplet, _ := range droplets {
		area += 6 - c.Count(
			droplet.GetAllNeighbours(),
			func(d m.Droplet) bool { return droplets[d] },
		)
	}

	return area, nil
}
