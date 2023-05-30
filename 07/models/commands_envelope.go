package models

import (
	e "aoc/envelope"
	f "aoc/functional"
)

type CommandsEnvelope struct {
	commands []Command
}

func (ce CommandsEnvelope) Get() []Command {
	return f.Map(Command.Copy, ce.commands)
}

func CreateCommandsEnvelope(commands ...Command) e.Envelope[[]Command] {
	return CommandsEnvelope{commands}
}

func CommandsEnvelopeEqualityFunction(lhs, rhs e.Envelope[[]Command]) bool {
	return f.ArrayEqualWith(CommandEqualityFunc)(lhs.Get(), rhs.Get())
}
