package guide

import (
	`github.com/pangum/pangu`
)

type (
	in struct {
		qa     *qa
		answer *answer
	}

	testIn struct {
		pangu.In

		Qa     *qa
		Answer *answer
	}
)

func newIn(ti testIn) *in {
	return &in{
		qa:     ti.Qa,
		answer: ti.Answer,
	}
}
