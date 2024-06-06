package main

import "fmt"

func main() {
	fmt.Println("a")
}

type Point struct {
	x int
	y int
}

type Direction int32
const (
	Direction_EAST Direction = 1
	Direction_SOUTH Direction = 2
	Direction_WEST Direction = 3
	Direction_North Direction = 4
)

type MoveType int32
const (
	MOVE_FORWARD MoveType = 1
	MOVE_BACK	 MoveType = 2
)

type TurnType int32
const (
	TURN_LEFT 	MoveType = 1
	TURN_RIGHT	MoveType = 2
)

func init(){

}

func Move(type int){

}

func Turn(){

}