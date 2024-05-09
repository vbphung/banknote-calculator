package mike

type Once interface {
	Cal(amount uint64) []uint64
}

type once struct {
	pre [][]uint
}

func NewOnce(units []uint) (Once, bool) {
	o := &once{
		pre: make([][]uint, 9),
	}

	if !o.gen(units) {
		return nil, false
	}

	return o, true
}

func (o *once) Cal(amount uint64) []uint64 {
	dec, used := uint64(1), make([]uint64, 0)
	for amount > 0 {
		n := uint(amount % 10)
		if n > 0 {
			u := o.pre[n-1]
			for _, n := range u {
				used = append(used, dec*uint64(n))
			}
		}

		dec *= 10
		amount /= 10
	}

	return used
}

func (o *once) gen(units []uint) bool {
	for n := uint(1); n <= 9; n++ {
		u := o.of(units, n)
		if u == nil {
			return false
		}

		o.pre[n-1] = u
	}

	return true
}

func (o once) of(units []uint, amt uint) []uint {
	v := make([]bool, len(units))

	var recur func(n uint) []uint
	recur = func(n uint) []uint {
		if n == 0 {
			return []uint{}
		}

		used, usedUnit := make([]uint, 9), -1
		for i, unit := range units {
			if v[i] || n < unit {
				continue
			}

			v[i] = true
			u := recur(n - unit)
			if u != nil && len(u) < len(used) {
				used, usedUnit = u, int(unit)
			}
			v[i] = false
		}

		if usedUnit < 0 {
			return nil
		}

		return append(used, uint(usedUnit))
	}

	return recur(amt)
}
