package models

func GenerateShapeAt(
	height int,
	leftmost_position int,
	shape_supplier func() Shape,
) Shape {
	shape := shape_supplier()
	for i := range shape {
		shape[i] = MakePoint(HeightOf(shape[i])+height, PositionOf(shape[i])+leftmost_position)
	}
	return shape
}

func HorizontalLine() Shape {
	return []Point{
		MakePoint(0, 0),
		MakePoint(0, 1),
		MakePoint(0, 2),
		MakePoint(0, 3),
	}
}

func PlusShape() Shape {
	return []Point{
		MakePoint(0, 1),
		MakePoint(1, 0),
		MakePoint(1, 1),
		MakePoint(1, 2),
		MakePoint(2, 1),
	}
}

func L_Shape() Shape {
	return []Point{
		MakePoint(0, 0),
		MakePoint(0, 1),
		MakePoint(0, 2),
		MakePoint(1, 2),
		MakePoint(2, 2),
	}
}

func VerticalLine() Shape {
	return []Point{
		MakePoint(0, 0),
		MakePoint(1, 0),
		MakePoint(2, 0),
		MakePoint(3, 0),
	}
}

func Square() Shape {
	return []Point{
		MakePoint(0, 0),
		MakePoint(0, 1),
		MakePoint(1, 0),
		MakePoint(1, 1),
	}
}
