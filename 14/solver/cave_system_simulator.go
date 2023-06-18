package solver

type CaveSystemSimulator interface {
	// simulate dropping of the sand;
	// returns boolean indicating whether the sand was added successfully
	dropSandUnit() bool
}
