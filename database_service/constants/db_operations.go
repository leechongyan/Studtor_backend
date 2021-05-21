package constants

type Operation int

const (
	Get Operation = iota + 1
	Add
	Delete
)
