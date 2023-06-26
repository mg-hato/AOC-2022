package models

type Droplet struct {
	X, Y, Z int
}

func MakeDroplet(x, y, z int) Droplet {
	return Droplet{
		X: x,
		Y: y,
		Z: z,
	}
}

func (droplet Droplet) GetAllNeighbours() []Droplet {
	return []Droplet{
		MakeDroplet(droplet.X+1, droplet.Y, droplet.Z),
		MakeDroplet(droplet.X-1, droplet.Y, droplet.Z),
		MakeDroplet(droplet.X, droplet.Y+1, droplet.Z),
		MakeDroplet(droplet.X, droplet.Y-1, droplet.Z),
		MakeDroplet(droplet.X, droplet.Y, droplet.Z+1),
		MakeDroplet(droplet.X, droplet.Y, droplet.Z-1),
	}
}
