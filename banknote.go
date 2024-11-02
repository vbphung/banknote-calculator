package banknotecalculator

import (
	"slices"
	"sort"
)

type Banknoter interface {
	Cal(supply uint64) []uint64
}

type banknoter struct {
	units           []uint64
	decimal         uint64
	next, baseUnits int
}

func NewBanknoter(units []uint) Banknoter {
	sort.Slice(units, func(i, j int) bool {
		return units[i] < units[j]
	})

	u := make([]uint64, len(units))
	for i := len(units) - 1; i >= 0; i-- {
		u[i] = uint64(units[i])
	}

	return &banknoter{u, 10, 0, len(units)}
}

func (b *banknoter) Cal(supply uint64) []uint64 {
	if b.units[len(b.units)-1] >= supply {
		for i := len(b.units) - 1; i >= 0; i-- {
			if b.units[i] > supply {
				continue
			}

			return b.copy(i + 1)
		}
	}

	for {
		u := b.decimal * b.units[b.next]
		if u > supply {
			break
		}

		b.units = append(b.units, u)

		if b.next == b.baseUnits-1 {
			b.next, b.decimal = 0, b.decimal*10
		} else {
			b.next++
		}
	}

	return b.copy(len(b.units))
}

func (b *banknoter) copy(pre int) []uint64 {
	cp := make([]uint64, pre)
	copy(cp, b.units[:pre])

	slices.Reverse(cp)

	return cp
}
