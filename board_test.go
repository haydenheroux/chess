package main

import "testing"

func TestBoardMove(t *testing.T) {
	board := NewBoard()
	from := Notation("a2")
	to := Notation("a3")

	_ = board.Move(from, to)

	if board.IsOccupied(from) {
		t.Errorf("From is not empty after move")
	}

	if !board.IsOccupied(to) {
		t.Errorf("To is empty after move")
	}
}

func TestBoardMoveOntoOccupied(t *testing.T) {
	board := NewBoard()
	from := Notation("a2")
	to := Notation("a7")

	ok := board.Move(from, to)

	if ok {
		t.Errorf("Move was okay when it should not have been")
	}

	if !board.IsOccupied(from) {
		t.Errorf("From is empty after move")
	}

	if !board.IsOccupied(to) {
		t.Errorf("To is empty after move")
	}
}

func TestBoardMoveThereAndBack(t *testing.T) {
	board := NewBoard()
	from := Notation("a2")
	to := Notation("b4")

	ok := board.Move(from, to)

	if !ok {
		t.Errorf("Move was not okay, but it should have been okay")
	}

	if board.IsOccupied(from) {
		t.Errorf("From is not empty after first move")
	}

	if !board.IsOccupied(to) {
		t.Errorf("To is empty after first move")
	}

	ok = board.Move(to, from)

	if !ok {
		t.Errorf("Move was not okay, but it should have been okay")
	}

	if !board.IsOccupied(from) {
		t.Errorf("From is empty after final move")
	}

	if board.IsOccupied(to) {
		t.Errorf("To is not empty after final move")
	}
}

