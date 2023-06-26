package solver

import (
	c "aoc/common"
	m "aoc/d18/models"
)

func generate_water_droplets(lava_droplets map[m.Droplet]bool) map[m.Droplet]bool {
	water_droplets := map[m.Droplet]bool{}
	corners, in_range_func := encapsulating_cube_edges(c.GetKeys(lava_droplets))
	q := c.Queue[m.Droplet](corners...)
	for !q.IsEmpty() {
		water_droplet, _ := q.Dequeue()
		water_droplets[water_droplet] = true
		for _, new_droplet := range water_droplet.GetAllNeighbours() {
			if !water_droplets[new_droplet] && !lava_droplets[new_droplet] && in_range_func(new_droplet) {
				water_droplets[new_droplet] = true
				q.Enqueue(new_droplet)
			}
		}
	}
	return water_droplets
}

func encapsulating_cube_edges(droplets []m.Droplet) ([]m.Droplet, func(m.Droplet) bool) {
	getx := func(d m.Droplet) int { return d.X }
	gety := func(d m.Droplet) int { return d.Y }
	getz := func(d m.Droplet) int { return d.Z }

	xs, ys, zs := c.Map(getx, droplets), c.Map(gety, droplets), c.Map(getz, droplets)

	min_x, max_x := c.Minimum(xs), c.Maximum(xs)
	min_y, max_y := c.Minimum(ys), c.Maximum(ys)
	min_z, max_z := c.Minimum(zs), c.Maximum(zs)

	x_in_range := c.InInclusiveRange(min_x-1, max_x+1)
	y_in_range := c.InInclusiveRange(min_y-1, max_y+1)
	z_in_range := c.InInclusiveRange(min_z-1, max_z+1)

	return []m.Droplet{
			m.MakeDroplet(min_x-1, min_y-1, min_z-1),
			m.MakeDroplet(max_x+1, min_y-1, min_z-1),

			m.MakeDroplet(min_x-1, max_y+1, min_z-1),
			m.MakeDroplet(max_x+1, max_y+1, min_z-1),

			m.MakeDroplet(min_x-1, min_y-1, max_z+1),
			m.MakeDroplet(max_x+1, min_y-1, max_z+1),

			m.MakeDroplet(min_x-1, max_y+1, max_z+1),
			m.MakeDroplet(max_x+1, max_y+1, max_z+1),
		}, func(d m.Droplet) bool {
			return x_in_range(d.X) && y_in_range(d.Y) && z_in_range(d.Z)
		}
}
