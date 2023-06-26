package solver

import (
	c "aoc/common"
	m "aoc/d18/models"
)

func CountAreaOfWaterAdjacentSurfaces(envelope m.SolverInput) (int, error) {
	lava_droplets := c.CreateSet(envelope.Get(), c.Identity[m.Droplet])
	water_droplets := generate_water_droplets(lava_droplets)

	area := 0
	for droplet := range lava_droplets {
		area += c.Count(
			droplet.GetAllNeighbours(),
			func(d m.Droplet) bool { return water_droplets[d] },
		)
	}

	return area, nil
}
