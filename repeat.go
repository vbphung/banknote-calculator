package mike

import (
	"slices"
	"sort"
)

type Repeat interface {
	Cal(amount uint64) map[uint64]uint
}

type repeat struct {
	units []uint64
}

func NewRepeat(units []uint, supply uint64) Repeat {
	r := &repeat{}
	r.gen(units, supply)

	return r
}

func (r *repeat) Cal(amount uint64) map[uint64]uint {
	used := make(map[uint64]uint)
	for _, unit := range r.units {
		if amount == 0 {
			break
		}

		if amount < unit {
			continue
		}

		used[unit] = uint(amount / unit)
		amount %= unit
	}

	return used
}

func (r *repeat) gen(units []uint, supply uint64) {
	r.units = AllBills(units, supply)
}

func AllBills(units []uint, supply uint64) []uint64 {
	sort.Slice(units, func(i, j int) bool {
		return units[i] < units[j]
	})

	bills := make([]uint64, 0)
	nx, dec := 0, 1
	for {
		u := uint64(units[nx] * uint(dec))
		if u > supply {
			break
		}

		bills = append(bills, u)

		if nx == len(units)-1 {
			nx, dec = 0, dec*10
		} else {
			nx++
		}
	}

	slices.Reverse(bills)
	return bills
}
