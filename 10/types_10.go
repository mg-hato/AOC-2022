package main

import (
	. "aoc/functional"
	"fmt"
	"strings"
)

type Instruction interface {
	CycleLength() int
	DeltaX() int
	String() string
}

// ADDX

type Addx struct {
	arg int
}

func NewAddx(arg int) Instruction {
	return &Addx{arg}
}

func (a *Addx) CycleLength() int {
	return 2
}

func (a *Addx) DeltaX() int {
	return a.arg
}

func (a *Addx) String() string {
	return fmt.Sprintf("addx %d", a.arg)
}

// NOOP

type Noop struct {
}

func NewNoop() Instruction {
	return &Noop{}
}

func (*Noop) CycleLength() int {
	return 1
}

func (*Noop) DeltaX() int {
	return 0
}

func (*Noop) String() string {
	return "noop"
}

type Result interface {
	String() string
}

type ResultInt struct {
	result int
}

func (ri ResultInt) String() string {
	return fmt.Sprintf("%d", ri.result)
}

type ResultCRT struct {
	crt string
}

func (rcrt ResultCRT) String() string {
	runes := []rune(rcrt.crt)
	screen_rows := []string{}
	for len(runes) > 0 {
		screen_rows = append(screen_rows, string(Take(40, runes)))
		runes = Drop(40, runes)
	}
	return fmt.Sprintf("\n%s", strings.Join(screen_rows, "\n"))
}
