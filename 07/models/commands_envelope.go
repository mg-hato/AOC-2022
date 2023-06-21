package models

import (
	c "aoc/common"
)

type CommandsEnvelope struct {
	commands []Command
}

func (ce CommandsEnvelope) Get() []Command {
	return c.Map(Command.Copy, ce.commands)
}

func CreateCommandsEnvelope(commands ...Command) c.Envelope[[]Command] {
	return CommandsEnvelope{commands}
}

func CommandsEnvelopeEqualityFunction(lhs, rhs c.Envelope[[]Command]) bool {
	return c.ArrayEqualWith(CommandEqualityFunc)(lhs.Get(), rhs.Get())
}
