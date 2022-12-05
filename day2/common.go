package day2

func conv(inp string) shape {
	if inp == "A" || inp == "X" {
		return rock
	} else if inp == "B" || inp == "Y" {
		return paper
	} else {
		return scissors
	}
}

type shape int

const (
	rock shape = iota + 1
	paper
	scissors
)

type state int

const (
	lose state = iota
	draw       = iota * 3
	win
)

func getState(opInp, yourInp shape) state {
	diff := int(yourInp) - int(opInp)
	if diff == 0 {
		return draw
	} else if diff == -1 || diff == 2 {
		return lose
	} else {
		return win
	}
}

func score(onInp, yourInp shape) int {
	state := getState(onInp, yourInp)
	return int(state) + int(yourInp)
}
