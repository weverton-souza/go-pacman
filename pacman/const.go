package pacman

const (
	backgroundImageSize = 100
	stageBlocSize       = 32
)

type elem int

const (
	w0 elem = iota
	w1
	w2
	w3
	w4
	w5
	w6
	w7
	w8
	w9
	w10
	w11
	w12
	w13
	w14
	w15
	w16
	w17
	w18
	w19
	w20
	w21
	w22
	w23
	w24
	playerElem
	bigDotElem
	dotElem
	empty
	blinkyElem
	clydeElem
	inkyElem
	pinkyElem
	fruitElem
	backgroundElem
)

type input int

const (
	_ input = iota
	up
	right
	down
	left
)
