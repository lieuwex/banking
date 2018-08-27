package view

type GroupBy int

const (
	ByDay GroupBy = iota
	ByWeek
	ByMonth
	ByYear
)
