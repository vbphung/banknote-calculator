package mike

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
	r.units = NewBanknoter(units).Cal(supply)
}
